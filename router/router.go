package router

import (
	"GoServer/config"
	"GoServer/config/util"
	"GoServer/docs"
	"GoServer/handler"
	"GoServer/logger"
	"GoServer/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(config *config.Config, handler *handler.Handler) *gin.Engine {
	r := gin.New()
	if config.IsDevelop() {
		r.Use(gin.Logger())
	}
	r.Use(logger.GinLogrus())
	r.Use(logger.PanicMiddleware())
	r.Use(util.RealIP())

	// 개발 swagger
	if config.IsDevelop() {
		swag := r.Group("/swagger")
		{
			docs.SwaggerInfo.BasePath = ""
			swag.GET("/*any", swagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	// Public Router
	public := r.Group("/")
	{
		public.POST("/user/regist", handler.UserHandler.Register)
		public.POST("/user/login", handler.UserHandler.Login)
	}

	// JWT Router
	auth := r.Group("/").Use(middleware.JWTMiddleware())
	{
		auth.GET("/app/menu", handler.AppHandler.GetMenu)
	}

	return r
}
