package queue

import (
	"github.com/idsulik/go-collections/v2/iterator"
)

// Iterator implements iterator.Iterator for Queue
type Iterator[T any] struct {
	current int
	items   []T
}

// NewIterator creates a new iterator for the queue
func NewIterator[T any](q *Queue[T]) iterator.Iterator[T] {
	it := &Iterator[T]{
		current: 0,
		items:   make([]T, q.Len()),
	}

	// Take a snapshot of current queue items
	// This ensures modifications to the queue won't affect iteration
	if q.Len() > 0 {
		copy(it.items, q.GetItems())
	}

	return it
}

// HasNext returns true if there are more elements to iterate over
func (it *Iterator[T]) HasNext() bool {
	return it.current < len(it.items)
}

// Next returns the next element in the iteration
// Returns the zero value and false if there are no more elements
func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}

	value := it.items[it.current]
	it.current++
	return value, true
}

// Reset restarts iteration from the beginning
// Uses the same snapshot of items from when iterator was created
func (it *Iterator[T]) Reset() {
	it.current = 0
}
