package handler

import (
	"GoServer/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) (UserHandler, error) {
	if userUsecase == nil {
		return nil, fmt.Errorf("userUsecase is nil")
	}

	return &userHandler{
		userUsecase: userUsecase,
	}, nil
}

func (h *userHandler) Register(c *gin.Context) {

}

func (h *userHandler) Login(c *gin.Context) {

}
