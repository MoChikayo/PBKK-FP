package models

import (
	"github.com/MoChikayo/PBKK-FP/pkg/config"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
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

// func GetBookById(id int64) (*Book, error) {
// 	var book Book
// 	err := db.First(&book, id).Error
// 	return &book, err
// }

// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID = ?", Id).Find(&getBook)
// 	return &getBook, db
// }

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.First(&book, id) // This returns a *gorm.DB object as the second value
	return &book, db
}

// func DeleteBook(id int64) error {
// 	return db.Delete(&Book{}, id).Error
// }

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
