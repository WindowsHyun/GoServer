package usecase

import (
	"GoServer/model"
)

type UserUsecase struct {
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Register(user *model.User) error {
	return nil
}

func (u *UserUsecase) Login(email, password string) (*model.User, error) {
	return &model.User{}, nil
}
