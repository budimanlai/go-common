package gocommon

import (
	"testing"
	"time"
)

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Already normalized
		{"628123456789", "628123456789"},
		// With leading zero
		{"08123456789", "628123456789"},
		// With leading 8
		{"8123456789", "628123456789"},
		// With non-digit characters
		{"+62 812-3456-789", "628123456789"},
		{"(0812) 3456-789", "628123456789"},
		// Already normalized with non-digit characters
		{"62-812-3456-789", "628123456789"},
		// Random string, no digits
		{"abc", ""},
		// Empty string
		{"", ""},
		// Only non-digit prefix
		{"+0812-3456-789", "628123456789"},
		// Number with country code and spaces
		{"62 812 3456 789", "628123456789"},
	}

	for _, tt := range tests {
		result := NormalizePhoneNumber(tt.input)
		if result != tt.expected {
			t.Errorf("NormalizePhoneNumber(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestGenerateTransactionID_Uniqueness(t *testing.T) {
	ids := make(map[string]struct{})
	for i := 0; i < 1000; i++ {
		id := GenerateTransactionID()
		if _, exists := ids[id]; exists {
			t.Errorf("Duplicate transaction ID generated: %s", id)
		}
		ids[id] = struct{}{}
	}
}

func TestGenerateTransactionID_Format(t *testing.T) {
	id := GenerateTransactionID()
	if len(id) != 20 {
		t.Errorf("Expected length 16, got %d", len(id))
	}
	for i, r := range id {
		if r < '0' || r > '9' {
			t.Errorf("Non-digit character at position %d: %q", i, r)
		}
	}
	timestamp := id[:12]
	_, err := time.Parse("060102150405", timestamp)
	if err != nil {
		t.Errorf("Invalid timestamp format in transaction ID: %v", err)
	}
}
func TestGenerateUnique6Digits_LengthAndDigits(t *testing.T) {
	for i := 0; i < 100; i++ {
		code := GenerateUnique6Digits()
		if len(code) != 6 {
			t.Errorf("Expected length 6, got %d: %q", len(code), code)
		}
		for j, r := range code {
			if r < '0' || r > '9' {
				t.Errorf("Non-digit character at position %d: %q in %q", j, r, code)
			}
		}
	}
}

func TestGenerateUnique6Digits_Uniqueness(t *testing.T) {
	codes := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		code := GenerateUnique6Digits()
		if code == "" {
			t.Errorf("GenerateUnique6Digits returned empty string")
		}
		if _, exists := codes[code]; exists {
			t.Errorf("Duplicate code generated: %s", code)
		}
		codes[code] = struct{}{}
	}
}

func TestGenerateUnique6Digits_NotEmpty(t *testing.T) {
	code := GenerateUnique6Digits()
	if code == "" {
		t.Error("GenerateUnique6Digits returned empty string")
	}
}
func TestGenerateUUIDv4_FormatAndVersion(t *testing.T) {
	for i := 0; i < 100; i++ {
		uuid := GenerateUUIDv4()
		if uuid == "" {
			t.Errorf("GenerateUUIDv4 returned empty string")
			continue
		}
		// UUID should have 36 characters: 8-4-4-4-12 (including 4 dashes)
		if len(uuid) != 36 {
			t.Errorf("Expected UUID length 36, got %d: %q", len(uuid), uuid)
		}
		// Check dashes at correct positions
		if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
			t.Errorf("UUID dashes in wrong positions: %q", uuid)
		}
		// Check version (should be '4')
		if uuid[14] != '4' {
			t.Errorf("UUID version is not 4: %q", uuid)
		}
		// Check variant (should be 8, 9, a, or b)
		variant := uuid[19]
		if variant != '8' && variant != '9' && variant != 'a' && variant != 'b' {
			t.Errorf("UUID variant is not RFC 4122 compliant: %q", uuid)
		}
	}
}

func TestGenerateUUIDv4_Uniqueness(t *testing.T) {
	uuids := make(map[string]struct{})
	for i := 0; i < 1000; i++ {
		uuid := GenerateUUIDv4()
		if uuid == "" {
			t.Errorf("GenerateUUIDv4 returned empty string")
			continue
		}
		if _, exists := uuids[uuid]; exists {
			t.Errorf("Duplicate UUID generated: %s", uuid)
		}
		uuids[uuid] = struct{}{}
	}
}
