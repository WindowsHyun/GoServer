package main

import (
	"GoServer/config"
	"GoServer/database/mongo"
	"GoServer/database/mysql"
	"GoServer/database/redis"
	"GoServer/logger"
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
		log.Panicf("Failed to initialize MongoDB clients: %v", err)
	}
	defer mongo.Close(ctx)

	mysqlRepo, err := mysql.Initialize(ctx, cf)
	if err != nil {
		log.Panicf("Failed to initialize MySQL clients: %v", err)
	}
	defer mysql.Close(ctx)

	redisRepo, err := redis.Initialize(ctx, cf)
	if err != nil {
		log.Panicf("Failed to initialize Redis clients: %v", err)
	}
	defer redis.Close(ctx)

	// Initialize repositories
	// repos := repository.NewRepositories()

	// Initialize usecase
	usecase.InitUsecase(mongoRepo, mysqlRepo, redisRepo)

	// Initialize handlers
	// userHandler := handler.NewUserHandler(userUsecase)

	// Set up router
	// r := router.SetupRouter(userHandler)

	// Start the server
	// log.Fatal(r.Run(":8080"))
}
