package service

import (
	"fenmo-ai-assignment/models"
	"fenmo-ai-assignment/repository"
	"fenmo-ai-assignment/utils"
	"time"
)

// ExpenseService handles business logic for expenses
type ExpenseService struct {
	repo *repository.ExpenseRepository
}

// NewExpenseService creates a new expense service
func NewExpenseService(repo *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

// CreateExpense creates a new expense with validation
func (s *ExpenseService) CreateExpense(req models.CreateExpenseRequest) (*models.Expense, error) {
	// Validate amount
	if err := utils.ValidateAmount(req.Amount); err != nil {
		return nil, err
	}

	// Validate date
	if err := utils.ValidateDate(req.Date); err != nil {
		return nil, err
	}

	// Validate category and description are not empty
	if req.Category == "" {
		return nil, &ValidationError{Message: "category is required"}
	}
	if req.Description == "" {
		return nil, &ValidationError{Message: "description is required"}
	}

	// Create expense model
	expense := &models.Expense{
		ID:          utils.GenerateUUID(),
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		Date:        req.Date,
		CreatedAt:   time.Now(),
	}

	// Save to database
	if err := s.repo.Create(expense); err != nil {
		return nil, err
	}

	return expense, nil
}

// GetExpenses retrieves expenses with optional filtering and sorting
func (s *ExpenseService) GetExpenses(category string, sort string) ([]models.Expense, error) {
	var expenses []models.Expense
	var err error

	// Determine which repository method to call based on filters
	if category != "" && sort == "date_desc" {
		expenses, err = s.repo.GetByCategorySortedByDateDesc(category)
	} else if category != "" {
		expenses, err = s.repo.GetByCategory(category)
	} else if sort == "date_desc" {
		expenses, err = s.repo.GetAllSortedByDateDesc()
	} else {
		expenses, err = s.repo.GetAll()
	}

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
