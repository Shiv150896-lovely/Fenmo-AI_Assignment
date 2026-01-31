package utils

import "testing"

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  string
		wantErr bool
	}{
		{"valid amount", "100.50", false},
		{"valid integer", "100", false},
		{"valid decimal", "0.99", false},
		{"empty string", "", true},
		{"negative amount", "-10.50", true},
		{"invalid format", "abc", true},
		{"whitespace", "  100.50  ", false},
		{"zero", "0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAmount(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAmount(%q) error = %v, wantErr %v", tt.amount, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{"valid date", "2024-01-15", false},
		{"valid date with leading zeros", "2024-01-05", false},
		{"empty string", "", true},
		{"invalid format - missing dashes", "20240115", true},
		{"invalid format - wrong separator", "2024/01/15", true},
		{"invalid format - short year", "24-01-15", true},
		{"invalid format - short month", "2024-1-15", true},
		{"invalid format - short day", "2024-01-5", true},
		{"invalid format - non-numeric", "abcd-01-15", true},
		{"whitespace", "  2024-01-15  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDate(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDate(%q) error = %v, wantErr %v", tt.date, err, tt.wantErr)
			}
		})
	}
}
