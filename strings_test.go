package gocommon

import (
	"strings"
	"testing"
)

func TestCapitalizeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"jOHN doe", "John Doe"},
		{"ALICE", "Alice"},
		{"bob smith", "Bob Smith"},
		{"", ""},
		{"  john   doe  ", "John Doe"},
		{"éLÉOnore d'ARTAGNAN", "Éléonore D'artagnan"},
		{"mArY aNnE o'connor", "Mary Anne O'connor"},
		{"123 abc", "123 Abc"},
		{"a", "A"},
		{"A", "A"},
		{"  ", ""},
	}

	for _, tt := range tests {
		result := CapitalizeName(tt.input)
		if result != tt.expected {
			t.Errorf("CapitalizeName(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		length   int
		expected int
	}{
		{0, 0},
		{-5, 0},
		{1, 1},
		{5, 5},
		{12, 12},
		{100, 100},
	}

	for _, tt := range tests {
		result := GenerateRandomString(tt.length)
		if len(result) != tt.expected {
			t.Errorf("GenerateRandomString(%d) length = %d; want %d", tt.length, len(result), tt.expected)
		}
	}

	// Test that the generated string only contains allowed characters
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-"
	for i := 0; i < 10; i++ {
		str := GenerateRandomString(32)
		for _, c := range str {
			if !strings.ContainsRune(charset, c) {
				t.Errorf("GenerateRandomString(32) contains invalid character: %q", c)
			}
		}
	}

	// Test randomness: generate multiple strings and ensure they are not all the same
	set := make(map[string]struct{})
	for i := 0; i < 10; i++ {
		s := GenerateRandomString(16)
		set[s] = struct{}{}
	}
	if len(set) < 2 {
		t.Error("GenerateRandomString(16) does not appear to generate random strings")
	}
}
