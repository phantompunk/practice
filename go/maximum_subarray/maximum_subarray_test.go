package kata

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
  testCases := []struct {
    name     string
    input    []int
    expected int
  }{
    { "Example 1", []int{-2,1,-3,4,-1,2,1,-5,4}, 6 },
    { "Example 2", []int{1}, 1 },
    { "Example 3", []int{5,4,-1,7,8}, 23 },
    { "All Negative", []int{-3,-2,-1,-4}, -1 },
    { "Single Element Negative", []int{-5}, -5 },
    { "Mixed Elements", []int{-2, -1, 2, 3, -4, 5}, 6 },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      result := maxSubArray(tc.input)
      assert.Equal(t, tc.expected, result)
    })
  }
}
