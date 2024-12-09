package models

import (
	"time"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Transaction struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID"`
	BookID     uint      `json:"book_id"`
	Book       Book      `gorm:"foreignKey:BookID"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &User{}, &Transaction{})
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

func (u *User) CreateUser() *User {
	if err := db.Create(&u).Error; err != nil {
		return nil
	}
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var user User
	db := db.First(&user, id)
	return &user, db
}

func (t *Transaction) CreateTransaction() *Transaction {
	if err := db.Create(&t).Error; err != nil {
		return nil
	}
	return t
}

func GetAllTransactions() []Transaction {
	var transactions []Transaction
	db.Preload("User").Preload("Book").Find(&transactions)
	return transactions
}

func GetTransactionById(id int64) (*Transaction, *gorm.DB) {
	var transaction Transaction
	db := db.Preload("User").Preload("Book").First(&transaction, id)
	return &transaction, db
}
