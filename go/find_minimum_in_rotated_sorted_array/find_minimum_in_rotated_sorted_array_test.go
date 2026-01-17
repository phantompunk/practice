package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		// Add your test cases here
		{"init", []int{1}, 1},
		{"init1", []int{3, 1}, 1},
		{"init2", []int{3, 4, 5, 1, 2}, 1},
		{"init3", []int{4, 5, 6, 7, 0, 1, 2}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMin(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
