package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connect() {
	dsn := "root:2010961523@tcp(127.0.0.1:3306)/taskmanager_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database ping error:", err)
	}

	fmt.Println("✅ Connected to MySQL successfully!")
	DB = db
}
