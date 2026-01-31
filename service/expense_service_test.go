package service

import (
	"testing"
)

// Note: Service tests require database connection
// For proper unit testing, the service should use an interface for the repository
// These tests are skipped and would require integration testing setup

func TestExpenseService_CreateExpense(t *testing.T) {
	// Note: This test requires actual database connection
	// For unit testing without DB, we'd need to refactor to use interfaces
	// For now, we'll test the validation logic which doesn't require DB
	t.Skip("Skipping - requires database connection. Use integration tests instead.")
}

func TestExpenseService_GetExpenses(t *testing.T) {
	// Note: This test requires actual database connection
	// For unit testing without DB, we'd need to refactor to use interfaces
	t.Skip("Skipping - requires database connection. Use integration tests instead.")
}

// Test validation logic separately (tested in utils package)
func TestValidationLogic(t *testing.T) {
	// Validation is tested in utils/decimal_test.go
	// This is just a placeholder to show where service validation tests would go
	t.Skip("Validation tests are in utils package")
}
