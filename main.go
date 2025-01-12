package main

import (
	"GoServer/config"
	"GoServer/database/mongo"
	"GoServer/database/mysql"
	"GoServer/database/redis"
	"GoServer/handler"
	"GoServer/logger"
	"GoServer/middleware"
	"GoServer/router"
	"GoServer/usecase"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cf := config.InitConfig()
	if cf == nil {
		log.Panic("init config failed\n")
		return
	}

	if err := logger.InitLogger(cf, "GoServer"); err != nil {
		log.Panicf("init logger failed, err : %s\n", err.Error())
		return
	}

	mongoRepo, err := mongo.Initialize(ctx, cf)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize MongoDB clients: %v", err)
		return
	}
	defer mongo.Close(ctx)

	mysqlRepo, err := mysql.Initialize(ctx, cf)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize MySQL clients: %v", err)
		return
	}
	defer mysql.Close(ctx)

	redisRepo, err := redis.Initialize(ctx, cf)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize Redis clients: %v", err)
		return
	}
	defer redis.Close(ctx)

	// Initialize usecase
	usecase, err := usecase.InitUsecase(cf, mongoRepo, mysqlRepo, redisRepo)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize usecase: %v", err)
		return
	}

	// Initialize handlers
	handler, err := handler.InitHandler(cf, usecase)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize handler: %v", err)
		return
	}

	// Initialize middleware
	err = middleware.Init(cf, redisRepo)
	if err != nil {
		logger.LogPanic(ctx, "Failed to initialize middleware: %v", err)
		return
	}

	// Set up router
	r := router.SetupRouter(cf, handler)

	// Start the server
	err = r.Run(":" + cf.GetServer().Port)
	if err != nil {
		logger.LogPanic(ctx, "Failed to start server: %v", err)
		return
	}
}
