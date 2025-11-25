package domain

import (
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
    ViewLink   string    `json:"view_link"`
}

// Predefined statuses
const (
    StatusBorrowed = "borrowed"
    StatusReturned = "returned"
    StatusOverdue  = "overdue"
)
