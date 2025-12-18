package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		// Add your test cases here
		{"one is dupe", []int{1, 2, 1}, true},
		{"no dupes", []int{1, 2, 3, 4}, false},
		{"all dupes", []int{5, 5, 5, 5}, true},
		{"empty array", []int{}, false},
		{"single element", []int{42}, false},
		{"large array with dupes", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, true},
		{"large array without dupes", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := containsDuplicate(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
