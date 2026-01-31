package repository

import (
	"database/sql"
	"fenmo-ai-assignment/models"
	"time"
)

// ExpenseRepository handles database operations for expenses
type ExpenseRepository struct {
	db *sql.DB
}

// NewExpenseRepository creates a new expense repository
func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

// Create creates a new expense in the database
func (r *ExpenseRepository) Create(expense *models.Expense) error {
	query := `
		INSERT INTO expenses (id, amount, category, description, date, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		expense.ID,
		expense.Amount,
		expense.Category,
		expense.Description,
		expense.Date,
		expense.CreatedAt,
	)

	return err
}

// GetAll retrieves all expenses from the database
func (r *ExpenseRepository) GetAll() ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, created_at FROM expenses`
	return r.queryExpenses(query)
}

// GetByCategory retrieves expenses filtered by category
func (r *ExpenseRepository) GetByCategory(category string) ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, created_at FROM expenses WHERE category = ?`
	return r.queryExpenses(query, category)
}

// GetAllSortedByDateDesc retrieves all expenses sorted by date descending
func (r *ExpenseRepository) GetAllSortedByDateDesc() ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, created_at 
			  FROM expenses 
			  ORDER BY date DESC, created_at DESC`
	return r.queryExpenses(query)
}

// GetByCategorySortedByDateDesc retrieves expenses filtered by category and sorted by date descending
func (r *ExpenseRepository) GetByCategorySortedByDateDesc(category string) ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, created_at 
			  FROM expenses 
			  WHERE category = ? 
			  ORDER BY date DESC, created_at DESC`
	return r.queryExpenses(query, category)
}

// queryExpenses executes a query and returns expenses
func (r *ExpenseRepository) queryExpenses(query string, args ...interface{}) ([]models.Expense, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		var createdAtStr string

		err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Category,
			&expense.Description,
			&expense.Date,
			&createdAtStr,
		)
		if err != nil {
			return nil, err
		}

		// Parse created_at timestamp - SQLite DATETIME format
		// Try multiple formats
		formats := []string{
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05Z",
			"2006-01-02T15:04:05Z07:00",
			time.RFC3339,
		}
		
		var parsed bool
		for _, format := range formats {
			if expense.CreatedAt, err = time.Parse(format, createdAtStr); err == nil {
				parsed = true
				break
			}
		}
		
		if !parsed {
			expense.CreatedAt = time.Now()
		}

		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}
