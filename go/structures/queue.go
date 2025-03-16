package structures

import "errors"

// Queue is an interface representing a First-In-First-Out (FIFO) data structure.
type Queue[T any] interface {
	// Enqueue adds an element to the rear of the queue.
	Enqueue(item T)
	// Dequeue removes and returns the front element of the queue.
	// Returns an error if the queue is empty.
	Dequeue() (T, error)
	// Peek returns the front element of the queue without removing it.
	// Returns an error if the queue is empty.
	Peek() (T, error)
	// IsEmpty returns true if the queue is empty, false otherwise.
	IsEmpty() bool
	// Size returns the number of elements in the queue.
	Size() int
	// Clear removes all elements from the queue.
	Clear()
}

var ErrEmptyQueue = errors.New("queue is empty")

type ListQueue[T any] struct {
	head *AnyNode[T]
	tail *AnyNode[T]
	size int
}

func (q *ListQueue[T]) Enqueue(item T) {
	q.size++
	node := &AnyNode[T]{value: item}
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *ListQueue[T]) Dequeue() (T, error) {
	if q.head == nil {
		var zero T
		return zero, ErrEmptyQueue
	}

	q.size--
	current := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	return current.value, nil
}

func (q *ListQueue[T]) Peek() (T, error) {
	if q.head == nil {
		var zero T
		return zero, ErrEmptyQueue
	}
	return q.head.value, nil
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *ListQueue[T]) Size() int {
	return q.size
}

func (q *ListQueue[T]) Clear() {
	q.size = 0
	q.head = nil
	q.tail = nil
}
