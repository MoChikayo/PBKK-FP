package models

import (
	"strconv"
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
	ViewLink   string    `json:"view_link"` // add view link for each transaction
}

// Predefined transaction statuses
const (
	StatusBorrowed = "borrowed"
	StatusReturned = "returned"
	StatusOverdue  = "overdue"
)

// CreateTransaction creates a new transaction in the database
func (t *Transaction) CreateTransaction() *Transaction {
	if err := db.Create(&t).Error; err != nil {
		return nil
	}
	return t
}

// GetAllTransactions retrieves all transactions from the database and adds the ViewLink for each
func GetAllTransactions() []Transaction {
	var transactions []Transaction
	db.Preload("Customer").Preload("Book").Find(&transactions)

	// Adding ViewLink for each transaction
	for i, t := range transactions {
		transactions[i].ViewLink = "/transaction/" + strconv.Itoa(int(t.ID))
	}

	return transactions
}

// GetTransactionById retrieves a single transaction by its ID
func GetTransactionById(id int64) (*Transaction, *gorm.DB) {
	var transaction Transaction
	db := db.Preload("Customer").Preload("Book").First(&transaction, id)

	// Add ViewLink for this transaction
	transaction.ViewLink = "/transaction/" + strconv.Itoa(int(transaction.ID))

	return &transaction, db
}
