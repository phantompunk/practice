package kata

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		given []int
		want  int
	}{
		{[]int{1, 0, 1, 2}, 3},
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %d", tc.given, tc.want), func(t *testing.T) {
			got := longestConsecutive(tc.given)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
