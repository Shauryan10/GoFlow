package service

import (
	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/models"
	"github.com/Shauryan10/GoFlow/internal/queue"
)

func CreateTask(name string) (*models.Task, error) {

	task := models.Task{
		Name:   name,
		Status: "PENDING",
	}

	result := database.DB.Create(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	// Send task to worker queue
	queue.TaskQueue <- task

	return &task, nil
}
