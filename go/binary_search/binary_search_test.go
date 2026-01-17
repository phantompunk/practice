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
		{"target present", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"target absent", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"empty array", []int{}, 5, -1},
		{"single element not found", []int{10}, 5, -1},
		{"multiple elements target at start", []int{2, 4, 6, 8, 10}, 2, 0},
		{"multiple elements target at end", []int{2, 4, 6, 8, 10}, 10, 4},
		{"multiple elements target in middle", []int{1, 3, 5, 7, 9}, 5, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
