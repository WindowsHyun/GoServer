package usecase

import (
	IMongo "GoServer/database/mongo"
	"GoServer/model"
	"context"

	"github.com/pkg/errors"
)

type userRepository struct {
	userInfo IMongo.MongoInterface
	menu     IMongo.MongoInterface
}

type UserUsecase interface {
	GetUserByID(int) (*model.User, error)
	SaveUser(*model.User) error
	RegisterUser(ctx context.Context, user model.User) error
}

func NewUserRepository(mongo map[string]IMongo.MongoInterface) (UserUsecase, error) {
	if err := IsValidRepoKey(mongo, "UserInfo", "Menu"); err != nil {
		return nil, errors.Wrap(err, "Invalid Mongo key")
	}

	return &userRepository{
		userInfo: mongo["UserInfo"],
		menu:     mongo["Menu"],
	}, nil
}

func (u *userRepository) RegisterUser(ctx context.Context, user model.User) error {
	err := u.userInfo.Insert(ctx, user)
	return err
}

func (u *userRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (u *userRepository) SaveUser(user *model.User) error {
	return nil
}
