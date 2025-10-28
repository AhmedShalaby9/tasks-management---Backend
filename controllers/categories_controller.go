package controllers

import (
	"strconv"
	"taskmanager/database"
	"taskmanager/helpers"
	models "taskmanager/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category

	res := database.DB.Find(&categories)
	if res.Error != nil {
		helpers.Respond(c, false, nil, res.Error.Error())
		return
	}
	helpers.Respond(c, true, categories, "Categories retrieved successfully")

}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	res := database.DB.Create(&category)
	if res.Error != nil {
		helpers.Respond(c, false, nil, res.Error.Error())
		return
	}

	helpers.Respond(c, true, category, "Category created successfully")
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := database.DB.Delete(&models.Category{}, id)
	if res.Error != nil {
		helpers.Respond(c, false, nil, res.Error.Error())
		return
	}

	helpers.Respond(c, true, nil, "Category deleted successfully")
}
