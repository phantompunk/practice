package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		// Add your test cases here
		{"base", 1, true},
		{"example 1", 121, true},
		{"example 2", -121, false},
		{"example 3", 10, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
