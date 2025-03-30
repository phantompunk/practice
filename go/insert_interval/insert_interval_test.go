package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		given  [][]int
		target []int
		want   [][]int
	}{
		{[][]int{}, []int{2, 5}, [][]int{{2, 5}}},
		{[][]int{{1, 5}}, []int{5, 7}, [][]int{{1, 7}}},
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {4, 5}, {6, 9}}, []int{1, 10}, [][]int{{1, 10}}},
		{[][]int{{1, 2}, {6, 9}}, []int{3, 5}, [][]int{{1, 2}, {3, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %v", tc.given, tc.target), func(t *testing.T) {
			got := insert(tc.given, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
