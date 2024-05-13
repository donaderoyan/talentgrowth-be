package config

import (
	"context"
	"strings"
	"time"

	model "github.com/donaderoyan/talentgrowth-be/models"
	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Increased timeout to avoid premature timeout issues
	defer cancel()

	mongoURI := util.GodotEnv("MONGO_URI")
	if mongoURI == "" {
		logrus.Fatal("MONGO_URI must be set in the environment variables")
	}

	// Adding retryWrites=false to the URI to handle potential issues with connection stability
	if !strings.Contains(mongoURI, "retryWrites") {
		mongoURI += "&retryWrites=false"
	}

	clientOptions := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5 * time.Second) // Setting server selection timeout
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri": mongoURI,
		}).Fatal("Failed to connect to MongoDB:", err)
	}

	// Check the connection with a longer timeout context
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri": mongoURI,
		}).Fatal("Failed to ping MongoDB:", err)
	}

	for _, modelItem := range model.RegisterModels() {
		db := client.Database(util.GodotEnv("MONGO_DBNAME")) // Replace 'yourDatabaseName' with your actual database name
		collection := db.Collection(modelItem)               // Assuming ModelName is a field that holds the collection name
		if collection == nil {
			logrus.Fatal("Failed to get collection for model:", modelItem)
		}
	}

	logrus.Info("Connected to MongoDB successfully")
	return client
}
