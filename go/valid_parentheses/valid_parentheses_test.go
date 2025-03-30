package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		target   string
		expected bool
	}{
		{name: "valid anagram", input: "anagram", target: "nagram", expected: true},
		{name: "invalid anagram", input: "rat", target: "car", expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
