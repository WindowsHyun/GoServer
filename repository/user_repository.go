package repository

import (
	"GoServer/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByID(int) (*model.User, error)
	SaveUser(*model.User) error
	RegisterUser(ctx context.Context, user model.User) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) RegisterUser(ctx context.Context, user model.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (r *MongoUserRepository) SaveUser(user *model.User) error {
	return nil
}
