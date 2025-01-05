package mongo

import (
	"GoServer/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{collection: db.Collection("users")}
}

func (r *MongoUserRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (r *MongoUserRepository) SaveUser(user *model.User) error {
	return nil
}
