package models

import (
	//"github.com/MoChikayo/PBKK-FP/pkg/config"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	BookID     uint      `json:"book_id"`
	Book       Book      `gorm:"foreignKey:BookID"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
}

// Predefined transaction statuses
const (
	StatusBorrowed = "borrowed"
	StatusReturned = "returned"
	StatusOverdue  = "overdue"
)

func (t *Transaction) CreateTransaction() *Transaction {
	if err := db.Create(&t).Error; err != nil {
		return nil
	}
	return t
}

func GetAllTransactions() []Transaction {
	var transactions []Transaction
	db.Preload("Customer").Preload("Book").Find(&transactions)
	return transactions
}

func GetTransactionById(id int64) (*Transaction, *gorm.DB) {
	var transaction Transaction
	db := db.Preload("Customer").Preload("Book").First(&transaction, id)
	return &transaction, db
}
