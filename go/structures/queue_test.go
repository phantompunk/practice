package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := &ListQueue[int]{}
	queue.Enqueue(5)
	assertEqual(t, queue.Size(), 1)

	got, err := queue.Dequeue()
	assertNilError(t, err)
	assertEqual(t, got, 5)
	assertEqual(t, queue.Size(), 0)

	queue.Enqueue(6)
	assertEqual(t, queue.Size(), 1)
	val, err := queue.Peek()
	assertNilError(t, err)
	assertEqual(t, val, 6)
}

func TestListQueue(t *testing.T) {
	t.Run("Empty Queue", func(t *testing.T) {
		q := &ListQueue[int]{}
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())

		_, err := q.Dequeue()
		assert.Error(t, err)

		_, err = q.Peek()
		assert.Error(t, err)
	})

	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		q := &ListQueue[string]{}

		q.Enqueue("Alice")
		assert.False(t, q.IsEmpty())
		assert.Equal(t, 1, q.Size())

		q.Enqueue("Bob")
		assert.Equal(t, 2, q.Size())

		val, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "Alice", val)
		assert.Equal(t, 1, q.Size())

		val, err = q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, "Bob", val)
		assert.True(t, q.IsEmpty())

		_, err = q.Dequeue()
		assert.Error(t, err)
	})

	t.Run("Peek", func(t *testing.T) {
		q := &ListQueue[int]{}
		q.Enqueue(10)
		q.Enqueue(20)

		val, err := q.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
		assert.Equal(t, 2, q.Size()) // Peek should not change size

		q.Dequeue()
		val, err = q.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 20, val)

		q.Dequeue()
		_, err = q.Peek()
		assert.Error(t, err)

	})

	t.Run("Clear", func(t *testing.T) {
		q := &ListQueue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Clear()
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())
		_, err := q.Dequeue()
		assert.Error(t, err)
		_, err = q.Peek()
		assert.Error(t, err)
	})

	t.Run("One Element Queue", func(t *testing.T) {
		q := &ListQueue[int]{}
		q.Enqueue(1)
		val, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())
		_, err = q.Dequeue()
		assert.Error(t, err)
	})
}
