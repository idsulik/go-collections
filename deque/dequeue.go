package deque

// Deque represents a double-ended queue (deque) where elements can be added
// or removed from both the front and the back.
type Deque[T any] struct {
	items []T
}

// New creates and returns a new, empty deque with an initial capacity specified by
// the `initialCapacity` parameter.
func New[T any](initialCapacity int) *Deque[T] {
	return &Deque[T]{
		items: make([]T, 0, initialCapacity),
	}
}

// PushFront adds an item to the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// PushBack adds an item to the back of the deque.
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// PopFront removes and returns the item at the front of the deque.
// Returns false if the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

// PopBack removes and returns the item at the back of the deque.
// Returns false if the deque is empty.
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

// PeekFront returns the item at the front of the deque without removing it.
// Returns false if the deque is empty.
func (d *Deque[T]) PeekFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	return d.items[0], true
}

// PeekBack returns the item at the back of the deque without removing it.
// Returns false if the deque is empty.
func (d *Deque[T]) PeekBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}

	return d.items[len(d.items)-1], true
}

// Len returns the number of items in the deque.
func (d *Deque[T]) Len() int {
	return len(d.items)
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

// Clear removes all items from the deque, leaving it empty.
func (d *Deque[T]) Clear() {
	d.items = make([]T, 0)
}
