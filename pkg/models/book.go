package models

import (
	"github.com/MoChikayo/PBKK-FP/pkg/config"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

//var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Customer{}, &Transaction{})
}

// func (b *Book) CreateBook() *Book {
// 	db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

func (b *Book) CreateBook() *Book {
	if err := db.Create(&b).Error; err != nil {
		return nil
	}
	return b
}

// func GetAllBooks() ([]Book, error) {
// 	var books []Book
// 	err := db.Find(&books).Error
// 	return books, err
// }

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.First(&book, id) // This returns a *gorm.DB object as the second value
	return &book, db
}

func DeleteBook(ID int64) error {
	var book Book
	if err := db.Where("ID = ?", ID).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

// func GetBookById(id int64) (*Book, *gorm.DB) {
// 	var book Book
// 	db := db.First(&book, id) // This returns a *gorm.DB object as the second value
// 	return &book, db
// }

// func DeleteBook(ID int64) Book {
// 	var book Book
// 	db.Where("ID = ?", ID).Delete(book)
// 	return book
// }
