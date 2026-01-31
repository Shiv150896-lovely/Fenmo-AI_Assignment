package utils

import (
	"regexp"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	// UUID v4 format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	tests := []struct {
		name string
	}{
		{"generate uuid 1"},
		{"generate uuid 2"},
		{"generate uuid 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid := GenerateUUID()
			if !uuidPattern.MatchString(uuid) {
				t.Errorf("GenerateUUID() = %v, does not match UUID v4 pattern", uuid)
			}
			if len(uuid) != 36 {
				t.Errorf("GenerateUUID() = %v, length = %d, want 36", uuid, len(uuid))
			}
		})
	}

	// Test uniqueness
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUID() generated duplicate UUIDs: %v", uuid1)
	}
}
