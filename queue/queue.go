package queue

import "github.com/idsulik/go-collections/v2/deque"

type Queue[T any] struct {
	d *deque.Deque[T]
}

// New creates and returns a new, empty queue.
func New[T any](initialCapacity int) *Queue[T] {
	return &Queue[T]{
		d: deque.New[T](initialCapacity),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.d.PushBack(item)
}

// Dequeue removes and returns the item at the front of the queue.
// Returns false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	return q.d.PopFront()
}

// Peek returns the item at the front of the queue without removing it.
// Returns false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	return q.d.PeekFront()
}

// Len returns the number of items currently in the queue.
func (q *Queue[T]) Len() int {
	return q.d.Len()
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.d.IsEmpty()
}

// Clear removes all items from the queue, leaving it empty.
func (q *Queue[T]) Clear() {
	q.d.Clear()
}
