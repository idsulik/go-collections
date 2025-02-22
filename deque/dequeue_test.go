package deque

import (
	"math"
	"testing"

	"github.com/idsulik/go-collections/v3/internal/slices"
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

func TestClearBehavior(t *testing.T) {
	t.Run(
		"capacity reset on clear", func(t *testing.T) {
			d := New[int](2)
			// Force several capacity increases
			for i := 0; i < 100; i++ {
				d.PushBack(i)
			}

			largeCapacity := d.Cap()
			d.Clear()

			if d.Cap() >= largeCapacity {
				t.Error("Capacity should be reduced after Clear()")
			}
			if d.Cap() > defaultCapacity*2 {
				t.Error("Capacity should be closer to default after Clear()")
			}
		},
	)
}

func TestCapacity(t *testing.T) {
	t.Run(
		"initial capacity", func(t *testing.T) {
			d := New[int](5)
			if got := d.Cap(); got != 5 {
				t.Errorf("Cap() = %d; want 5", got)
			}
		},
	)

	t.Run(
		"default capacity", func(t *testing.T) {
			d := New[int](0)
			if got := d.Cap(); got != defaultCapacity {
				t.Errorf("Cap() = %d; want %d", got, defaultCapacity)
			}
		},
	)

	t.Run(
		"capacity growth", func(t *testing.T) {
			d := New[int](2)
			initialCap := d.Cap()

			// Fill beyond initial capacity
			for i := 0; i < 3; i++ {
				d.PushBack(i)
			}

			if got := d.Cap(); got != initialCap*resizeFactor {
				t.Errorf("Cap() after growth = %d; want %d", got, initialCap*resizeFactor)
			}
		},
	)
}

func TestGetItems(t *testing.T) {
	d := New[int](4)
	expected := []int{1, 2, 3, 4}

	for _, v := range expected {
		d.PushBack(v)
	}

	items := d.GetItems()

	if len(items) != len(expected) {
		t.Errorf("GetItems() length = %d; want %d", len(items), len(expected))
	}

	for i, v := range expected {
		if items[i] != v {
			t.Errorf("GetItems()[%d] = %d; want %d", i, items[i], v)
		}
	}
}

func TestClone(t *testing.T) {
	d := New[int](4)
	original := []int{1, 2, 3}

	for _, v := range original {
		d.PushBack(v)
	}

	clone := d.Clone()

	t.Run(
		"identical content", func(t *testing.T) {
			if clone.Len() != d.Len() {
				t.Errorf("Clone length = %d; want %d", clone.Len(), d.Len())
			}

			for i := 0; i < d.Len(); i++ {
				orig, _ := d.PopFront()
				cloned, _ := clone.PopFront()
				if orig != cloned {
					t.Errorf("Clone mismatch at position %d: got %d; want %d", i, cloned, orig)
				}
			}
		},
	)

	t.Run(
		"independent modification", func(t *testing.T) {
			d.PushBack(4)
			if d.Len() == clone.Len() {
				t.Error("Clone should be independent of original")
			}
		},
	)
}

func TestWraparound(t *testing.T) {
	d := New[int](4)

	// Fill the deque
	for i := 0; i < 4; i++ {
		d.PushBack(i)
	}

	// Create wrap-around by removing from front and adding to back
	d.PopFront()
	d.PopFront()
	d.PushBack(4)
	d.PushBack(5)

	expected := []int{2, 3, 4, 5}
	items := d.GetItems()

	for i, v := range expected {
		if items[i] != v {
			t.Errorf("Wraparound error: items[%d] = %d; want %d", i, items[i], v)
		}
	}
}

func TestOverflowProtection(t *testing.T) {
	t.Run(
		"new with large capacity", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Expected panic for too large initial capacity")
				}
			}()

			New[int](math.MaxInt)
		},
	)
}

func TestEdgeCases(t *testing.T) {
	t.Run(
		"rapid push/pop alternation", func(t *testing.T) {
			d := New[int](2)
			for i := 0; i < 1000; i++ {
				d.PushBack(i)
				if v, ok := d.PopFront(); !ok || v != i {
					t.Errorf("Push/Pop alternation failed at i=%d", i)
				}
			}
		},
	)

	t.Run(
		"mixed front/back operations", func(t *testing.T) {
			d := New[int](4)
			d.PushFront(1)
			d.PushBack(2)
			d.PushFront(3)
			d.PushBack(4)

			expected := []int{3, 1, 2, 4}
			items := d.GetItems()

			for i, v := range expected {
				if items[i] != v {
					t.Errorf("Mixed operations: items[%d] = %d; want %d", i, items[i], v)
				}
			}
		},
	)
}

func TestForEach(t *testing.T) {
	d := New[int](4)

	// Test empty deque
	var emptyResult []int
	d.ForEach(
		func(value int) {
			emptyResult = append(emptyResult, value)
		},
	)
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty result, got %v", emptyResult)
	}

	// Test non-wrapped deque
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	var result []int
	d.ForEach(
		func(value int) {
			result = append(result, value)
		},
	)
	expected := []int{1, 2, 3}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test wrapped deque
	d.PushFront(0)
	d.PushBack(4)
	result = nil
	d.ForEach(
		func(value int) {
			result = append(result, value)
		},
	)
	expected = []int{0, 1, 2, 3, 4}
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
