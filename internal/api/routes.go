package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", HealthCheck)

	router.POST("/tasks", CreateTask)
	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTaskByID)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)
}
