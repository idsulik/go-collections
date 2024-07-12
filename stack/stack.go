package stack

// Stack represents a LIFO (last-in, first-out) stack.
type Stack[T any] struct {
	items []T
}

// New creates and returns a new, empty stack with the specified initial capacity.
func New[T any](initialCapacity int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, initialCapacity),
	}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// Returns false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

// Peek returns the item at the top of the stack without removing it.
// Returns false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// Len returns the number of items currently in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all items from the stack, leaving it empty.
func (s *Stack[T]) Clear() {
	s.items = make([]T, 0)
}
