package kata

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		given []int
		want  bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("given %v expect %v", tc.given, tc.want), func(t *testing.T) {
			got := containsDuplicate(tc.given)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
