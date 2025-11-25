package models

import (
	//"github.com/MoChikayo/PBKK-FP/pkg/config"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Phone string `json:"phone"`
}

func (u *Customer) CreateCustomer() *Customer {
	if err := db.Create(&u).Error; err != nil {
		return nil
	}
	return u
}

func GetAllCustomers() []Customer {
	var customers []Customer
	db.Find(&customers)
	return customers
}

func GetCustomerById(id int64) (*Customer, *gorm.DB) {
	var customer Customer
	db := db.First(&customer, id)
	return &customer, db
}
