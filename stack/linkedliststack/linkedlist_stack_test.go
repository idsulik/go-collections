package linkedliststack

import (
	"testing"
)

// TestNew tests the creation of a new stack with an initial capacity.
func TestNew(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()

	if got := s.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestLinkedListStackPush tests adding items to the stack.
func TestLinkedListStackPush(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Len(); got != 3 {
		t.Errorf("Len() = %d; want 3", got)
	}
	if got, ok := s.Pop(); !ok || got != 3 {
		t.Errorf("Pop() = %d, %v; want 3, true", got, ok)
	}
}

// TestLinkedListStackPopEmpty tests popping from an empty stack.
func TestLinkedListStackPopEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop() should return false for an empty stack")
	}
}

// TestLinkedListStackPeek tests peeking at the top of the stack.
func TestLinkedListStackPeek(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	s.Push(1)
	s.Push(2)

	if got, ok := s.Peek(); !ok || got != 2 {
		t.Errorf("Peek() = %d, %v; want 2, true", got, ok)
	}

	// Ensure Peek does not remove the item
	if got, ok := s.Peek(); !ok || got != 2 {
		t.Errorf("Peek() = %d, %v; want 2, true after re-peeking", got, ok)
	}
}

// TestLinkedListStackPeekEmpty tests peeking into an empty stack.
func TestLinkedListStackPeekEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	if _, ok := s.Peek(); ok {
		t.Errorf("Peek() should return false for an empty stack")
	}
}

// TestLinkedListStackLen tests the length of the stack.
func TestLinkedListStackLen(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	if got := s.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}

	s.Push(1)
	s.Push(2)
	if got := s.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}
}

// TestLinkedListStackIsEmpty tests checking if the stack is empty.
func TestLinkedListStackIsEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true after Pop")
	}
}

// TestLinkedListStackClear tests clearing the stack.
func TestLinkedListStackClear(t *testing.T) {
	// Create a new stack with an initial capacity of
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Clear()

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true after Clear")
	}
	if got := s.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0 after Clear", got)
	}
}
