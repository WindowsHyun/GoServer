package repository

import (
	"GoServer/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppRepository interface {
	GetMenu(ctx context.Context) ([]model.Menu, error)
	// Additional methods
}

type MongoAppRepository struct {
	collection *mongo.Collection
}

func NewMongoAppRepository(collection *mongo.Collection) *MongoAppRepository {
	return &MongoAppRepository{collection: collection}
}

func (r *MongoAppRepository) GetMenu(ctx context.Context) ([]model.Menu, error) {
	// Implement the method
	return nil, nil
}
