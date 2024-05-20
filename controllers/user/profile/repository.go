package profileController

import (
	"context"
	"fmt"
	"time"

	model "github.com/donaderoyan/talentgrowth-be/models"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UpdateProfile(userID string, user *model.User) (*model.User, error)
}

type repository struct {
	db *mongo.Database
}

func NewProfileRepository(db *mongo.Database) *repository {
	return &repository{db: db}
}

type UserProfileUpdateError struct {
	*util.BaseError
}

func (e *UserProfileUpdateError) Error() string {
	return fmt.Sprintf("Profile update error: %s - %s", e.Code, e.Message)
}

func (r *repository) UpdateProfile(userID string, user *model.User) (*model.User, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, &UserProfileUpdateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
	}

	// Transaction handling
	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {

		// Check if user exists and retrieve current user data
		var existingUser model.User
		err = r.db.Collection("users").FindOne(sc, bson.M{"_id": userIDPrimitive}).Decode(&existingUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &UserProfileUpdateError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
			return err
		}
		// Ensure correct field names are used for MongoDB document
		correctFieldNames := bson.M{
			"firstName":          user.FirstName,
			"lastName":           user.LastName,
			"phone":              user.Phone,
			"birthday":           user.Birthday,
			"gender":             user.Gender,
			"nationality":        user.Nationality,
			"bio":                user.Bio,
			"updatedAt":          time.Now(),
			"profilePicture":     user.ProfilePicture,
			"address.street":     user.Address.Street,
			"address.city":       user.Address.City,
			"address.state":      user.Address.State,
			"address.postalCode": user.Address.PostalCode,
			"address.country":    user.Address.Country,
		}

		// Update user profile with only the specified fields
		_, err = r.db.Collection("users").UpdateByID(sc, userIDPrimitive, bson.M{"$set": correctFieldNames})
		if err != nil {
			fmt.Printf("Error updating user profile: %v", err)
			return &UserProfileUpdateError{util.NewBaseError("USER_PROFILE_UPDATE_ERROR", "Profile update failed")}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	// Retrieve the updated user profile
	var updatedUser model.User
	err = r.db.Collection("users").FindOne(context.Background(), bson.M{"_id": userIDPrimitive}).Decode(&updatedUser)
	if err != nil {
		return nil, &UserProfileUpdateError{util.NewBaseError("USER_PROFILE_RETRIEVE_ERROR", "Failed to retrieve updated profile")}
	}

	return &updatedUser, nil
}
