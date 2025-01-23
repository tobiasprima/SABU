package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient     *mongo.Client
	MealsCollection *mongo.Collection
)

// InitMongo initializes the MongoDB connection
func InitMongo() error {
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URL environment variable is not set")
	}

	databaseName := os.Getenv("MONGO_DB_NAME")
	if databaseName == "" {
		return fmt.Errorf("MONGO_DB_NAME environment variable is not set")
	}

	collectionName := os.Getenv("MONGO_COLLECTION")
	if collectionName == "" {
		return fmt.Errorf("MONGO_COLLECTION environment variable is not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	// Initialize MealsCollection
	MealsCollection = client.Database(databaseName).Collection(collectionName)
	MongoClient = client
	return nil
}