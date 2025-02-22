package linkedlist

import (
	"testing"
)

func TestAddFront(t *testing.T) {
	list := New[int]()
	list.AddFront(1)
	list.AddFront(2)

	if got := list.Size(); got != 2 {
		t.Errorf("Size() = %d; want 2", got)
	}

	if got, _ := list.RemoveFront(); got != 2 {
		t.Errorf("RemoveFront() = %d; want 2", got)
	}

	if got, _ := list.RemoveFront(); got != 1 {
		t.Errorf("RemoveFront() = %d; want 1", got)
	}
}

func TestAddBack(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	if got := list.Size(); got != 2 {
		t.Errorf("Size() = %d; want 2", got)
	}

	if got, _ := list.RemoveFront(); got != 1 {
		t.Errorf("RemoveFront() = %d; want 1", got)
	}

	if got, _ := list.RemoveFront(); got != 2 {
		t.Errorf("RemoveFront() = %d; want 2", got)
	}
}

func TestRemoveFrontEmpty(t *testing.T) {
	list := New[int]()
	if _, ok := list.RemoveFront(); ok {
		t.Errorf("RemoveFront() should return false on empty list")
	}
}

func TestRemoveBackEmpty(t *testing.T) {
	list := New[int]()
	if _, ok := list.RemoveBack(); ok {
		t.Errorf("RemoveBack() should return false on empty list")
	}
}

func TestRemoveBack(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	if got := list.Size(); got != 2 {
		t.Errorf("Size() = %d; want 2", got)
	}

	if got, _ := list.RemoveBack(); got != 2 {
		t.Errorf("RemoveBack() = %d; want 2", got)
	}

	if got, _ := list.RemoveBack(); got != 1 {
		t.Errorf("RemoveBack() = %d; want 1", got)
	}
}

// TestIterate tests the Iterate method of the LinkedList.
func TestIterate(t *testing.T) {
	// Helper function to use as callback in Iterate
	calledValues := []int{}
	fn := func(value int) bool {
		calledValues = append(calledValues, value)
		return true
	}

	// Create a new linked list and add values
	l := New[int]()
	l.AddBack(1)
	l.AddBack(2)
	l.AddBack(3)

	// Call Iterate
	l.Iterate(fn)

	// Verify that Iterate visited all nodes and called the function with correct values
	expectedValues := []int{1, 2, 3}
	if len(calledValues) != len(expectedValues) {
		t.Fatalf("Expected %d calls, but got %d", len(expectedValues), len(calledValues))
	}
	for i, value := range calledValues {
		if value != expectedValues[i] {
			t.Errorf("At index %d, expected %d but got %d", i, expectedValues[i], value)
		}
	}
}

// TestIterateStopsOnFalse verifies that iteration stops when the function returns false.
func TestIterateStopsOnFalse(t *testing.T) {
	// Helper function to use as callback in Iterate
	var calledValues []int
	fn := func(value int) bool {
		calledValues = append(calledValues, value)
		return value != 2 // Stop iteration when value equals 2
	}

	// Create a new linked list and add values
	l := New[int]()
	l.AddBack(1)
	l.AddBack(2)
	l.AddBack(3)
	l.AddBack(4)

	// Call Iterate
	l.Iterate(fn)

	// Verify that Iterate stopped at value 2
	expectedValues := []int{1, 2}
	if len(calledValues) != len(expectedValues) {
		t.Fatalf("Expected %d calls, but got %d", len(expectedValues), len(calledValues))
	}
	for i, value := range calledValues {
		if value != expectedValues[i] {
			t.Errorf("At index %d, expected %d but got %d", i, expectedValues[i], value)
		}
	}
}

// TestIterateEmptyList verifies that Iterate does nothing on an empty list.
func TestIterateEmptyList(t *testing.T) {
	// Helper function to use as callback in Iterate
	fnCalled := false
	fn := func(value int) bool {
		fnCalled = true
		return true
	}

	// Create a new linked list and do not add any values
	l := New[int]()

	// Call Iterate
	l.Iterate(fn)

	// Verify that Iterate did not call the function
	if fnCalled {
		t.Errorf("Function should not have been called on an empty list")
	}
}

func TestSize(t *testing.T) {
	list := New[int]()
	if got := list.Size(); got != 0 {
		t.Errorf("Size() = %d; want 0", got)
	}

	list.AddBack(1)
	if got := list.Size(); got != 1 {
		t.Errorf("Size() = %d; want 1", got)
	}

	list.AddBack(2)
	if got := list.Size(); got != 2 {
		t.Errorf("Size() = %d; want 2", got)
	}

	list.RemoveFront()
	if got := list.Size(); got != 1 {
		t.Errorf("Size() = %d; want 1", got)
	}

	list.RemoveFront()
	if got := list.Size(); got != 0 {
		t.Errorf("Size() = %d; want 0", got)
	}
}

func TestClear(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	list.Clear()
	if got := list.Size(); got != 0 {
		t.Errorf("Size() = %d; want 0", got)
	}
}

func TestForEach(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)

	var sum int
	list.ForEach(
		func(value int) {
			sum += value
		},
	)

	if sum != 6 {
		t.Errorf("ForEach() = %d; want 6", sum)
	}
}
