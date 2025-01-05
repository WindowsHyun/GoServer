package main

import (
	"GoServer/router"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
