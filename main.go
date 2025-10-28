package main

import (
	"fmt"
	"taskmanager/database"
	"taskmanager/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Connect()
	router := gin.Default()
	routes.RegisterTaskRoutes(router)
	routes.RegisterCategoriesRoutes(router)

	fmt.Println("🚀 Server running on http://localhost:8080")
	router.Run(":8080")
}
