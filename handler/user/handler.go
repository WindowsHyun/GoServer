package user

import (
	"GoServer/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.UserUsecase
}

func NewHandler(usecase *usecase.UserUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}
