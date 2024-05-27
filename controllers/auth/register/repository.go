package register

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
	RegisterRepository(user *model.User) (*model.User, error)
}

type repository struct {
	db *mongo.Database
}

func NewRegisterRepository(db *mongo.Database) *repository {
	return &repository{db: db}
}

type UserAlreadyExistsError struct {
	*util.BaseError
}

type UserRegistrationError struct {
	*util.BaseError
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("Register error: %s - %s", e.Code, e.Message)
}

func (e *UserRegistrationError) Error() string {
	return fmt.Sprintf("Register error: %s - %s", e.Code, e.Message)
}

func (r *repository) RegisterRepository(user *model.User) (*model.User, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Transaction handling
	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		// Check if user already exists
		count, err := r.db.Collection("users").CountDocuments(sc, bson.M{"email": user.Email})
		if err != nil {
			return err
		}
		if count > 0 {
			return &UserAlreadyExistsError{util.NewBaseError("USER_ALREADY_EXISTS", "Email already registered")}
		}

		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			return &UserRegistrationError{util.NewBaseError("USER_REGISTRATION_ERROR", "Registration failed")}
		}
		user.Password = hashedPassword

		// Insert the new user
		result, err := r.db.Collection("users").InsertOne(sc, user)
		if err != nil {
			return &UserRegistrationError{util.NewBaseError("USER_REGISTRATION_ERROR", "Registration failed")}
		}

		// Update the user's ID with the new InsertedID
		user.ID = result.InsertedID.(primitive.ObjectID)
		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return user, nil
}
