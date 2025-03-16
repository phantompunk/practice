package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListStack(t *testing.T) {
	stack := &ListStack[int]{}
	stack.Push(1)
	assert.Equal(t, 1, stack.Size())

	got, err := stack.Pop()
	assertNilError(t, err)
	assertEqual(t, got, 1)

	stack.Push(5)
	got, err = stack.Peek()
	assertNilError(t, err)
	assertEqual(t, got, 5)

	stack.Clear()
	assertEqual(t, stack.Size(), 0)

	got, err = stack.Pop()
	assertError(t, err)
}
