package profile

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
	UpdateProfile(userID string, updates bson.M) (*model.User, error)
	PatchProfile(userID string, updates bson.M) (*model.User, error)
	GetProfile(userID string) (*model.User, error)
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

type GetUserProfileError struct {
	*util.BaseError
}

func (e *GetUserProfileError) Error() string {
	return fmt.Sprintf("Get user profile error: %s - %s", e.Code, e.Message)
}

func (r *repository) UpdateProfile(userID string, updates bson.M) (*model.User, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, &UserProfileUpdateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
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

		// Update user profile with only the specified fields
		_, err = r.db.Collection("users").UpdateByID(sc, userIDPrimitive, bson.M{"$set": updates})
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

func (r *repository) PatchProfile(userID string, updates bson.M) (*model.User, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, &UserProfileUpdateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
	}
	var dataUser model.User
	// Transaction handling
	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {

		// Check if user exists and retrieve current user data
		err = r.db.Collection("users").FindOne(sc, bson.M{"_id": userIDPrimitive}).Decode(&dataUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &UserProfileUpdateError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
			return err
		}

		for key, value := range updates {
			switch key {
			case "firstName":
				dataUser.FirstName = value.(string)
			case "lastName":
				dataUser.LastName = value.(string)
			case "phone":
				dataUser.Phone = value.(string)
			case "birthday":
				parsedBirthday, err := time.Parse("02-01-2006", value.(string))
				if err != nil {
					return err
				}
				dataUser.Birthday = parsedBirthday
			case "gender":
				dataUser.Gender = value.(string)
			case "nationality":
				dataUser.Nationality = value.(string)
			case "bio":
				dataUser.Bio = value.(string)
			case "profilePicture":
				dataUser.ProfilePicture = value.(string)
			}
		}

		// handle embeded Address
		if address, ok := updates["address"]; ok {
			addressMap, _ := address.(map[string]interface{})

			for key, value := range addressMap {
				switch key {
				case "street":
					dataUser.Address.Street = value.(string)
				case "city":
					dataUser.Address.City = value.(string)
				case "state":
					dataUser.Address.State = value.(string)
				case "country":
					dataUser.Address.Country = value.(string)
				case "postalCode":
					dataUser.Address.PostalCode = value.(string)
				}
			}
		}

		// Update user profile with only the specified fields
		_, err = r.db.Collection("users").UpdateByID(sc, userIDPrimitive, bson.M{"$set": dataUser})
		if err != nil {
			fmt.Printf("Error updating user profile: %v", err)
			return &UserProfileUpdateError{util.NewBaseError("USER_PROFILE_UPDATE_ERROR", "Profile update failed")}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return &dataUser, nil
}

func (r *repository) GetProfile(userID string) (*model.User, error) {
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Convert userID to primitive.ObjectID
	userIDPrimitive, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, &UserProfileUpdateError{util.NewBaseError("INVALID_USER_ID", "Invalid user ID format")}
	}
	var dataUser model.User
	// Transaction handling
	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {

		err = r.db.Collection("users").FindOne(sc, bson.M{"_id": userIDPrimitive}).Decode(&dataUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &GetUserProfileError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
		}

		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return &dataUser, nil
}
