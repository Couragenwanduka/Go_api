package Userservice

import (
	"api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"fmt"
)
type User struct {
	FirstName string
	LastName string
	Email string
	Password string
	Address string
	Phone string
}

func SaveUser(user User, client *mongo.Client) (*mongo.InsertOneResult, error) {
	// Implement user creation logic here
    // Example: Save user to database
    // db.collection.insertOne(user)
	collection := client.Database("simple-api").Collection("users")
	newMovie := model.Item{
		ID:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
        Password: user.Password,
        Address: user.Address,
        Phone: user.Phone,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, newMovie)
	if err != nil {
	return nil, fmt.Errorf("failed to insert user: %v", err)
	}
	fmt.Printf("Inserted user with ID: %v\n", result.InsertedID)
	return result, nil
}

// Implement other CRUD operations here
func FindUserByEmail(email string, client *mongo.Client) (*User, error) {
	collection := client.Database("simple-api").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a BSON filter
	filter := bson.M{"email": email}

	// Create a User object to store the result
	var user User

	// Perform the query and decode the result
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err // Return other errors
	}

	return &user, nil
}
