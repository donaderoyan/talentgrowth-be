package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "github.com/donaderoyan/talentgrowth-be/configs"
)

func main() {
	mongo := config.ConnectMongoDB()
	client := mongo.Database("talentgrowth")
	defer mongo.Disconnect(context.Background())

	fieldsToAdd, collectionName := prepareMigrationData()
	for i, collName := range collectionName {
		if err := MigrateData(client, collName, fieldsToAdd[i]); err != nil {
			log.Fatalf("Migration failed for collection %s: %v", collName, err)
		}
	}

	log.Println("Migration completed successfully")
}

// MigrateData dynamically updates all documents in the specified collection to include specified fields with default values
func MigrateData(client *mongo.Database, collectionName string, fieldsToAdd bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	collection := client.Collection(collectionName)

	update := bson.M{"$set": fieldsToAdd}
	opts := options.Update().SetUpsert(true)

	// Update all documents
	_, err := collection.UpdateMany(ctx, bson.M{}, update, opts)
	if err != nil {
		log.Printf("Failed to migrate %v data: %v", collectionName, err)
		return err
	}

	log.Printf("%v data migration completed successfully", collectionName)
	return nil
}

func prepareMigrationData() ([]bson.M, []string) {
	var migrations []bson.M
	var collectionNames []string

	// Migration for user profile pictures and addresses
	userFieldsJSON := `{"profilePicture": "", "address": {"street": "", "city": "", "state": "", "postalCode": "", "country": ""}}`
	var userFields bson.M
	if err := bson.UnmarshalExtJSON([]byte(userFieldsJSON), true, &userFields); err != nil {
		log.Fatalf("Invalid JSON for user fields to add: %v", err)
	}
	migrations = append(migrations, userFields)
	collectionNames = append(collectionNames, "users")

	// Additional migrations can be appended here in a similar pattern
	// Example:
	// projectFieldsJSON := `{"projectTitle": "", "projectDescription": ""}`
	// var projectFields bson.M
	// if err := bson.UnmarshalExtJSON([]byte(projectFieldsJSON), true, &projectFields); err != nil {
	//     log.Fatalf("Invalid JSON for project fields to add: %v", err)
	// }
	// migrations = append(migrations, projectFields)
	// collectionNames = append(collectionNames, "projects")

	return migrations, collectionNames
}
