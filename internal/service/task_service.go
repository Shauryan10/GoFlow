package service

import (
	"time"

	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/metrics"
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
	metrics.QueueLength.Set(float64(len(queue.TaskQueue)))
	
	return &task, nil
}

func CompleteTask(taskID uint) error {

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {
		return result.Error
	}

	task.Status = "COMPLETED"

	task.Progress = 100

	now := time.Now()

	task.CompletedAt = &now

	task.Result = "Task processed successfully"

	result = database.DB.Save(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func StartTask(taskID uint) error {

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {
		return result.Error
	}

	task.Status = "RUNNING"

	result = database.DB.Save(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTaskProgress(taskID uint, progress uint) error {

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {
		return result.Error
	}

	task.Progress = progress

	result = database.DB.Save(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IncrementRetry(taskID uint) error {

	var task models.Task

	if err := database.DB.First(&task, taskID).Error; err != nil {
		return err
	}

	task.RetryCount++

	return database.DB.Save(&task).Error
}

func FailTask(taskID uint, message string) error {

	var task models.Task

	if err := database.DB.First(&task, taskID).Error; err != nil {
		return err
	}

	now := time.Now()

	task.Status = "FAILED"
	task.Error = message
	task.CompletedAt = &now

	return database.DB.Save(&task).Error
}
