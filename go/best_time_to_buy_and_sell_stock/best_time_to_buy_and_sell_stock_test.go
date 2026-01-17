package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		// Add your test cases here
		{"single day", []int{5}, 0},
		{"two days increasing", []int{1, 5}, 4},
		{"two days decreasing", []int{5, 1}, 0},
		{"multiple peaks", []int{3, 2, 6, 1, 4}, 4},
		{"constant prices", []int{2, 2, 2, 2}, 0},
		{"example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example 2", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProfit(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
