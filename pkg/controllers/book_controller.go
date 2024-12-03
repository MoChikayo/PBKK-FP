package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/models"
	"github.com/MoChikayo/PBKK-FP/pkg/utils"
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
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
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if err := db.Save(&bookDetails).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content_Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
