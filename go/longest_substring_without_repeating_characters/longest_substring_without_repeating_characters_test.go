package kata

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		given string
		want  int
	}{
		{"a", 1},
		{"aab", 2},
		{"abca", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvadf", 4},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %d", tc.given, tc.want), func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.given)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
