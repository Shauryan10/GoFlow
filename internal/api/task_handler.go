package api

import (
	"net/http"
	"strconv"

	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/models"
	"github.com/Shauryan10/GoFlow/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := service.CreateTask(req.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := TaskResponse{
		ID:     task.ID,
		Name:   task.Name,
		Status: task.Status,
	}

	c.JSON(http.StatusCreated, response)
}

func GetTasks(c *gin.Context) {

	var tasks []models.Task

	result := database.DB.Find(&tasks)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	responses := []TaskResponse{}

	for _, task := range tasks {

		responses = append(responses, TaskResponse{
			ID:     task.ID,
			Name:   task.Name,
			Status: task.Status,
		})

	}

	c.JSON(http.StatusOK, responses)

}

func GetTaskByID(c *gin.Context) {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})

		return
	}

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})

		return
	}

	response := TaskResponse{
		ID:     task.ID,
		Name:   task.Name,
		Status: task.Status,
	}

	c.JSON(http.StatusOK, response)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	taskID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var req UpdateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	if req.Name != "" {
		task.Name = req.Name
	}

	if req.Status != "" {
		task.Status = req.Status
	}

	database.DB.Save(&task)

	response := TaskResponse{
		ID:     task.ID,
		Name:   task.Name,
		Status: task.Status,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	taskID, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task

	result := database.DB.First(&task, taskID)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	database.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
