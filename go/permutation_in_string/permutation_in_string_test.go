package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	testCases := []struct {
		name     string
		perm     string
		input    string
		expected bool
	}{
		// Add your test cases here
		{"invalid", "ab", "a", false},
		{"base", "a", "ab", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkInclusion(tc.perm, tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
