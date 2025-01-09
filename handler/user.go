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

// User
// @Summary
// @Description User Regist
// @Tags User
// @Accept json
// @Produce json
// @Param reqBody body structure.ReqRegist true "User registration request body"
// @Success 200 {object} structure.ResRegist
// @Failure 400 "Bad Request Error"
// @Failure 502 "Could Not Be Searched In DB Collection"
// @Failure 506 "DB Collection Update Error"
// @Failure 509 "JWT Token Error"
// @Router /user/regist [post]
func (h *userHandler) Register(c *gin.Context) {

}

// User
// @Summary
// @Description User Login
// @Tags User
// @Accept json
// @Produce json
// @Param reqBody body structure.ReqLogin true "User login request body"
// @Success 200 {object} structure.ResLogin
// @Failure 400 "Bad Request Error"
// @Router /user/login [post]
func (h *userHandler) Login(c *gin.Context) {

}
