package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// تحميل ملف .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables")
	}

	// قراءة القيم من الـ .env
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// بناء الاتصال
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection error: %v", err)
	}

	log.Println("✅ Connected to MySQL successfully")

	// إنشاء الجداول تلقائيًا لو مش موجودة
	// err = DB.AutoMigrate(
	// 	&model.Category{},
	// 	&model.Task{},
	// )
	if err != nil {
		log.Fatalf("❌ AutoMigrate error: %v", err)
	}
}
