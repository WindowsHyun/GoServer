package usecase

import (
	IMongo "GoServer/database/mongo"
	"GoServer/model"
	"context"
)

type userRepository struct {
	mongo IMongo.MongoInterface
}

type UserUsecase interface {
	GetUserByID(int) (*model.User, error)
	SaveUser(*model.User) error
	RegisterUser(ctx context.Context, user model.User) error
}

func NewUserRepository(mongo IMongo.MongoInterface) UserUsecase {
	return &userRepository{
		mongo: mongo,
	}
}

func (u *userRepository) RegisterUser(ctx context.Context, user model.User) error {
	err := u.mongo.Insert(ctx, user)
	return err
}

func (u *userRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (u *userRepository) SaveUser(user *model.User) error {
	return nil
}
