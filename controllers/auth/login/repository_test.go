package loginController

import (
	"context"
	"os"
	"testing"

	model "github.com/donaderoyan/talentgrowth-be/models"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoginRepositorySuite struct {
	suite.Suite
	db     *mongo.Database
	repo   Repository
	client *mongo.Client
}

func (suite *LoginRepositorySuite) SetupSuite() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:password123@localhost:27017/?authSource=admin" // Default URI, change if necessary
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.Require().NoError(err)
	suite.client = client
	suite.db = client.Database("testdb")
	suite.repo = NewLoginRepository(suite.db)
}

func TestLoginRepositorySuite(t *testing.T) {
	suite.Run(t, new(LoginRepositorySuite))
}

func (suite *LoginRepositorySuite) TearDownSuite() {
	suite.client.Disconnect(context.Background())
}

func (suite *LoginRepositorySuite) TestLoginRepository() {
	// Clear the users collection before starting the test
	err := suite.db.Collection("users").Drop(context.Background())
	suite.Require().NoError(err, "Expected no error in dropping collection")

	hashedPassword, hashErr := util.HashPassword("password123")
	suite.Require().NoError(hashErr, "Expected no error in hashing password")

	user := &model.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	testUser := &model.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Insert a user to simulate login
	insertResult, err := suite.db.Collection("users").InsertOne(context.Background(), user)
	suite.Require().NoError(err, "Expected no error in inserting user")
	suite.Require().NotNil(insertResult.InsertedID, "Expected an inserted ID for the new user")

	// Test login
	loggedInUser, err := suite.repo.LoginRepository(testUser)
	suite.Require().NoError(err, "Expected no error during login")
	suite.NotNil(loggedInUser, "Expected user to be logged in")
	suite.Equal(loggedInUser.Email, user.Email, "Expected email to match")
}

func (suite *LoginRepositorySuite) TestLoginRepositoryWithWrongPassword() {
	user := &model.User{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	// Test login with wrong password
	_, err := suite.repo.LoginRepository(user)
	suite.Require().Error(err, "Expected error due to wrong password")
	suite.IsType(&UserLoginError{}, err)
}

func (suite *LoginRepositorySuite) TestLoginRepositoryWithNonExistentUser() {
	user := &model.User{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	// Test login with non-existent user
	_, err := suite.repo.LoginRepository(user)
	suite.Require().Error(err, "Expected error due to non-existent user")
	suite.IsType(&UserLoginNotFoundError{}, err)
}
