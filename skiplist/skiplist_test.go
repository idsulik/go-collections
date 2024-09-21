package skiplist

import (
	"testing"
)

// Test basic insertion and search
func TestInsertAndSearch(t *testing.T) {
	sl := New[int](16, 0.5)
	values := []int{3, 6, 7, 9, 12, 19, 17, 26, 21, 25}

	for _, v := range values {
		sl.Insert(v)
	}

	for _, v := range values {
		if !sl.Search(v) {
			t.Errorf("SkipList should contain %d", v)
		}
	}

	if sl.Search(100) {
		t.Errorf("SkipList should not contain %d", 100)
	}
}

// Test deletion
func TestDelete(t *testing.T) {
	sl := New[int](16, 0.5)
	values := []int{3, 6, 7, 9, 12, 19, 17}

	for _, v := range values {
		sl.Insert(v)
	}

	sl.Delete(6)
	if sl.Search(6) {
		t.Errorf("SkipList should not contain %d after deletion", 6)
	}

	sl.Delete(9)
	if sl.Search(9) {
		t.Errorf("SkipList should not contain %d after deletion", 9)
	}

	sl.Delete(100) // Deleting non-existing element should not cause error
}

// Test Len and IsEmpty
func TestLenAndIsEmpty(t *testing.T) {
	sl := New[int](16, 0.5)

	if !sl.IsEmpty() {
		t.Error("SkipList should be empty")
	}

	if sl.Len() != 0 {
		t.Errorf("Expected length 0, got %d", sl.Len())
	}

	sl.Insert(1)

	if sl.IsEmpty() {
		t.Error("SkipList should not be empty after insertion")
	}

	if sl.Len() != 1 {
		t.Errorf("Expected length 1, got %d", sl.Len())
	}
}

// Test Clear
func TestClear(t *testing.T) {
	sl := New[int](16, 0.5)
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(3)

	sl.Clear()

	if !sl.IsEmpty() {
		t.Error("SkipList should be empty after Clear")
	}

	if sl.Len() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", sl.Len())
	}
}

// Test with strings
func TestStrings(t *testing.T) {
	sl := New[string](16, 0.5)
	values := []string{"apple", "banana", "cherry", "date", "elderberry"}

	for _, v := range values {
		sl.Insert(v)
	}

	if !sl.Search("banana") {
		t.Error("SkipList should contain 'banana'")
	}

	sl.Delete("cherry")
	if sl.Search("cherry") {
		t.Error("SkipList should not contain 'cherry' after deletion")
	}

	if sl.Len() != 4 {
		t.Errorf("Expected length 4 after deletion, got %d", sl.Len())
	}
}

// Test inserting duplicate values
func TestInsertDuplicate(t *testing.T) {
	sl := New[int](16, 0.5)
	sl.Insert(10)
	sl.Insert(10)

	if sl.Len() != 1 {
		t.Errorf("Expected length 1 after inserting duplicate, got %d", sl.Len())
	}
}

// Test large dataset
func TestLargeDataSet(t *testing.T) {
	sl := New[int](32, 0.5)
	const numElements = 10000

	for i := 0; i < numElements; i++ {
		sl.Insert(i)
	}

	if sl.Len() != numElements {
		t.Errorf("Expected length %d, got %d", numElements, sl.Len())
	}

	for i := 0; i < numElements; i++ {
		if !sl.Search(i) {
			t.Errorf("SkipList should contain %d", i)
		}
	}
}
