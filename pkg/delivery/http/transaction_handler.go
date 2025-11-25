package http

import (
    "net/http"
    "strconv"

    "github.com/MoChikayo/PBKK-FP/pkg/domain"
    "github.com/MoChikayo/PBKK-FP/pkg/service"
    "github.com/MoChikayo/PBKK-FP/pkg/utils"
    "github.com/gin-gonic/gin"
)

type TransactionHandler struct {
    service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
    return &TransactionHandler{service: s}
}

// GET /api/transactions
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
    txs, err := h.service.ListTransactions()
    if err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondSuccess(c, http.StatusOK, txs)
}

// GET /api/transactions/:id
func (h *TransactionHandler) GetTransaction(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid transaction ID")
        return
    }

    tx, err := h.service.GetTransaction(uint(id))
    if err != nil {
        utils.RespondError(c, http.StatusNotFound, err.Error())
        return
    }
    utils.RespondSuccess(c, http.StatusOK, tx)
}

// POST /api/transactions/borrow
func (h *TransactionHandler) Borrow(c *gin.Context) {
    var input domain.Transaction
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    created, err := h.service.BorrowBook(input)
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusCreated, created)
}

// PUT /api/transactions/return/:id
func (h *TransactionHandler) Return(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid transaction ID")
        return
    }

    updated, err := h.service.ReturnBook(uint(id))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusOK, updated)
}

// DELETE /api/transactions/:id
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.RespondError(c, http.StatusBadRequest, "invalid transaction ID")
        return
    }

    if err := h.service.DeleteTransaction(uint(id)); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    utils.RespondSuccess(c, http.StatusOK, gin.H{
        "message": "transaction deleted",
    })
}
