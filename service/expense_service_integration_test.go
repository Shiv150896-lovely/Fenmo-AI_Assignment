package service

import (
	"fenmo-ai-assignment/database"
	"fenmo-ai-assignment/models"
	"fenmo-ai-assignment/repository"
	"os"
	"testing"
)

// Integration tests that require a real database
// Run with: go test -tags=integration ./service

func TestExpenseService_CreateExpense_Integration(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Setup test database
	testDBPath := "./test_expenses.db"
	defer os.Remove(testDBPath) // Clean up after test

	err := database.Init(testDBPath)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer database.Close()

	// Create repository and service
	repo := repository.NewExpenseRepository(database.DB)
	service := NewExpenseService(repo)

	tests := []struct {
		name    string
		req     models.CreateExpenseRequest
		wantErr bool
	}{
		{
			name: "valid expense",
			req: models.CreateExpenseRequest{
				Amount:      "100.50",
				Category:    "Food",
				Description: "Lunch",
				Date:        "2024-01-15",
			},
			wantErr: false,
		},
		{
			name: "missing category",
			req: models.CreateExpenseRequest{
				Amount:      "100.50",
				Category:    "",
				Description: "Lunch",
				Date:        "2024-01-15",
			},
			wantErr: true,
		},
		{
			name: "invalid amount - negative",
			req: models.CreateExpenseRequest{
				Amount:      "-10.50",
				Category:    "Food",
				Description: "Lunch",
				Date:        "2024-01-15",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expense, err := service.CreateExpense(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateExpense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if expense == nil {
					t.Errorf("CreateExpense() returned nil expense")
					return
				}
				if expense.ID == "" {
					t.Errorf("CreateExpense() expense.ID is empty")
				}
				if expense.Amount != tt.req.Amount {
					t.Errorf("CreateExpense() expense.Amount = %v, want %v", expense.Amount, tt.req.Amount)
				}
			}
		})
	}
}

func TestExpenseService_GetExpenses_Integration(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Setup test database
	testDBPath := "./test_expenses.db"
	defer os.Remove(testDBPath) // Clean up after test

	err := database.Init(testDBPath)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer database.Close()

	// Create repository and service
	repo := repository.NewExpenseRepository(database.DB)
	service := NewExpenseService(repo)

	// Create test expenses
	_, _ = service.CreateExpense(models.CreateExpenseRequest{
		Amount:      "100.50",
		Category:    "Food",
		Description: "Lunch",
		Date:        "2024-01-15",
	})

	_, _ = service.CreateExpense(models.CreateExpenseRequest{
		Amount:      "50.00",
		Category:    "Transport",
		Description: "Taxi",
		Date:        "2024-01-14",
	})

	_, _ = service.CreateExpense(models.CreateExpenseRequest{
		Amount:      "75.25",
		Category:    "Food",
		Description: "Dinner",
		Date:        "2024-01-16",
	})

	tests := []struct {
		name     string
		category string
		sort     string
		wantLen  int
	}{
		{"get all", "", "", 3},
		{"filter by category", "Food", "", 2},
		{"filter by non-existent category", "NonExistent", "", 0},
		{"sort by date", "", "date_desc", 3},
		{"filter and sort", "Food", "date_desc", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expenses, err := service.GetExpenses(tt.category, tt.sort)
			if err != nil {
				t.Errorf("GetExpenses() error = %v", err)
				return
			}
			if len(expenses) != tt.wantLen {
				t.Errorf("GetExpenses() len = %v, want %v", len(expenses), tt.wantLen)
			}
		})
	}
}
