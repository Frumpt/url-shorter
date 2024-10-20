package random

import (
	"testing"
)

func TestNewRandomString(t *testing.T) {
	var tests = []struct {
		name        string
		stingLength int
		expected    int
	}{
		{
			name:        "string length to 2",
			stingLength: 2,
			expected:    2,
		},
		{
			name:        "string length to 6",
			stingLength: 6,
			expected:    6,
		},
		{
			name:        "string length to 12",
			stingLength: 12,
			expected:    12,
		},
		{
			name:        "string length to 25",
			stingLength: 25,
			expected:    25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewRandomString(tt.stingLength)
			if len(actual) != tt.expected {
				t.Errorf("actual: %v, but expected: %v", len(actual), tt.expected)
			}
		})
	}
}
