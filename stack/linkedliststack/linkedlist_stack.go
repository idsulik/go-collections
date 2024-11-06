package linkedliststack

import "github.com/idsulik/go-collections/v2/linkedlist"

// LinkedListStack represents a LIFO (last-in, first-out) stack implemented using a linked list.
type LinkedListStack[T any] struct {
	linkedList *linkedlist.LinkedList[T]
}

// New creates and returns a new, empty linked list stack.
func New[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		linkedList: linkedlist.New[T](),
	}
}

// Push adds an item to the top of the stack.
func (s *LinkedListStack[T]) Push(item T) {
	s.linkedList.AddFront(item)
}

// Pop removes and returns the item from the top of the stack.
// Returns false if the stack is empty.
func (s *LinkedListStack[T]) Pop() (T, bool) {
	return s.linkedList.RemoveFront()
}

// Peek returns the item at the top of the stack without removing it.
// Returns false if the stack is empty.
func (s *LinkedListStack[T]) Peek() (T, bool) {
	return s.linkedList.PeekFront()
}

// Len returns the number of items currently in the stack.
func (s *LinkedListStack[T]) Len() int {
	return s.linkedList.Size()
}

// IsEmpty checks if the stack is empty.
func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.linkedList.IsEmpty()
}

// Clear removes all items from the stack, leaving it empty.
func (s *LinkedListStack[T]) Clear() {
	s.linkedList.Clear()
}
