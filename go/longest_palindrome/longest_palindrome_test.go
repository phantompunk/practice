package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		// Add your test cases here
		{"longest is 4", "cccc", 4},
		{"longest is 7", "abccccdd", 7},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
