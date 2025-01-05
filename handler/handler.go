package handler

import (
	"GoServer/handler/admin"
	"GoServer/handler/user"
	"GoServer/usecase"
)

type UserHandler struct {
}

func NewUserHandler() *user.Handler {
	uc := usecase.NewUserUsecase()
	return user.NewHandler(uc)
}

type AdminHandler struct {
}

func NewAdminHandler() *admin.Handler {
	return admin.NewHandler()
}
