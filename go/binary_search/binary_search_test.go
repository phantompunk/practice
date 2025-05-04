package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		// Add your test cases here
		{"base case", []int{1}, 1, 0},
		{"base case", []int{1, 2}, 2, 1},
		{"base case", []int{1, 2, 5}, 5, 2},
		{"base case", []int{-4}, 1, -1},
		{"initial case", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"initial case", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
