package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{name: "Example 1", input: []int{1, 1, 1, 2, 2, 3}, target: 2, expected: []int{1, 2}},
		{name: "Example 2", input: []int{1}, target: 1, expected: []int{1}},
		{name: "Example 3", input: []int{4, 4, 4, 6, 6, 2, 2, 2, 2}, target: 2, expected: []int{2, 4}},
		{name: "Example 4", input: []int{5, 5, 5, 5, 5, 1, 1, 1, 3, 3, 2}, target: 3, expected: []int{5, 1, 3}},
		{name: "Example 5", input: []int{10, 9, 9, 8, 8, 8, 7}, target: 2, expected: []int{8, 9}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := topKFrequent(tc.input, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}
