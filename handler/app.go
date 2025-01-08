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

func NewAppHandler(usecase *usecase.Usecase) AppHandler {
	return &appHandler{
		appUsecase: usecase.AppUsecase,
	}
}

func (a *appHandler) GetMenu(c *gin.Context) {

}
