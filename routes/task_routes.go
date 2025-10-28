package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}
