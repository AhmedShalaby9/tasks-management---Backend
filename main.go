package main

import (
	"net/http"
	"strconv"
	"taskmanager/tasks"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.PUT("/tasks/:id/complete", completeTask) // مثال لــ multi-level path

}

// getTasks - GET /tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks.Tasks)
}

// createTask - POST /tasks
func createTask(c *gin.Context) {
	var newTask tasks.Task
	if err := c.ShouldBindJSON(&newTask); err != nil { // bind JSON آمن
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = tasks.NextID
	tasks.NextID++
	tasks.Tasks = append(tasks.Tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// getTaskByID - GET /tasks/:id
func getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for _, t := range tasks.Tasks {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

// updateTask - PUT /tasks/:id
func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updated tasks.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Title = updated.Title
			tasks.Tasks[i].Description = updated.Description
			tasks.Tasks[i].Done = updated.Done
			c.JSON(http.StatusOK, tasks.Tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

// deleteTask - DELETE /tasks/:id
func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

// completeTask - PUT /tasks/:id/complete
func completeTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Done = true
			c.JSON(http.StatusOK, tasks.Tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
