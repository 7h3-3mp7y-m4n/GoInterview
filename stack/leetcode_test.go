package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"basic valid case", "({[]})", true},
		{"unbalanced brackets", "({[})", false},
		{"only opening brackets", "(((", false},
		{"only closing brackets", ")))", false},
		{"empty string", "", true},
		{"nested valid", "{[()()]}", true},
		{"mismatched pair", "{[(])}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.expected {
				t.Errorf("isValid(%q) = %v; expected %v", tt.s, got, tt.expected)
			}
		})
	}
}
