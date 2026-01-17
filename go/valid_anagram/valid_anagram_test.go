package kata

import (
	"fmt"
	"strings"
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
		{"listen", "silent", true},
		{"triangle", "integral", true},
		{"apple", "pale", false},
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

func BenchmarkIsAnagram(b *testing.B) {
	s := strings.Repeat("abcdefghij", 100) // n = 1000
	t := strings.Repeat("jihgfedcba", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

func BenchmarkIsAnagramSort(b *testing.B) {
	s := strings.Repeat("abcdefghij", 100) // n = 1000
	t := strings.Repeat("jihgfedcba", 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagramSort(s, t)
	}
}
