package router

import (
	"GoServer/handler"
	"GoServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	// 공개 Router
	public := r.Group("/")
	{
		public.POST("/user/regist", handler.UserHandler.Register)
		public.POST("/user/login", handler.UserHandler.Login)
	}

	// JWT 사용
	auth := r.Group("/").Use(middleware.JWTMiddleware())
	{
		auth.GET("/admin/info", handler.AppHandler.GetMenu)
	}

	return r
}
