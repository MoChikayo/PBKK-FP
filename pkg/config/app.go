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

// ResetDatabase clears all data from the database
func ResetDatabaseEndpoint() error {
	// Use PRAGMA statements to disable/enable foreign key checks
	if err := db.Exec("PRAGMA foreign_keys = OFF").Error; err != nil {
		return err
	}

	// List your table names to truncate
	tables := []string{"customers", "books", "transactions"}
	for _, table := range tables {
		if err := db.Exec("DELETE FROM " + table).Error; err != nil {
			return err
		}
	}

	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return err
	}

	return nil
}
