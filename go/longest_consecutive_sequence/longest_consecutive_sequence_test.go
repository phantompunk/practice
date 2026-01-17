package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		// Add your test cases here
		{"example 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"example 2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"single element", []int{10}, 1},
		{"no consecutive", []int{10, 30, 20}, 1},
		{"negative numbers", []int{-1, -2, -3, -4}, 4},
		{"mixed numbers", []int{1, 2, 0, 1}, 3},
		{"duplicates", []int{1, 2, 2, 3}, 3},
		{"large range", []int{10, 5, 12, 3, 55, 30, 4, 11, 2}, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestConsecutive(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
