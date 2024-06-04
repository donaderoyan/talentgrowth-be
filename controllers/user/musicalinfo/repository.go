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
	CreateMusicalInfo(userID string, data bson.M) (*model.MusicalInfo, error)
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
	return fmt.Sprintf("Musical Information update error: %s - %s", e.Code, e.Message)
}

type MusicalInfoCreateError struct {
	*util.BaseError
}

func (e *MusicalInfoCreateError) Error() string {
	return fmt.Sprintf("Musical Information update error: %s - %s", e.Code, e.Message)
}

func (r *repository) CreateMusicalInfo(userID string, data bson.M) (*model.MusicalInfo, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, &MusicalInfoCreateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
	}
	var updatedMusicalInfo model.MusicalInfo

	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {

		// Check if user exists and retrieve current user data
		var existingUser model.User
		err = r.db.Collection("users").FindOne(sc, bson.M{"_id": userIDPrimitive}).Decode(&existingUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &MusicalInfoCreateError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
			return err
		}

		var existingMusicalInfo *model.MusicalInfo
		err := r.db.Collection("musicalinfo").FindOne(sc, bson.M{"userID": userIDPrimitive}).Decode(&existingMusicalInfo)
		if err == mongo.ErrNoDocuments {
			// userID does not exist in musicalinfo, continue with creation
			data["userID"] = userIDPrimitive
		} else if err != nil {
			return err
		} else {
			// Musical info already exists, return error
			return &MusicalInfoCreateError{util.NewBaseError("MUSICAL_INFO_EXIST", "Musical information already exists")}
		}

		// Insert musical information
		result, errInsert := r.db.Collection("musicalinfo").InsertOne(sc, data)
		if errInsert != nil {
			return &MusicalInfoCreateError{util.NewBaseError("ADD_MUSICALINFO_ERROR", "Create musical information failed")}
		}

		// Update the User document with the MusicalInfoID
		filter := bson.M{"_id": userIDPrimitive}
		userUpdate := bson.M{
			"$set": bson.M{
				"musicalInfoId": result.InsertedID.(primitive.ObjectID),
			},
		}

		_, err = r.db.Collection("users").UpdateOne(sc, filter, userUpdate)
		if err != nil {
			return &MusicalInfoCreateError{util.NewBaseError("ADD_MUSICALINFO_ERROR", "Create musical information failed")}
		}

		// Retrieve the updated MusicalInfo
		err = r.db.Collection("musicalinfo").FindOne(context.Background(), bson.M{"userID": userIDPrimitive}).Decode(&updatedMusicalInfo)
		if err != nil {
			return &MusicalInfoCreateError{util.NewBaseError("MUSICAL_INFO_RETRIEVE_ERROR", "Failed to retrieve updated musical information")}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return &updatedMusicalInfo, nil
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

		// Insert musical information
		_, errInsert := r.db.Collection("musicalinfo").UpdateOne(sc, bson.M{"userID": userIDPrimitive}, bson.M{"$set": updates})
		if errInsert != nil {
			fmt.Printf("INSERT MUSICAL INFORMATION ERRORR >>>>>>>>>>> %v", errInsert)
			return &MusicalInfoUpdateError{util.NewBaseError("ADD_MUSICALINFO_ERROR", "Update musical information failed")}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	// Retrieve the updated MusicalInfo
	var updatedMusicalInfo model.MusicalInfo
	err = r.db.Collection("musicalinfo").FindOne(context.Background(), bson.M{"userID": userIDPrimitive}).Decode(&updatedMusicalInfo)
	if err != nil {
		return nil, &MusicalInfoUpdateError{util.NewBaseError("MUSICAL_INFO_RETRIEVE_ERROR", "Failed to retrieve updated musical information")}
	}

	return &updatedMusicalInfo, nil
}
