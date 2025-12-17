package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "example 1 - duplicates that sum to target",
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "example 2 - target in middle",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3 - target at beginning",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "target at end",
			input:    []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "large numbers",
			input:    []int{1000000, 2000000, 3000000},
			target:   5000000,
			expected: []int{1, 2},
		},
		{
			name:     "zero in array",
			input:    []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "target with zero",
			input:    []int{-1, 0, 1, 2},
			target:   1,
			expected: []int{1, 2},
		},
		{
			name:     "minimum array size",
			input:    []int{1, 2},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "longer array",
			input:    []int{1, 5, 3, 7, 9, 2, 4, 6, 8},
			target:   10,
			expected: []int{2, 3},
		},
		{
			name:     "same number appears multiple times",
			input:    []int{5, 5, 5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "negative target",
			input:    []int{-5, -3, -1, 0, 2},
			target:   -4,
			expected: []int{1, 2},
		},
		{
			name:     "all negative numbers",
			input:    []int{-10, -20, -30, -40},
			target:   -50,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
