package kata

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range testCases {
		got := productExceptSelf(tc.nums)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}
