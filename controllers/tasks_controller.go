package controllers

import (
	"database/sql"
	"strconv"
	"taskmanager/database"
	"taskmanager/helpers"
	models "taskmanager/model"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, done,category_id FROM tasks")
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}
	defer rows.Close()

	var tasksList []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CategoryId); err != nil {
			helpers.Respond(c, false, nil, err.Error())
			return
		}
		tasksList = append(tasksList, t)
	}

	helpers.Respond(c, true, tasksList, "Tasks retrieved successfully")
}

func CreateTask(c *gin.Context) {
	var t models.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		helpers.Respond(c, false, nil, "Invalid JSON format")
		return
	}

	res, err := database.DB.Exec("INSERT INTO tasks (title, description, done, category_id) VALUES (?, ?, ?, ?)", t.Title, t.Description, t.Done, t.CategoryId)
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}
	id, _ := res.LastInsertId()
	t.ID = int(id)

	helpers.Respond(c, true, t, "Task created successfully")
}

func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t models.Task
	err := database.DB.QueryRow("SELECT id, title, description, done, category_id FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Description, &t.Done, t.CategoryId)
	if err == sql.ErrNoRows {
		helpers.Respond(c, false, nil, "Task not found")
		return
	} else if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, t, "Task retrieved successfully")
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t models.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		helpers.Respond(c, false, nil, "Invalid JSON")
		return
	}

	_, err := database.DB.Exec("UPDATE tasks SET title=?, description=?, done=? WHERE id=?", t.Title, t.Description, t.Done, id)
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}

	helpers.Respond(c, true, t, "Task updated successfully")
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}
	helpers.Respond(c, true, nil, "Task deleted successfully")
}

func CompleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DB.Exec("UPDATE tasks SET done=true WHERE id=?", id)
	if err != nil {
		helpers.Respond(c, false, nil, err.Error())
		return
	}
	helpers.Respond(c, true, gin.H{"id": id, "done": true}, "Task marked as complete")
}
