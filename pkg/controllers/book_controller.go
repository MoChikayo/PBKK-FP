package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/models"
	"github.com/MoChikayo/PBKK-FP/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	//"gorm.io/gorm"
)

var NewBook models.Book

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	newBooks, err := models.GetAllBooks()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	res, _ := json.Marshal(newBooks)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // No error return value here
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil { // Check for errors from the *gorm.DB
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	//w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook handles the deletion of a book by its ID
func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Fetch the book details to ensure it exists before deleting
	var bookDetails models.Book
	db := config.GetDB().Where("id = ?", ID).First(&bookDetails)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	if bookDetails.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Soft delete the book
	if err := config.GetDB().Delete(&bookDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBook handles the updating of a book's details by its ID
func UpdateBook(c *gin.Context) {
	var updateBook models.Book
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Fetch the book details explicitly using the "books" table
	var bookDetails models.Book
	db := config.GetDB().Table("books").Where("id = ?", ID).First(&bookDetails)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save changes
	if err := config.GetDB().Table("books").Where("id = ?", ID).Save(&bookDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookDetails)
}

// func DeleteBook(c *gin.Context) {
// 	bookId := c.Param("bookId")
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
// 		return
// 	}
// 	book := models.DeleteBook(ID)
// 	c.JSON(http.StatusOK, book)
// }

// func UpdateBook(c *gin.Context) {
// 	var updateBook models.Book
// 	if err := c.ShouldBindJSON(&updateBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	bookId := c.Param("bookId")
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
// 		return
// 	}

// 	bookDetails, db := models.GetBookById(ID)
// 	if db.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
// 		return
// 	}

// 	// Update fields
// 	if updateBook.Name != "" {
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != "" {
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != "" {
// 		bookDetails.Publication = updateBook.Publication
// 	}

// 	if err := db.Save(&bookDetails).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, bookDetails)
// }
