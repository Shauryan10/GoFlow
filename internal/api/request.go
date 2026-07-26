package api

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
