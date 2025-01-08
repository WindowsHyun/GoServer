package handler

import (
	"GoServer/config"
	"GoServer/usecase"

	"github.com/pkg/errors"
)

type Handler struct {
	UserHandler UserHandler
	AppHandler  AppHandler
}

func InitHandler(config *config.Config, usecase *usecase.Usecase) (*Handler, error) {
	userHandler, err := NewUserHandler(usecase.UserUsecase)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize userHandler")
	}
	appHandler, err := NewAppHandler(usecase.AppUsecase)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize appHandler")
	}

	return &Handler{
		UserHandler: userHandler,
		AppHandler:  appHandler,
	}, nil
}
