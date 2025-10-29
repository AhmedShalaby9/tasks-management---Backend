package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:2010961523@tcp(127.0.0.1:3306)/taskmanager_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Database connection error:", err)
	}

	if err != nil {
		log.Fatal("❌ Database ping error:", err)
	}
	fmt.Println("✅ Connected to MySQL successfully!")
	DB = db
}
