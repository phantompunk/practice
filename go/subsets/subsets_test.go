package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{"base", []int{0}, [][]int{{}, {0}}},
		// Add your test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := subsets(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
