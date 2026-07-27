package queue

import "github.com/Shauryan10/GoFlow/internal/models"

var TaskQueue chan models.Task

func InitQueue(buffer int) {
	TaskQueue = make(chan models.Task, buffer)
}
