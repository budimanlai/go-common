package gocommon

import (
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	tt := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "Valid time",
			input:    time.Date(2023, 10, 1, 12, 34, 56, 0, time.UTC),
			expected: "2023-10-01 12:34:56",
		},
		{
			name:     "Zero time",
			input:    time.Time{},
			expected: "0001-01-01 00:00:00",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := TimeToString(&tc.input, "")
			if got != tc.expected {
				t.Errorf("TimeToString(%v) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestStringToTime(t *testing.T) {
	tt := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "Valid string",
			input:    "2023-10-01 12:34:56",
			expected: time.Date(2023, 10, 1, 12, 34, 56, 0, time.UTC),
		},
		{
			name:        "Invalid format",
			input:       "2023/10/01 12:34:56",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expectError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := StringToTime(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("StringToTime(%q) expected error, got nil", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToTime(%q) unexpected error: %v", tc.input, err)
				}
				// Compare time values in UTC to avoid location mismatch
				if !got.UTC().Equal(tc.expected) {
					t.Errorf("StringToTime(%q) = %v; want %v", tc.input, got, tc.expected)
				}
			}
		})
	}
}
func TestTimeToStringAndBack(t *testing.T) {
	tt := []struct {
		name  string
		input time.Time
	}{
		{
			name:  "Round-trip normal time",
			input: time.Date(2022, 5, 17, 8, 30, 45, 0, time.UTC),
		},
		{
			name:  "Round-trip zero time",
			input: time.Time{},
		},
		{
			name:  "Round-trip with non-UTC location",
			input: time.Date(2023, 12, 25, 23, 59, 59, 0, time.FixedZone("WIB", 7*3600)),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			str := TimeToString(&tc.input, "")
			_, err := StringToTime(str)
			if err != nil {
				t.Fatalf("StringToTime(%q) unexpected error: %v", str, err)
			}
		})
	}
}
func TestConvertToLocalTime(t *testing.T) {
	// Save and restore the original location to avoid side effects
	origLoc := time.Local
	defer func() { time.Local = origLoc }()

	// Set a known local timezone for testing
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load location: %v", err)
	}
	time.Local = loc

	tt := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "UTC to Asia/Jakarta",
			input:    time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 10, 1, 19, 0, 0, 0, loc),
		},
		{
			name:     "Already in Asia/Jakarta",
			input:    time.Date(2023, 10, 1, 19, 0, 0, 0, loc),
			expected: time.Date(2023, 10, 1, 19, 0, 0, 0, loc),
		},
		{
			name:     "Zero time",
			input:    time.Time{},
			expected: time.Time{}.In(loc),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertToLocalTime(tc.input)
			if !got.Equal(tc.expected) || got.Location().String() != loc.String() {
				t.Errorf("ConvertToLocalTime(%v) = %v (%v); want %v (%v)", tc.input, got, got.Location(), tc.expected, tc.expected.Location())
			}
		})
	}
}
func TestStringWithTZToLocalTime(t *testing.T) {
	// Save and restore the original location to avoid side effects
	origLoc := time.Local
	defer func() { time.Local = origLoc }()

	// Set a known local timezone for testing
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load location: %v", err)
	}
	time.Local = loc

	tt := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "RFC3339 with +07:00",
			input:    "2023-10-01T12:34:56+07:00",
			expected: time.Date(2023, 10, 1, 12, 34, 56, 0, loc),
		},
		{
			name:     "RFC3339 with Z (UTC)",
			input:    "2023-10-01T05:34:56Z",
			expected: time.Date(2023, 10, 1, 12, 34, 56, 0, loc),
		},
		{
			name:     "Space format with -07:00",
			input:    "2023-10-01 20:34:56-07:00",
			expected: time.Date(2023, 10, 2, 10, 34, 56, 0, loc),
		},
		{
			name:        "Invalid format",
			input:       "2023/10/01 12:34:56+07:00",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expectError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := StringWithTZToLocalTime(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("StringWithTZToLocalTime(%q) expected error, got nil", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringWithTZToLocalTime(%q) unexpected error: %v", tc.input, err)
				}
				if !got.Equal(tc.expected) || got.Location().String() != loc.String() {
					t.Errorf("StringWithTZToLocalTime(%q) = %v (%v); want %v (%v)", tc.input, got, got.Location(), tc.expected, tc.expected.Location())
				}
			}
		})
	}
}
func TestGetCurrentTimeInLocalZone(t *testing.T) {
	// Save and restore the original location to avoid side effects
	origLoc := time.Local
	defer func() { time.Local = origLoc }()

	// Set a known local timezone for testing
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load location: %v", err)
	}
	time.Local = loc

	before := time.Now().In(loc)
	got := GetCurrentTimeInLocalZone()
	after := time.Now().In(loc)

	// Check that the returned time is between before and after, and in the correct location
	if got.Before(before) || got.After(after) {
		t.Errorf("GetCurrentTimeInLocalZone() = %v; want between %v and %v", got, before, after)
	}
	if got.Location().String() != loc.String() {
		t.Errorf("GetCurrentTimeInLocalZone() location = %v; want %v", got.Location(), loc)
	}
}
