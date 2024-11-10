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

// resize increases the capacity of the deque.
func (d *Deque[T]) resize() {
	newCapacity := d.capacity * resizeFactor
	d.reallocate(newCapacity)
}

// reallocate creates a new buffer with the specified capacity and copies elements.
func (d *Deque[T]) reallocate(newCapacity int) {
	newBuffer := make([]T, newCapacity)
	if d.size > 0 {
		if d.tail > d.head {
			copy(newBuffer, d.buffer[d.head:d.tail])
		} else {
			n := copy(newBuffer, d.buffer[d.head:])
			copy(newBuffer[n:], d.buffer[:d.tail])
		}
	}

	d.buffer = newBuffer
	d.capacity = newCapacity
	d.head = 0
	d.tail = d.size
}

// PushFront inserts an item at the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	if d.size == d.capacity {
		d.resize()
	}
	if d.head == 0 {
		d.head = d.capacity - 1
	} else {
		d.head--
	}
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
	var zero T
	d.buffer[d.head] = zero // Clear reference
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
	if d.tail == 0 {
		d.tail = d.capacity - 1
	} else {
		d.tail--
	}
	item := d.buffer[d.tail]
	var zero T
	d.buffer[d.tail] = zero // Clear reference
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
	index := d.tail
	if index == 0 {
		index = d.capacity - 1
	} else {
		index--
	}
	return d.buffer[index], true
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int {
	return d.size
}

// Cap returns the current capacity of the deque.
func (d *Deque[T]) Cap() int {
	return d.capacity
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Clear removes all elements from the deque.
func (d *Deque[T]) Clear() {
	// Clear references to help GC
	for i := range d.buffer {
		var zero T
		d.buffer[i] = zero
	}
	d.head = 0
	d.tail = 0
	d.size = 0
	// Reset to default capacity if current capacity is much larger
	if d.capacity > defaultCapacity*2 {
		d.buffer = make([]T, defaultCapacity)
		d.capacity = defaultCapacity
	}
}

// GetItems returns a new slice containing the deque's elements in order.
func (d *Deque[T]) GetItems() []T {
	items := make([]T, d.size)
	if d.size == 0 {
		return items
	}

	if d.tail > d.head {
		copy(items, d.buffer[d.head:d.tail])
	} else {
		n := copy(items, d.buffer[d.head:])
		copy(items[n:], d.buffer[:d.tail])
	}
	return items
}

// Clone returns a deep copy of the deque.
func (d *Deque[T]) Clone() *Deque[T] {
	newDeque := &Deque[T]{
		buffer:   make([]T, d.capacity),
		head:     d.head,
		tail:     d.tail,
		size:     d.size,
		capacity: d.capacity,
	}
	copy(newDeque.buffer, d.buffer)
	return newDeque
}
