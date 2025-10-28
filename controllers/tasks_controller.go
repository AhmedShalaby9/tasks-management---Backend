package controllers

import (
	"strconv"
	"taskmanager/database"
	"taskmanager/helpers"
	models "taskmanager/model"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	var res = database.DB.Find(&tasks)

	if res.Error != nil {
		helpers.Respond(c, false, nil, res.Error.Error())
		return
	}
	helpers.Respond(c, true, tasks, "Tasks retrieved successfully")

}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		helpers.Respond(c, false, nil, "Invalid JSON format")
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, task, "Task created successfully")
}

// 🟢 Get task by ID
func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		helpers.Respond(c, false, nil, "Task not found")
		return
	}

	helpers.Respond(c, true, task, "Task retrieved successfully")
}

// 🟢 Update task
func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		helpers.Respond(c, false, nil, "Task not found")
		return
	}

	var updated models.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		helpers.Respond(c, false, nil, "Invalid JSON")
		return
	}

	task.Title = updated.Title
	task.Description = updated.Description
	task.Done = updated.Done
	task.CategoryId = updated.CategoryId

	if err := database.DB.Save(&task).Error; err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, task, "Task updated successfully")
}

// 🟢 Delete task
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, nil, "Task deleted successfully")
}

// 🟢 Mark task as complete
func CompleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		helpers.Respond(c, false, nil, "Task not found")
		return
	}

	task.Done = true
	if err := database.DB.Save(&task).Error; err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, task, "Task marked as complete")
}
