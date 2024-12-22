package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}