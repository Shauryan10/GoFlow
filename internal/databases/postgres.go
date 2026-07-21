package database

import (
	"fmt"
	"log"

	"github.com/Shauryan10/GoFlow/config"
	"github.com/Shauryan10/GoFlow/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	DB = db
	err = DB.AutoMigrate(&models.Task{})

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Database connected successfully!")
	fmt.Println("Tasks table migrated successfully!")

}
