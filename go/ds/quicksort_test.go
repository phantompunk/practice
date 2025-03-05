package ds

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	given := []int{9, 5, 8, 2}
	want := []int{2, 5, 8, 9}

	got := quicksort(given)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
