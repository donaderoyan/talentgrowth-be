package registerController

import (
	"context"
	"testing"

	"github.com/benweissmann/memongo"
	model "github.com/donaderoyan/talentgrowth-be/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestRegisterRepository(t *testing.T) {
	// Create a new in-memory MongoDB server
	server, err := memongo.StartWithOptions(&memongo.Options{MongoVersion: "5.0.0"})
	if err != nil {
		t.Fatalf("Failed to start MongoDB server: %v", err)
	}
	defer server.Stop()

	// Connect to the in-memory MongoDB server
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(server.URI()))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB server: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Create a new database
	db := client.Database("testdb")

	// Create a new repository instance
	repo := NewRegisterRepository(db)

	// Define test cases
	tests := []struct {
		name    string
		user    *model.User
		wantErr bool
		errType error
	}{
		{
			name: "Successful Registration",
			user: &model.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "User Already Exists",
			user: &model.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantErr: true,
			errType: &UserAlreadyExistsError{},
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Register user
			result, err := repo.RegisterRepository(tt.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				// Check if the error type is correct
				if _, ok := err.(*UserAlreadyExistsError); !ok {
					t.Errorf("Expected error of type *UserAlreadyExistsError, got %T", err)
				}
			} else if !tt.wantErr && result != nil {
				// Check if the user ID is set
				if result.ID == primitive.NilObjectID {
					t.Errorf("Expected user ID to be set, got %v", result.ID)
				}
			}
		})
	}
}
