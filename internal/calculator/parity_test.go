package calculator

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"zero is even", 0, true},
		{"positive even", 2, true},
		{"positive odd", 3, false},
		{"negative even", -4, true},
		{"negative odd", -5, false},
		{"large even", 1000000, true},
		{"large odd", 999999, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.input); got != tt.expected {
				t.Errorf("IsEven(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"zero is not odd", 0, false},
		{"positive even", 2, false},
		{"positive odd", 3, true},
		{"negative even", -4, false},
		{"negative odd", -5, true},
		{"large even", 1000000, false},
		{"large odd", 999999, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.input); got != tt.expected {
				t.Errorf("IsOdd(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
