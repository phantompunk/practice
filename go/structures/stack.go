package structures

import "errors"

// Stack is an interface representing a Last-In-First-Out (LIFO) data structure.
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(item T)
	// Pop removes and returns the top element of the stack.
	// Returns an error if the stack is empty.
	Pop() (T, error)
	// Peek returns the top element of the stack without removing it.
	// Returns an error if the stack is empty.
	Peek() (T, error)
	// IsEmpty returns true if the stack is empty, false otherwise.
	IsEmpty() bool
	// Size returns the number of elements in the stack.
	Size() int
	// Clear removes all elements from the stack.
	Clear()
}

var ErrEmptyStack = errors.New("stack is empty")

type AnyNode[T any] struct {
	value T
	next  *AnyNode[T]
}

type ListStack[T any] struct {
	head *AnyNode[T]
	size int
}

func (s *ListStack[T]) Push(item T) {
	s.size++
	node := &AnyNode[T]{value: item}
	if s.head == nil {
		s.head = node
	}
}

func (s *ListStack[T]) Pop() (T, error) {
	if s.head == nil {
		var zero T
		return zero, ErrEmptyStack
	}

	top := s.head
	s.head = top.next
	s.size--
	return top.value, nil
}

func (s *ListStack[T]) Peek() (T, error) {
	if s.head == nil {
		var zero T
		return zero, ErrEmptyStack
	}

	return s.head.value, nil
}

func (s *ListStack[T]) Size() int {
	return s.size
}

func (s *ListStack[T]) Clear() {
	s.size = 0
	s.head = nil
}
