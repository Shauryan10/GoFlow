package main

import (
	"fmt"

	"github.com/Shauryan10/GoFlow/config"
	"github.com/Shauryan10/GoFlow/internal/api"
	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/queue"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	fmt.Println("Database Host :", cfg.DBHost)
	fmt.Println("Database Port :", cfg.DBPort)
	fmt.Println("Application Port :", cfg.AppPort)

	database.ConnectDatabase(cfg) //connecting to postgres

	queue.InitQueue(100)

	router := gin.Default() //creating gin router

	// Register all API routes
	api.RegisterRoutes(router)

	fmt.Println("🚀 GoFlow API is running on port", cfg.AppPort)
	router.Run(":" + cfg.AppPort)

}
