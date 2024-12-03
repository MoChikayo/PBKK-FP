package models

import (
	"github.com/MoChikayo/PBKK-FP/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	if err := db.Create(&b).Error; err != nil {
		return nil
	}
	return b
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	return books, err
}

func GetBookById(id int64) (*Book, error) {
	var book Book
	err := db.First(&book, id).Error
	return &book, err
}

func DeleteBook(id int64) error {
	return db.Delete(&Book{}, id).Error
}
