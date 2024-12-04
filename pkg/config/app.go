package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "database.db"

	// Open a SQLite connection
	d, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}
	db = d
	log.Println("Connected to the SQLite database")
}

func GetDB() *gorm.DB {
	return db
}
