package handler

import (
	"GoServer/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type appHandler struct {
	appUsecase usecase.AppUsecase
}

type AppHandler interface {
	GetMenu(c *gin.Context)
}

func NewAppHandler(appUsecase usecase.AppUsecase) (AppHandler, error) {
	if appUsecase == nil {
		return nil, fmt.Errorf("appUsecase is nil")
	}

	return &appHandler{
		appUsecase: appUsecase,
	}, nil
}

// App
// @Summary
// @Description App Menu
// @Tags App
// @Accept json
// @Produce json
// @Success 200 {object} structure.ResDefaultMessage
// @Failure 400 "Bad Request Error"
// @Router /app/menu [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (a *appHandler) GetMenu(c *gin.Context) {

}
