package config

import (
	"context"
	"time"

	util "github.com/donaderoyan/talentgrowth-be/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := util.GodotEnv("MONGO_URI")
	if mongoURI == "" {
		logrus.Fatal("MONGO_URI must be set in the environment variables")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri": mongoURI,
		}).Fatal("Failed to connect to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri": mongoURI,
		}).Fatal("Failed to ping MongoDB:", err)
	}

	logrus.Info("Connected to MongoDB successfully")
	return client
}
