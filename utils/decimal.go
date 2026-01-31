package utils

import (
	"errors"
	"strconv"
	"strings"
)

// ValidateAmount validates that the amount string is a valid positive decimal number
func ValidateAmount(amount string) error {
	if amount == "" {
		return errors.New("amount cannot be empty")
	}

	// Remove any whitespace
	amount = strings.TrimSpace(amount)

	// Parse as float64 to validate it's a number
	value, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return errors.New("amount must be a valid number")
	}

	// Check if negative
	if value < 0 {
		return errors.New("amount must be positive")
	}

	return nil
}

// ValidateDate validates that the date string is in ISO format (YYYY-MM-DD)
func ValidateDate(date string) error {
	if date == "" {
		return errors.New("date cannot be empty")
	}

	date = strings.TrimSpace(date)

	// Basic format check: YYYY-MM-DD
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return errors.New("date must be in YYYY-MM-DD format")
	}

	// Try to parse each part as integers
	year := parts[0]
	month := parts[1]
	day := parts[2]

	if len(year) != 4 || len(month) != 2 || len(day) != 2 {
		return errors.New("date must be in YYYY-MM-DD format")
	}

	// Validate numeric values
	if _, err := strconv.Atoi(year); err != nil {
		return errors.New("date must be in YYYY-MM-DD format")
	}
	if _, err := strconv.Atoi(month); err != nil {
		return errors.New("date must be in YYYY-MM-DD format")
	}
	if _, err := strconv.Atoi(day); err != nil {
		return errors.New("date must be in YYYY-MM-DD format")
	}

	return nil
}
