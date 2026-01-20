package kata

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
  testCases := []struct {
    name     string
    input    []int
    expected int
  }{
    {"Example Test Case 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
    {"Example Test Case 2", []int{1, 1}, 1},
    {"Additional Test Case 1", []int{4, 3, 2, 1, 4}, 16},
    {"Additional Test Case 2", []int{1, 2, 1}, 2},
    {"Additional Test Case 3", []int{1, 3, 2, 5, 25, 24, 5}, 24},
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      result := maxArea(tc.input)
      assert.Equal(t, tc.expected, result)
    })
  }
}
