package kata

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1}, 1},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 2, 1, 1, 1, 3, 2, 3, 2}, 2},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 3, 2, 3, 2}, 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given %v want %d", tc.nums, tc.want), func(t *testing.T) {
			got := majorityElement(tc.nums)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
