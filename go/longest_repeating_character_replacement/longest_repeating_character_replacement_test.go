package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		replace  int
		expected int
	}{
		// Add your test cases here
		{"empty", "", 1, 0},
		{"base", "A", 1, 1},
		{"simple", "AB", 1, 2},
		{"test1", "ABAB", 2, 4},
		{"test2", "AABABBA", 1, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := characterReplacement(tc.input, tc.replace)
			assert.Equal(t, tc.expected, result)
		})
	}
}
