package ringbuffer

// RingBuffer represents a circular buffer of fixed size
type RingBuffer[T any] struct {
	buffer []T
	size   int
	head   int // points to the next write position
	tail   int // points to the next read position
	count  int // number of elements currently in buffer
}

// New creates a new RingBuffer with the specified capacity
func New[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &RingBuffer[T]{
		buffer: make([]T, capacity),
		size:   capacity,
	}
}

// Write adds an item to the buffer, overwriting the oldest item if the buffer is full
func (r *RingBuffer[T]) Write(item T) bool {
	if r.count == r.size {
		return false // Buffer is full
	}

	r.buffer[r.head] = item
	r.head = (r.head + 1) % r.size
	r.count++
	return true
}

// Read removes and returns the oldest item from the buffer
func (r *RingBuffer[T]) Read() (T, bool) {
	var zero T
	if r.count == 0 {
		return zero, false // Buffer is empty
	}

	item := r.buffer[r.tail]
	r.tail = (r.tail + 1) % r.size
	r.count--
	return item, true
}

// Peek returns the oldest item without removing it
func (r *RingBuffer[T]) Peek() (T, bool) {
	var zero T
	if r.count == 0 {
		return zero, false
	}
	return r.buffer[r.tail], true
}

// IsFull returns true if the buffer is at capacity
func (r *RingBuffer[T]) IsFull() bool {
	return r.count == r.size
}

// IsEmpty returns true if the buffer contains no items
func (r *RingBuffer[T]) IsEmpty() bool {
	return r.count == 0
}

// Cap returns the total capacity of the buffer
func (r *RingBuffer[T]) Cap() int {
	return r.size
}

// Len returns the current number of items in the buffer
func (r *RingBuffer[T]) Len() int {
	return r.count
}

// Clear removes all items from the buffer
func (r *RingBuffer[T]) Clear() {
	r.head = 0
	r.tail = 0
	r.count = 0
}
