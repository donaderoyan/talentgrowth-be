package loginController

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
	LoginRepository(user *model.User) (*model.User, error)
	UpdateRememberTokenRepository(userID primitive.ObjectID, token string) error
}

type repository struct {
	db *mongo.Database
}

func NewLoginRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

type UserLoginError struct {
	*util.BaseError
}

type UserLoginNotFoundError struct {
	*util.BaseError
}

func (e *UserLoginError) Error() string {
	return fmt.Sprintf("Login error: %s - %s", e.Code, e.Message)
}

func (e *UserLoginNotFoundError) Error() string {
	return fmt.Sprintf("Login error: %s - %s", e.Code, e.Message)
}

func (r *repository) LoginRepository(user *model.User) (*model.User, error) {
	// Start a session for transaction
	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Transaction handling
	transactionErr := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		// Check if user exists and password is correct
		var foundUser model.User
		err := r.db.Collection("users").FindOne(sc, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return &UserLoginNotFoundError{util.NewBaseError("USER_NOT_FOUND", "User not found")}
			}
			return err
		}
		// Compare hashed password
		passwordMatch := util.ComparePassword(foundUser.Password, user.Password)
		if passwordMatch != nil {
			return &UserLoginError{util.NewBaseError("INVALID_CREDENTIALS", "Invalid email or password")}
		}

		*user = foundUser
		return nil
	})

	if transactionErr != nil {
		return nil, transactionErr
	}

	return user, nil
}

func (r *repository) UpdateRememberTokenRepository(userID primitive.ObjectID, token string) error {
	ctx := context.Background()
	filter := bson.M{"_id": bson.M{"$eq": userID}}
	update := bson.M{"$set": bson.M{"rememberToken": token}}
	_, err := r.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
