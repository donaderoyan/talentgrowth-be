package register

import (
	"context"
	"os"
	"testing"

	model "github.com/donaderoyan/talentgrowth-be/models"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RegisterRepositorySuite struct {
	suite.Suite
	db     *mongo.Database
	repo   *repository
	client *mongo.Client
}

func (suite *RegisterRepositorySuite) SetupSuite() {
	// Retrieve the MongoDB URI from environment variables. Modify the default URI below if needed before starting tests, or consult the README.
	mongoURI := os.Getenv("MONGO_URI")
	// if mongoURI == "" {
	// 	mongoURI = "mongodb://root:password123@localhost:27017/?authSource=admin" // Default URI, change if necessary
	// }

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.Require().NoError(err)
	suite.client = client
	suite.db = client.Database("testdb")
	suite.repo = NewRegisterRepository(suite.db)
}

func TestRegisterRepositorySuite(t *testing.T) {
	suite.Run(t, new(RegisterRepositorySuite))
}

func (suite *RegisterRepositorySuite) TearDownSuite() {
	suite.client.Disconnect(context.Background())
}

func (suite *RegisterRepositorySuite) TestRegisterNewUser() {
	// Clear the users collection before starting the test
	err := suite.db.Collection("users").Drop(context.Background())
	suite.Require().NoError(err, "Expected no error in dropping collection")

	user := &model.User{
		Email:    "test10@example.com",
		Password: "password123",
	}

	registeredUser, err := suite.repo.RegisterRepository(user)
	suite.Require().NoError(err, "Expected no error")
	suite.NotNil(registeredUser, "Expected user to be registered")
	suite.Equal(registeredUser.Email, "test10@example.com")
}

func (suite *RegisterRepositorySuite) TestRegisterExistingUser() {
	// Clear the users collection before starting the test
	err := suite.db.Collection("users").Drop(context.Background())
	suite.Require().NoError(err, "Expected no error in dropping collection")

	user := &model.User{
		Email:    "test10@example.com",
		Password: "password123",
	}

	// First insert the user to simulate existing user scenario
	_, err = suite.repo.RegisterRepository(user)
	suite.Require().NoError(err, "Expected no error on initial user registration for test setup")

	// Attempt to register the same user again
	_, err = suite.repo.RegisterRepository(user)
	suite.Require().Error(err, "Expected error when registering an existing user")
	suite.IsType(&UserAlreadyExistsError{}, err)
}
