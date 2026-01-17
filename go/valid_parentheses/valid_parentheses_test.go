package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		// Add your test cases here
		{"base", "()", true},
		{"mismatched", "){", false},
		{"all", "()[]{}", true},
		{"invalid", "(}", false},
		{"nested", "([)]", false},
		{"complex valid", "{[]}", true},
		{"single opening", "(", false},
		{"single closing", ")", false},
		{"empty string", "", true},
		{"long valid", "((()))[]{}", true},
		{"long invalid", "((())]", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
