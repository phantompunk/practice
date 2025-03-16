package structures

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	ll := LinkedListInt{}
	assertBoolEqual(t, ll.IsEmpty(), true)
	assertEqual(t, ll.Len(), 0)

	ll.Append(4)
	assertEqual(t, ll.Len(), 1)

	ll.Append(3)
	got, err := ll.GetAt(1)
	assertNilError(t, err)
	assertEqual(t, got, 3)

	ll.Prepend(5)
	got, err = ll.GetAt(0)
	assertEqual(t, ll.Len(), 3)
	assertEqual(t, got, 5)

	ll.InsertAt(1, 6)
	got, err = ll.GetAt(1)
	assertEqual(t, ll.Len(), 4)
	assertEqual(t, got, 6)

	valid := ll.Contains(3)
	assertBoolEqual(t, valid, true)

	values := ll.ToSlice()
	want := []int{5, 6, 4, 3}
	if !reflect.DeepEqual(values, want) {
		t.Errorf("got %v, want %v", values, want)
	}

	ll.RemoveAt(1)
	got, err = ll.GetAt(1)
	assertEqual(t, got, 4)

	ll.RemoveFirst()
	got, err = ll.GetAt(0)
	assertEqual(t, got, 4)

	ll.RemoveLast()
	got, err = ll.GetAt(0)
	assertEqual(t, got, 4)

	ll.Clear()
	assertEqual(t, ll.Len(), 0)
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertBoolEqual(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, actual error) {
	t.Helper()

	if actual == nil {
		t.Errorf("got nil; expected: %v", actual)
	}
}

func assertNilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got %v; expected: nil", actual)
	}
}

func TestGenLinkedList(t *testing.T) {
	var ll LinkedList

	// Helper function to create a new linked list for each test
	newLinkedList := func() LinkedList {
		return &LinkedListInt{} // Replace with your actual implementation
	}

	t.Run("Empty List", func(t *testing.T) {
		ll = newLinkedList()
		assert.Equal(t, 0, ll.Len())
		assert.True(t, ll.IsEmpty())
		assert.Equal(t, []int{}, ll.ToSlice())
	})

	t.Run("Prepend", func(t *testing.T) {
		ll = newLinkedList()
		ll.Prepend(1)
		ll.Prepend(2)
		assert.Equal(t, 2, ll.Len())
		assert.Equal(t, []int{2, 1}, ll.ToSlice())
	})

	t.Run("Append", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		assert.Equal(t, 2, ll.Len())
		assert.Equal(t, []int{1, 2}, ll.ToSlice())
	})

	t.Run("InsertAt", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(3)
		err := ll.InsertAt(1, 2)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, ll.ToSlice())

		err = ll.InsertAt(0, 0)
		assert.NoError(t, err)
		assert.Equal(t, []int{0, 1, 2, 3}, ll.ToSlice())

		err = ll.InsertAt(4, 4)
		assert.NoError(t, err)
		assert.Equal(t, []int{0, 1, 2, 3, 4}, ll.ToSlice())

		err = ll.InsertAt(6, 5)
		assert.Error(t, err)
	})

	t.Run("GetAt", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		val, err := ll.GetAt(1)
		assert.NoError(t, err)
		assert.Equal(t, 2, val)

		_, err = ll.GetAt(3)
		assert.Error(t, err)
	})

	t.Run("RemoveAt", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		err := ll.RemoveAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{2}, ll.ToSlice())
		err = ll.RemoveAt(0)
		assert.NoError(t, err)
		assert.Equal(t, []int{}, ll.ToSlice())
		err = ll.RemoveAt(0)
		assert.Error(t, err)
	})

	t.Run("RemoveFirst", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		err := ll.RemoveFirst()
		assert.NoError(t, err)
		assert.Equal(t, []int{2}, ll.ToSlice())
		err = ll.RemoveFirst()
		assert.NoError(t, err)
		assert.Equal(t, []int{}, ll.ToSlice())
		err = ll.RemoveFirst()
		assert.Error(t, err)
	})

	t.Run("RemoveLast", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		err := ll.RemoveLast()
		assert.NoError(t, err)
		assert.Equal(t, []int{1}, ll.ToSlice())
		err = ll.RemoveLast()
		// assert.NoError(t, err)
		// assert.Equal(t, []int{}, ll.ToSlice())
		// err = ll.RemoveLast()
		// assert.Error(t, err)
	})

	t.Run("Contains", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		assert.True(t, ll.Contains(1))
		assert.False(t, ll.Contains(3))
	})

	t.Run("Clear", func(t *testing.T) {
		ll = newLinkedList()
		ll.Append(1)
		ll.Append(2)
		ll.Clear()
		assert.Equal(t, 0, ll.Len())
		assert.True(t, ll.IsEmpty())
		assert.Equal(t, []int{}, ll.ToSlice())
	})
}
