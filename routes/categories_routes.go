package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoriesRoutes(router *gin.Engine) {
	router.GET("/categories", controllers.GetCategories)
	router.POST("/categories", controllers.CreateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)

}
