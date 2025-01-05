package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID string             `json:"id" bson:"id"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
}
