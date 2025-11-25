package controllers

import (
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := customer.CreateCustomer()
	if u == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Customer"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func GetAllCustomers(c *gin.Context) {
	customers := models.GetAllCustomers() // Implement this in models
	c.JSON(http.StatusOK, customers)
}

func GetCustomerById(c *gin.Context) {
	customerId := c.Param("customerId")
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}
	customer, db := models.GetCustomerById(ID) // Implement this in models
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer updates a customer's details based on the provided data
func UpdateCustomer(c *gin.Context) {
	var updatedData models.Customer
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerId := c.Param("customerId")
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	// Fetch the customer details explicitly using the "customers" table
	var customerDetails models.Customer
	db := config.GetDB().Table("customers").Where("id = ?", ID).First(&customerDetails)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	// Update fields
	if updatedData.Name != "" {
		customerDetails.Name = updatedData.Name
	}
	if updatedData.Email != "" {
		customerDetails.Email = updatedData.Email
	}
	if updatedData.Phone != "" {
		customerDetails.Phone = updatedData.Phone
	}

	// Save changes
	if err := config.GetDB().Table("customers").Save(&customerDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customerDetails)
}

// DeleteCustomer deletes a customer by ID
func DeleteCustomer(c *gin.Context) {
	customerId := c.Param("customerId")
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
		return
	}

	// Fetch the customer details explicitly using the "customers" table
	var customerDetails models.Customer
	db := config.GetDB().Table("customers").Where("id = ?", ID).First(&customerDetails)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	// Soft delete by setting `deleted_at` field to the current timestamp
	if err := config.GetDB().Table("customers").Model(&customerDetails).UpdateColumn("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// // Update Customer
// func UpdateCustomer(c *gin.Context) {
// 	var updatedData models.Customer
// 	if err := c.ShouldBindJSON(&updatedData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	customerId := c.Param("customerId")
// 	ID, err := strconv.ParseInt(customerId, 0, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
// 		return
// 	}

// 	customerDetails, db := models.GetCustomerById(int64(ID))
// 	if db.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
// 		return
// 	}

// 	if updatedData.Name != "" {
// 		customerDetails.Name = updatedData.Name
// 	}
// 	if updatedData.Email != "" {
// 		customerDetails.Email = updatedData.Email
// 	}
// 	if err := db.Save(&customerDetails).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, customerDetails)
// }

// // Delete Customer
// func DeleteCustomer(c *gin.Context) {
// 	customerId := c.Param("customerId")
// 	ID, err := strconv.ParseInt(customerId, 0, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Customer ID"})
// 		return
// 	}

// 	customerDetails, db := models.GetCustomerById(ID)
// 	if db.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
// 		return
// 	}

// 	if err := db.Delete(&customerDetails).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
// }
