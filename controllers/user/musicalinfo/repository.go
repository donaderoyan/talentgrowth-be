package musicalinfo

import (
	"context"
	"fmt"

	model "github.com/donaderoyan/talentgrowth-be/models"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UpdateMusicalInfo(userID string, updates bson.M) (*model.MusicalInfo, error)
}

type repository struct {
	db *mongo.Database
}

func NewMusicalInfoRepository(db *mongo.Database) *repository {
	return &repository{db: db}
}

type MusicalInfoUpdateError struct {
	*util.BaseError
}

func (e *MusicalInfoUpdateError) Error() string {
	return fmt.Sprintf("Profile update error: %s - %s", e.Code, e.Message)
}

func (r *repository) UpdateMusicalInfo(userID string, updates bson.M) (*model.MusicalInfo, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, &MusicalInfoUpdateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
	}
	var updatedMusicalInfo model.MusicalInfo

	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {

		// Check if user exists and retrieve current user data
		var existingUser model.User
		err = r.db.Collection("users").FindOne(sc, bson.M{"_id": userIDPrimitive}).Decode(&existingUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &MusicalInfoUpdateError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
			return err
		}

		updates["userID"] = userIDPrimitive
		// Insert musical information
		result, errInsert := r.db.Collection("musicalinfo").InsertOne(sc, updates)
		if errInsert != nil {
			return &MusicalInfoUpdateError{util.NewBaseError("ADD_MUSICALINFO_ERROR", "Update musical information failed")}
		}

		// Update the User document with the MusicalInfoID
		filter := bson.M{"_id": userIDPrimitive}
		userUpdate := bson.M{
			"$set": bson.M{
				"musical_info_id": result.InsertedID.(primitive.ObjectID),
			},
		}

		_, err = r.db.Collection("users").UpdateOne(sc, filter, userUpdate)
		if err != nil {
			return &MusicalInfoUpdateError{util.NewBaseError("ADD_MUSICALINFO_ERROR", "Update musical information failed")}
		}

		// Retrieve the updated MusicalInfo
		err = r.db.Collection("musicalinfo").FindOne(context.Background(), bson.M{"userID": userIDPrimitive}).Decode(&updatedMusicalInfo)
		if err != nil {
			return &MusicalInfoUpdateError{util.NewBaseError("MUSICAL_INFO_RETRIEVE_ERROR", "Failed to retrieve updated musical information")}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return &updatedMusicalInfo, nil
}
