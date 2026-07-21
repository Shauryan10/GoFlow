package main

import (
	"fmt"

	"github.com/Shauryan10/GoFlow/config"
	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	database.ConnectDatabase(cfg)

	fmt.Println("Database Host :", cfg.DBHost)
	fmt.Println("Database Port :", cfg.DBPort)
	fmt.Println("Application Port :", cfg.AppPort)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GoFlow API Running",
		})
	})

	router.Run(":" + cfg.AppPort)
}
