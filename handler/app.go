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

func (a *appHandler) GetMenu(c *gin.Context) {

}
