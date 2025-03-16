package structures

import (
	"errors"
)

// LinkedList is an interface representing a singly linked list of integers.
type LinkedList interface {
	// Len returns the number of elements in the list.
	Len() int
	// IsEmpty returns true if the list is empty, false otherwise.
	IsEmpty() bool
	// Prepend adds a new element to the beginning of the list.
	Prepend(value int)
	// Append adds a new element to the end of the list.
	Append(value int)
	// InsertAt inserts a new element at the specified index.
	// Returns an error if the index is out of bounds.
	InsertAt(index int, value int) error
	// GetAt returns the element at the specified index.
	// Returns an error if the index is out of bounds.
	GetAt(index int) (int, error)
	// RemoveAt removes the element at the specified index.
	// Returns an error if the index is out of bounds.
	RemoveAt(index int) error
	// RemoveFirst removes the first element of the list.
	// Returns an error if the list is empty.
	RemoveFirst() error
	// RemoveLast removes the last element of the list.
	// Returns an error if the list is empty.
	RemoveLast() error
	// Contains returns true if the list contains the specified value, false otherwise.
	Contains(value int) bool
	// ToSlice returns a slice containing all elements of the list.
	ToSlice() []int
	// Clear removes all elements from the list.
	Clear()
}

var ErrIndexOutOfBound = errors.New("index out of bounds")
var ErrIsEmpty = errors.New("list is empty")

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type LinkedListInt struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedListInt) Len() int {
	return ll.size
}

func (ll *LinkedListInt) IsEmpty() bool {
	return ll.Len() == 0
}

func (ll *LinkedListInt) Prepend(value int) {
	ll.size++
	node := &Node{value: value}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
		return
	}

	node.next = ll.head
	ll.head.prev = node
	ll.head = node
}

func (ll *LinkedListInt) Append(value int) {
	ll.size++
	node := &Node{value: value}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	node.prev = ll.tail
	ll.tail = current.next
}

func (ll *LinkedListInt) InsertAt(index, value int) error {
	if index < 0 || index > ll.Len() {
		return ErrIndexOutOfBound
	} else if index == 0 {
		ll.Prepend(value)
		return nil
	} else if index == ll.Len() {
		ll.Append(value)
		return nil
	}

	ll.size++
	node := &Node{value: value}
	current := ll.head
	for range index - 1 {
		current = current.next
	}

	node.next = current.next
	node.prev = current
	if current.next != nil {
		current.next.prev = node
	}
	current.next = node
	return nil
}

func (ll *LinkedListInt) GetAt(index int) (int, error) {
	if index < 0 || index > ll.size {
		return 0, ErrIndexOutOfBound
	}

	i := 0
	current := ll.head
	for i < ll.Len() {
		if i == index {
			return current.value, nil
		}
		current = current.next
		i++
	}
	return 0, nil
}

func (ll *LinkedListInt) RemoveAt(index int) error {
	if index < 0 || index > ll.Len() {
		return ErrIndexOutOfBound
	} else if index == 0 {
		return ll.RemoveFirst()
	} else if index == ll.Len() {
		return ll.RemoveLast()
	}

	ll.size--
	current := ll.head
	for range index - 1 {
		current = current.next
	}

	if current.next.next != nil {
		current.next = current.next.next
	}

	return nil
}

func (ll *LinkedListInt) RemoveFirst() error {
	if ll.head == nil {
		return ErrIsEmpty
	}

	ll.size--
	ll.head = ll.head.next
	return nil
}

func (ll *LinkedListInt) RemoveLast() error {
	ll.size--
	current := ll.tail
	if current == nil {
		return ErrIsEmpty
	}
	if current.prev != nil {
		current.prev.next = nil
	}
	ll.tail = ll.tail.prev
	return nil
}

func (ll *LinkedListInt) Contains(value int) bool {
	current := ll.head
	for range ll.size {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedListInt) ToSlice() []int {
	results := []int{}

	current := ll.head
	for range ll.size {
		results = append(results, current.value)
		current = current.next
	}

	return results
}

func (ll *LinkedListInt) Clear() {
	ll.size = 0
	ll.head = nil
	ll.tail = nil
}
