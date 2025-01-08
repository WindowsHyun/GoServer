package handler

import (
	"GoServer/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.Usecase) UserHandler {
	return &userHandler{
		userUsecase: usecase.UserUsecase,
	}
}

func (h *userHandler) Register(c *gin.Context) {

}

func (h *userHandler) Login(c *gin.Context) {

}
