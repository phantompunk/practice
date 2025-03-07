package kata

import (
	"reflect"
	"testing"
)

func TestTopK(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 3}},
		{[]int{25, 21, 7, 8, 7, 7, 21, 21, 21, 21}, 3, []int{21, 7}},
	}

	for _, tc := range testCases {
		got := topKFrequent(tc.nums, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}
