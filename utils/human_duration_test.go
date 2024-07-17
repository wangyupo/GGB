package utils

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		hasError bool
	}{
		{"1d", 24 * time.Hour, false},
		{"1d2h", 26 * time.Hour, false},
		{"1h30m", 1*time.Hour + 30*time.Minute, false},
		{"45m", 45 * time.Minute, false},
		{"10s", 10 * time.Second, false},
		{"500ms", 500 * time.Millisecond, false},
		{"100us", 100 * time.Microsecond, false},
		{"50ns", 50 * time.Nanosecond, false},
		{"1d7h10m", 24*time.Hour + 7*time.Hour + 10*time.Minute, false},
		{"", 0, true},
		{"1d1x", 0, true},
		{"-1d", 0, true},
		{"1d-1h", 0, true},
	}

	for _, test := range tests {
		result, err := ParseDuration(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseDuration(%q) returned error %v, expected error: %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ParseDuration(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
