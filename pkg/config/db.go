package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	if db != nil {
		return db
	}

	// Change these to your MySQL login
	user := "root"
	pass := "" // your password
	host := "127.0.0.1"
	port := "3306"
	name := "bookstore" // database name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect MySQL: %v", err)
	}

	db = d
	return db
}

func GetDB() *gorm.DB {
	return db
}
