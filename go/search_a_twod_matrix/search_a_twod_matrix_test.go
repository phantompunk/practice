package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		target   int
		expected bool
	}{
		// Add your test cases here
		{"1d", [][]int{{1}}, 1, true},
		{"1x3", [][]int{{1, 3, 5}}, 1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchMatrix(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
