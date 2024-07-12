package deque

import (
	"testing"
)

func TestPushFront(t *testing.T) {
	d := New[int](2)
	d.PushFront(1)
	d.PushFront(2)

	if got := d.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}

	if got, ok := d.PopFront(); !ok || got != 2 {
		t.Errorf("PopFront() = %d, %v; want 2, true", got, ok)
	}

	if got, ok := d.PopFront(); !ok || got != 1 {
		t.Errorf("PopFront() = %d, %v; want 1, true", got, ok)
	}
}

func TestPushBack(t *testing.T) {
	d := New[int](2)
	d.PushBack(1)
	d.PushBack(2)

	if got := d.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}

	if got, ok := d.PopFront(); !ok || got != 1 {
		t.Errorf("PopFront() = %d, %v; want 1, true", got, ok)
	}

	if got, ok := d.PopFront(); !ok || got != 2 {
		t.Errorf("PopFront() = %d, %v; want 2, true", got, ok)
	}
}

func TestPopFrontEmpty(t *testing.T) {
	d := New[int](0)
	if _, ok := d.PopFront(); ok {
		t.Errorf("PopFront() should return false for an empty deque")
	}
}

func TestPopBackEmpty(t *testing.T) {
	d := New[int](0)
	if _, ok := d.PopBack(); ok {
		t.Errorf("PopBack() should return false for an empty deque")
	}
}

func TestPopBack(t *testing.T) {
	d := New[int](0)
	d.PushBack(1)
	d.PushBack(2)

	if got := d.Len(); got != 2 {
		t.Errorf("Len() = %d; want 2", got)
	}

	if got, ok := d.PopBack(); !ok || got != 2 {
		t.Errorf("PopBack() = %d, %v; want 2, true", got, ok)
	}

	if got, ok := d.PopBack(); !ok || got != 1 {
		t.Errorf("PopBack() = %d, %v; want 1, true", got, ok)
	}
}

func TestPeekFront(t *testing.T) {
	d := New[int](0)
	d.PushBack(1)
	d.PushBack(2)

	if got, ok := d.PeekFront(); !ok || got != 1 {
		t.Errorf("PeekFront() = %d, %v; want 1, true", got, ok)
	}
}

func TestPeekBack(t *testing.T) {
	d := New[int](2)
	d.PushBack(1)
	d.PushBack(2)

	if got, ok := d.PeekBack(); !ok || got != 2 {
		t.Errorf("PeekBack() = %d, %v; want 2, true", got, ok)
	}
}

func TestIsEmpty(t *testing.T) {
	d := New[int](1)
	if !d.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	d.PushBack(1)
	if d.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	d.PopFront()
	if !d.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

func TestClear(t *testing.T) {
	d := New[int](2)
	d.PushBack(1)
	d.PushBack(2)
	d.Clear()

	if !d.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true after Clear")
	}

	if got := d.Len(); got != 0 {
		t.Errorf("Len() = %d; want 0 after Clear", got)
	}
}
