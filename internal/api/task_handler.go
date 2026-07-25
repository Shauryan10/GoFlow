package api

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/Shauryan10/GoFlow/internal/databases"
	"github.com/Shauryan10/GoFlow/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("Received:", req.Name)

	task := models.Task{
		Name:   req.Name,
		Status: "PENDING",
	}

	result := database.DB.Create(&task)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
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
