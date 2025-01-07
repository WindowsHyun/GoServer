package usecase

import (
	IMongo "GoServer/database/mongo"
	"GoServer/model"
	"context"
)

type appRepository struct {
	mongo IMongo.MongoInterface
}

type AppUsecase interface {
	GetMenu(ctx context.Context) ([]model.Menu, error)
}

func NewAppRepository(mongo IMongo.MongoInterface) AppUsecase {
	return &appRepository{
		mongo: mongo,
	}
}

func (a *appRepository) GetMenu(ctx context.Context) ([]model.Menu, error) {
	return nil, nil
}
