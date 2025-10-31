package database

import (
	"fmt"
	"log"

	// models "taskmanager/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := "ballast.proxy.rlwy.net"
	port := "54178"
	user := "root"
	password := "NgjcxUaphKqtlbmXkxfrEqswRVUSioIN"
	dbname := "railway"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection error: %v", err)
	}
	// err = DB.AutoMigrate(
	// 	&models.Category{},
	// 	&models.Task{},
	// )
	log.Println("✅ Connected to Railway MySQL successfully")
}
