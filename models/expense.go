package models

import "time"

// Expense represents an expense entry
type Expense struct {
	ID          string    `json:"id" db:"id"`
	Amount      string    `json:"amount" db:"amount"`      // Decimal as string for precision
	Category    string    `json:"category" db:"category"`
	Description string    `json:"description" db:"description"`
	Date        string    `json:"date" db:"date"`          // ISO date format: YYYY-MM-DD
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// CreateExpenseRequest represents the request body for creating an expense
type CreateExpenseRequest struct {
	Amount      string `json:"amount" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
}
