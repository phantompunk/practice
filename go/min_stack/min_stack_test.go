package kata

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  obj := Constructor()
  obj.Push(-2)
  obj.Push(0)
  obj.Push(-3)
  param_4 := obj.GetMin()
  assert.Equal(t, -3, param_4)
  obj.Pop()
  param_3 := obj.Top()
  assert.Equal(t, 0, param_3)
  param_2 := obj.GetMin()
  assert.Equal(t, -2, param_2)

  // testCases := []struct {
  //   name     string
  //   input    string
  //   expected string
  // }{
  //   // Add your test cases here
  // }
  //
  // for _, tc := range testCases {
  //   t.Run(tc.name, func(t *testing.T) {
  //     result := (tc.input)
  //     assert.Equal(t, tc.expected, result)
  //   })
  // }
}
