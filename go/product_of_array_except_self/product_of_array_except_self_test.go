package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Example 2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"Single Element", []int{5}, []int{1}},
		{"Two Elements", []int{2, 3}, []int{3, 2}},
		{"With Negative", []int{-2, 4, -5}, []int{-20, 10, -8}},
		{"With Zeroes", []int{0, 0, 2}, []int{0, 0, 0}},
		{"All Ones", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := productExceptSelf(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
