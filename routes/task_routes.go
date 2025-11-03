package routes

import (
	"taskmanager/controllers"
	"taskmanager/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {

	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(middlewares.AuthMiddleware())

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}
