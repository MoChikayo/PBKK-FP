package http

import (
	"net/http"
	"strconv"

	"github.com/MoChikayo/PBKK-FP/pkg/domain"
	"github.com/MoChikayo/PBKK-FP/pkg/service"
	"github.com/MoChikayo/PBKK-FP/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(s service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

// GET /api/customers
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := h.service.ListCustomers()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondSuccess(c, http.StatusOK, customers)
}

// GET /api/customers/:id
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "invalid customer ID")
		return
	}

	customer, err := h.service.GetCustomer(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, customer)
}

// POST /api/customers
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var input domain.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.service.CreateCustomer(input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, created)
}

// PUT /api/customers/:id
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "invalid customer ID")
		return
	}

	var input domain.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.service.UpdateCustomer(uint(id), input)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, updated)
}

// DELETE /api/customers/:id
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "invalid customer ID")
		return
	}

	if err := h.service.DeleteCustomer(uint(id)); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}
