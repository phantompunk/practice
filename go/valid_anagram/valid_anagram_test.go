package kata

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		strA string
		strB string
		want bool
	}{
		{"a", "b", false},
		{"a", "a", true},
		{"a", "aa", false},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"aacc", "ccac", false},
		{"rat", "car", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s -> %s", tc.strA, tc.strB), func(t *testing.T) {
			got := isAnagram(tc.strA, tc.strB)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
