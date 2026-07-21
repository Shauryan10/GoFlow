package database

import (
	"fmt"
	"log"

	"github.com/Shauryan10/GoFlow/config"
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
		log.Fatal("Failed to connect database")
	}

	DB = db

	fmt.Println("Connected to PostgreSQL")
}
