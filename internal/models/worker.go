package models

type Worker struct {
	ID            int    `json:"id"`
	Status        string `json:"status"`
	CurrentTaskID uint   `json:"current_task_id"`
}
