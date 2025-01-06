package router

import (
	"GoServer/handler"
	"GoServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	// 공개 Router
	public := r.Group("/")
	{
		public.POST("/user/regist", userHandler.Register)
		public.POST("/user/login", userHandler.Login)
	}

	// JWT 사용
	auth := r.Group("/").Use(middleware.JWTMiddleware())
	{
		auth.GET("/admin/info")
	}

	return r
}
