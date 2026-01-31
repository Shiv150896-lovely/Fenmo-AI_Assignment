package handler

import (
	"fenmo-ai-assignment/models"
	"fenmo-ai-assignment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExpenseHandler handles HTTP requests for expenses
type ExpenseHandler struct {
	service *service.ExpenseService
}

// NewExpenseHandler creates a new expense handler
func NewExpenseHandler(service *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

// CreateExpense handles POST /expenses
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req models.CreateExpenseRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Create expense
	expense, err := h.service.CreateExpense(req)
	if err != nil {
		// Check if it's a validation error
		if validationErr, ok := err.(*service.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + validationErr.Message})
			return
		}

		// Database or other error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

// GetExpenses handles GET /expenses
func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	// Get query parameters
	category := c.Query("category")
	sort := c.Query("sort")

	// Get expenses
	expenses, err := h.service.GetExpenses(category, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return empty array if no expenses
	if expenses == nil {
		expenses = []models.Expense{}
	}

	c.JSON(http.StatusOK, expenses)
}
