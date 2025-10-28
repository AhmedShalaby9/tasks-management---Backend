package controllers

import (
	"taskmanager/database"
	"taskmanager/helpers"
	models "taskmanager/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title FROM categories")
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}
	defer rows.Close()

	var categoriesList []models.Category
	for rows.Next() {
		var t models.Category
		if err := rows.Scan(&t.ID, &t.Title); err != nil {
			helpers.Respond(c, false, nil, err.Error())
			return
		}
		categoriesList = append(categoriesList, t)
	}

	helpers.Respond(c, true, categoriesList, "Categories retrieved successfully")
}
