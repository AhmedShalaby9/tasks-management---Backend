package controllers

import (
	"taskmanager/database"
	"taskmanager/helpers"
	models "taskmanager/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, categories, "Categories retrieved successfully")
}
