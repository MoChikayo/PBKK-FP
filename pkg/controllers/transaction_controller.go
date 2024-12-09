package controllers

import (
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := transaction.CreateTransaction()
	if t == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func GetAllTransactions(c *gin.Context) {
	transactions := models.GetAllTransactions() // Implement this in models
	c.JSON(http.StatusOK, transactions)
}

func GetTransactionById(c *gin.Context) {
	transactionId := c.Param("transactionId")
	ID, err := strconv.ParseInt(transactionId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	transaction, db := models.GetTransactionById(ID) // Implement this in models
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// UpdateTransactionStatus updates the status of a transaction based on the provided data
func UpdateTransactionStatus(c *gin.Context) {
	var updatedData models.Transaction
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId := c.Param("transactionId")
	ID, err := strconv.ParseInt(transactionId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	// Fetch the transaction details explicitly using the "transactions" table
	var transactionDetails models.Transaction
	db := config.GetDB().Table("transactions").Where("id = ?", ID).First(&transactionDetails)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	// Validate the new status
	if updatedData.Status != models.StatusBorrowed && updatedData.Status != models.StatusReturned && updatedData.Status != models.StatusOverdue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}

	// Update transaction status
	transactionDetails.Status = updatedData.Status

	// Save changes
	if err := config.GetDB().Table("transactions").Save(&transactionDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactionDetails)
}
