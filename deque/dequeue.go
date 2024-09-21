package deque

const (
	defaultCapacity = 16 // Default initial capacity for the deque
	resizeFactor    = 2  // Factor by which the deque is resized when full
)

type Deque[T any] struct {
	buffer         []T // Underlying slice to hold elements
	head, tail     int // Indices for the front and back of the deque
	size, capacity int // Current size and maximum capacity of the deque
}

// New creates a new Deque with the specified initial capacity.
// If the initial capacity is less than 1, it uses the default capacity.
func New[T any](initialCapacity int) *Deque[T] {
	if initialCapacity < 1 {
		initialCapacity = defaultCapacity
	}

	return &Deque[T]{
		buffer:   make([]T, initialCapacity),
		capacity: initialCapacity,
	}
}

// resize doubles the capacity of the deque and repositions elements.
func (d *Deque[T]) resize() {
	newCapacity := d.capacity * resizeFactor
	newBuffer := make([]T, newCapacity)

	// Copy elements into the new buffer, handling wrap-around
	if d.tail >= d.head {
		copy(newBuffer, d.buffer[d.head:d.tail])
	} else {
		n := copy(newBuffer, d.buffer[d.head:])
		copy(newBuffer[n:], d.buffer[:d.tail])
	}

	d.head = 0
	d.tail = d.size
	d.buffer = newBuffer
	d.capacity = newCapacity
}

// PushFront inserts an item at the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	if d.size == d.capacity {
		d.resize()
	}
	d.head = (d.head - 1 + d.capacity) % d.capacity
	d.buffer[d.head] = item
	d.size++
}

// PushBack inserts an item at the back of the deque.
func (d *Deque[T]) PushBack(item T) {
	if d.size == d.capacity {
		d.resize()
	}
	d.buffer[d.tail] = item
	d.tail = (d.tail + 1) % d.capacity
	d.size++
}

// PopFront removes and returns the item at the front of the deque.
// Returns false if the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	item := d.buffer[d.head]
	d.head = (d.head + 1) % d.capacity
	d.size--
	return item, true
}

// PopBack removes and returns the item at the back of the deque.
// Returns false if the deque is empty.
func (d *Deque[T]) PopBack() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	d.tail = (d.tail - 1 + d.capacity) % d.capacity
	item := d.buffer[d.tail]
	d.size--
	return item, true
}

// PeekFront returns the item at the front of the deque without removing it.
// Returns false if the deque is empty.
func (d *Deque[T]) PeekFront() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	return d.buffer[d.head], true
}

// PeekBack returns the item at the back of the deque without removing it.
// Returns false if the deque is empty.
func (d *Deque[T]) PeekBack() (T, bool) {
	if d.size == 0 {
		var zero T
		return zero, false
	}
	index := (d.tail - 1 + d.capacity) % d.capacity
	return d.buffer[index], true
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int {
	return d.size
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Clear removes all elements from the deque.
func (d *Deque[T]) Clear() {
	d.buffer = make([]T, d.capacity)
	d.head = 0
	d.tail = 0
	d.size = 0
}
