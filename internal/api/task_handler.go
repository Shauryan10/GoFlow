package api

import (
	"fmt"
	"net/http"

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
	
}
