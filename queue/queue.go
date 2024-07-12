package queue

// Queue represents a FIFO (first-in, first-out) queue.
type Queue[T any] struct {
	items []T
}

// New creates and returns a new, empty queue.
func New[T any](initialCapacity int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, initialCapacity),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
// Returns false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

// Peek returns the item at the front of the queue without removing it.
// Returns false if the queue is empty.
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	return q.items[0], true
}

// Len returns the number of items currently in the queue.
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Clear removes all items from the queue, leaving it empty.
func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}
