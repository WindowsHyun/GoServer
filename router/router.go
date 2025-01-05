package router

import (
	"GoServer/handler"
	"GoServer/handler/admin"
	"GoServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	// 공개 Router
	public := r.Group("/")
	{
		userHandler := handler.NewUserHandler()
		public.POST("/user/regist", userHandler.Register)
		public.POST("/user/login", userHandler.Login)
	}

	// JWT 사용
	auth := r.Group("/").Use(middleware.JWTMiddleware())
	{
		adminHandler := admin.NewHandler()
		auth.GET("/admin/info", adminHandler.Info)
	}

	return r
}
