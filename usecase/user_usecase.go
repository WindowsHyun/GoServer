package usecase

import (
	"GoServer/model"
	"GoServer/repository"
	"context"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user model.User) error {
	return u.repo.RegisterUser(ctx, user)
}
