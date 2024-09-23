package arraystack

import (
	"testing"
)

// TestNew tests the creation of a new stack with an initial capacity.
func TestNew(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)

	if got := s.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestArrayStackPush tests adding items to the stack.
func TestArrayStackPush(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
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

// TestArrayStackPopEmpty tests popping from an empty stack.
func TestArrayStackPopEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop() should return false for an empty stack")
	}
}

// TestArrayStackPeek tests peeking at the top of the stack.
func TestArrayStackPeek(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
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

// TestArrayStackPeekEmpty tests peeking into an empty stack.
func TestArrayStackPeekEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
	if _, ok := s.Peek(); ok {
		t.Errorf("Peek() should return false for an empty stack")
	}
}

// TestArrayStackLen tests the length of the stack.
func TestArrayStackLen(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
	if got := s.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}

	s.Push(1)
	s.Push(2)
	if got := s.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}
}

// TestArrayStackIsEmpty tests checking if the stack is empty.
func TestArrayStackIsEmpty(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
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

// TestArrayStackClear tests clearing the stack.
func TestArrayStackClear(t *testing.T) {
	// Create a new stack with an initial capacity of 10
	s := New[int](10)
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
