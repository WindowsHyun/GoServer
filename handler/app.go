package handler

import (
	"GoServer/usecase"

	"github.com/gin-gonic/gin"
)

type appHandler struct {
	appUsecase usecase.AppUsecase
}

type AppHandler interface {
	GetMenu(c *gin.Context)
}

func NewAppHandler(appUsecase usecase.AppUsecase) AppHandler {
	return &appHandler{
		appUsecase: appUsecase,
	}
}

func (a *appHandler) GetMenu(c *gin.Context) {

}
