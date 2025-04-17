package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		hour     int
		expected int
	}{
		// Add your test cases here
		{"base", []int{3, 6, 7, 11}, 8, 4},
		{"base", []int{3, 6, 7, 11}, 8, 4},
		{"base", []int{30, 11, 23, 4, 20}, 5, 30},
		{"base", []int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minEatingSpeed(tc.input, tc.hour)
			assert.Equal(t, tc.expected, result)
		})
	}
}
