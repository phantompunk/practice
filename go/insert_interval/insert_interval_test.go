package kata

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
  testCases := []struct {
    name     string
    input    [][]int
    newInput []int
    expected [][]int
  }{
    {"insert into empty intervals", [][]int{}, []int{5, 7}, [][]int{{5, 7}}},
    {"no overlap, insert at beginning", [][]int{{8, 10}, {12, 16}}, []int{5, 7}, [][]int{{5, 7}, {8, 10}, {12, 16}}},
    {"no overlap, insert at end", [][]int{{1, 3}, {6, 9}}, []int{10, 11}, [][]int{{1, 3}, {6, 9}, {10, 11}}},
    {"overlap with one interval", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
    {"overlap with multiple intervals", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
    {"new interval inside existing interval", [][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
    {"new interval covers all existing intervals", [][]int{{2, 3}, {5, 7}}, []int{1, 8}, [][]int{{1, 8}}},
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      result := insert(tc.input, tc.newInput)
      assert.Equal(t, tc.expected, result)
    })
  }
}
