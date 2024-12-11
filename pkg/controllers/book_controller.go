package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/config"
	"github.com/MoChikayo/PBKK-FP/pkg/models"
	"github.com/MoChikayo/PBKK-FP/pkg/utils"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

var NewBook models.Book

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	newBooks := models.GetAllBooks() // No error return value here
// 	res, err := json.Marshal(newBooks)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func GetBook(c *gin.Context) {
	// Get all books
	newBooks := models.GetAllBooks()

	// Determine response format
	switch format := utils.GetFormat(c); format {
	case utils.FormatHTML:
		c.HTML(http.StatusOK, "books.list.html", gin.H{
			"title": "All Books",
			"books": newBooks,
		})
	case utils.FormatJSON:
		c.JSON(http.StatusOK, gin.H{"books": newBooks})
	}
}

func GetBookById(c *gin.Context) {
	// Extract bookId from the URL
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Fetch book details
	bookDetails, db := models.GetBookById(ID)
	if db.Error != nil { // Handle database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	// Determine response format
	switch utils.GetFormat(c) {
	case utils.FormatHTML:
		c.HTML(http.StatusOK, "books.view.html", gin.H{
			"title": "Book Details",
			"book":  bookDetails,
		})
	case utils.FormatJSON:
		c.JSON(http.StatusOK, gin.H{"book": bookDetails})
	}
}

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}

// 	bookDetails, db := models.GetBookById(ID)
// 	if db.Error != nil { // Check for errors from the *gorm.DB
// 		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	res, err := json.Marshal(bookDetails)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	b := CreateBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateBook(c *gin.Context) {
	var newBook models.Book

	// Determine the content type of the request
	if c.ContentType() == "application/json" {
		// Bind JSON input
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}
	} else {
		// Bind form data (such as when submitting an HTML form)
		if err := c.ShouldBind(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form input"})
			return
		}
	}

	// Create the book in the database
	createdBook := newBook.CreateBook()

	// Render the appropriate response based on the desired format
	switch utils.GetFormat(c) {
	case utils.FormatHTML:
		c.HTML(http.StatusOK, "books.create.html", gin.H{
			"title": "Create New Book",
			"book":  createdBook,
		})
	case utils.FormatJSON:
		c.JSON(http.StatusOK, gin.H{"book": createdBook})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
	}
}

// DeleteBook handles the deletion of a book by its ID
func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

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

	if err := config.GetDB().Delete(&bookDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

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

	if err := config.GetDB().Table("books").Where("id = ?", ID).Save(&bookDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch utils.GetFormat(c) {
	case utils.FormatHTML:
		c.HTML(http.StatusOK, "books.update.html", gin.H{
			"title": "Book Updated",
			"book":  bookDetails,
		})
	case utils.FormatJSON:
		c.JSON(http.StatusOK, gin.H{"book": bookDetails})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
	}

	//c.JSON(http.StatusOK, bookDetails)
}

// func UpdateBook(c *gin.Context) {
// 	var updateBook models.Book
// 	if err := c.ShouldBindJSON(&updateBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	bookId := c.Param("bookId")
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}

// 	// Fetch the book details explicitly using the "books" table
// 	var bookDetails models.Book
// 	db := config.GetDB().Table("books").Where("id = ?", ID).First(&bookDetails)
// 	if db.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
// 		return
// 	}
// 	if updateBook.Name != "" {
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != "" {
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != "" {
// 		bookDetails.Publication = updateBook.Publication
// 	}

// 	// Save changes
// 	if err := config.GetDB().Table("books").Where("id = ?", ID).Save(&bookDetails).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, bookDetails)
// }

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
