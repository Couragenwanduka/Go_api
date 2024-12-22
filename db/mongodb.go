package db

import (
	"log"
	"os"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB instance and returns the client
func Connect() *mongo.Client {
	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get MongoDB URI
	uri := os.Getenv("mongoDbUri")
	// uri := "mongodb+srv://courageobunike:nYut1ZfnApzcEfPq@cluster0.uqro2.mongodb.net/"
	if uri == "" {
		log.Fatal("No MongoDB URI provided in environment")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Defer disconnection
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	log.Println("successfully connected to MongoDB")
	return client
}
