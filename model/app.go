package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Link string             `json:"link" bson:"link"`
}
