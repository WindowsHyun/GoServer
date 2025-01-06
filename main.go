package main

import (
	"GoServer/config"
	"GoServer/database/mongo"
	"GoServer/handler"
	"GoServer/repository"
	"GoServer/router"
	"GoServer/usecase"
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	cf := config.InitConfig()
	if cf == nil {
		fmt.Printf("init config failed\n")
		return
	}

	if err := mongo.InitializeMongo(ctx, nil); err != nil {
		log.Fatalf("Failed to initialize MongoDB clients: %v", err)
	}
	defer mongo.CloseMongo(ctx)

	// Initialize repositories
	repos := repository.NewRepositories()

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(repos.User)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)

	// Set up router
	r := router.SetupRouter(userHandler)

	// Start the server
	log.Fatal(r.Run(":8080"))
}
