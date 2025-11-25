package domain

import "gorm.io/gorm"

type Customer struct {
    gorm.Model
    Name        string `json:"name"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
    Address     string `json:"address"`
}
