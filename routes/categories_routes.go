package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoriesRoutes(router *gin.Engine) {
	router.GET("/categories", controllers.GetCategories)
}
