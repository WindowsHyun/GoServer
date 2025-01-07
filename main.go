package main

import (
	"GoServer/config"
	"GoServer/database/mongo"
	"GoServer/logger"
	"GoServer/usecase"
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

	mongoRepo, err := mongo.InitializeMongo(ctx, cf)
	if err != nil {
		log.Panicf("Failed to initialize MongoDB clients: %v", err)
	}
	defer mongo.CloseMongo(ctx)

	// Initialize repositories
	// repos := repository.NewRepositories()

	// Initialize usecase
	usecase.InitUsecase(mongoRepo)

	// Initialize handlers
	// userHandler := handler.NewUserHandler(userUsecase)

	// Set up router
	// r := router.SetupRouter(userHandler)

	// Start the server
	// log.Fatal(r.Run(":8080"))
}
