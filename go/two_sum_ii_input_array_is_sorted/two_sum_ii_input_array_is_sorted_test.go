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
		{"Example Test Case 1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Example Test Case 2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"Example Test Case 3", []int{-1, 0}, -1, []int{1, 2}},
		{"Additional Test Case 1", []int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
		{"Additional Test Case 2", []int{5, 25, 75}, 100, []int{2, 3}},
		{"Additional Test Case 3", []int{1, 3, 5, 7, 9, 11, 12, 14}, 16, []int{3, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
