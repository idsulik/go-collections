package arraystack

// ArrayStack represents a LIFO (last-in, first-out) stack implemented using a slice.
type ArrayStack[T any] struct {
	items []T
}

// New creates and returns a new, empty stack with the specified initial capacity.
func New[T any](initialCapacity int) *ArrayStack[T] {
	return &ArrayStack[T]{
		items: make([]T, 0, initialCapacity),
	}
}

// Push adds an item to the top of the stack.
func (s *ArrayStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// Returns false if the stack is empty.
func (s *ArrayStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items[index] = *new(T) // remove reference
	s.items = s.items[:index]

	return item, true
}

// Peek returns the item at the top of the stack without removing it.
// Returns false if the stack is empty.
func (s *ArrayStack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// Len returns the number of items currently in the stack.
func (s *ArrayStack[T]) Len() int {
	return len(s.items)
}

// IsEmpty checks if the stack is empty.
func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all items from the stack, leaving it empty.
func (s *ArrayStack[T]) Clear() {
	for i := range s.items {
		var zero T
		s.items[i] = zero
	}
	s.items = s.items[:0]
}
