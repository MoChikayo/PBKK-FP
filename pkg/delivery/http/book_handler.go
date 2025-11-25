package http

import (
    "net/http"
    "strconv"

    "github.com/MoChikayo/PBKK-FP/pkg/domain"
    "github.com/MoChikayo/PBKK-FP/pkg/service"
    "github.com/MoChikayo/PBKK-FP/pkg/utils"
    "github.com/gin-gonic/gin"
)

type BookHandler struct {
    service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
    return &BookHandler{service: s}
}

// GET /api/books
func (h *BookHandler) GetBooks(c *gin.Context) {
    books, err := h.service.ListBooks()
    if err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondSuccess(c, http.StatusOK, books)
}

// GET /api/books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid book ID")
        return
    }

    book, err := h.service.GetBook(uint(id))
    if err != nil {
        utils.RespondError(c, http.StatusNotFound, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusOK, book)
}

// POST /api/books
func (h *BookHandler) CreateBook(c *gin.Context) {
    var input domain.Book

    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    created, err := h.service.CreateBook(input)
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusCreated, created)
}

// PUT /api/books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid book ID")
        return
    }

    var input domain.Book
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    updated, err := h.service.UpdateBook(uint(id), input)
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusOK, updated)
}

// DELETE /api/books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid book ID")
        return
    }

    if err := h.service.DeleteBook(uint(id)); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusOK, gin.H{
        "message": "book deleted",
    })
}
