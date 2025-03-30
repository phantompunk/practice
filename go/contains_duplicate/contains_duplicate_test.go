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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := containsDuplicate(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
