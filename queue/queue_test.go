package queue

import (
	"testing"
)

// TestNewQueue tests the creation of a new queue with initial capacity.
func TestNewQueue(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)

	if got := q.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestEnqueue tests adding items to the queue.
func TestEnqueue(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got := q.Len(); got != 3 {
		t.Errorf("Len() = %d; want 3", got)
	}
	if got, ok := q.Dequeue(); !ok || got != 1 {
		t.Errorf("Dequeue() = %d, %v; want 1, true", got, ok)
	}
}

// TestDequeueEmpty tests dequeueing from an empty queue.
func TestDequeueEmpty(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	if _, ok := q.Dequeue(); ok {
		t.Errorf("Dequeue() should return false for an empty queue")
	}
}

// TestPeek tests peeking at the front of the queue.
func TestPeek(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	q.Enqueue(1)
	q.Enqueue(2)

	if got, ok := q.Peek(); !ok || got != 1 {
		t.Errorf("Peek() = %d, %v; want 1, true", got, ok)
	}

	// Ensure Peek does not remove the item
	if got, ok := q.Peek(); !ok || got != 1 {
		t.Errorf("Peek() = %d, %v; want 1, true after re-peeking", got, ok)
	}
}

// TestPeekEmpty tests peeking into an empty queue.
func TestPeekEmpty(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	if _, ok := q.Peek(); ok {
		t.Errorf("Peek() should return false for an empty queue")
	}
}

// TestLen tests the length of the queue.
func TestLen(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	if got := q.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0", got)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	if got := q.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}
}

// TestIsEmpty tests checking if the queue is empty.
func TestIsEmpty(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true after Dequeue")
	}
}

// TestClear tests clearing the queue.
func TestClear(t *testing.T) {
	// Create a new queue with an initial capacity of 10
	q := New[int](10)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Clear()

	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true after Clear")
	}
	if got := q.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0 after Clear", got)
	}
}
