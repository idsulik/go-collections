package bst

import (
	"testing"
)

func TestInsertAndContains(t *testing.T) {
	bst := New[int]()
	values := []int{5, 3, 7, 2, 4, 6, 8}

	for _, v := range values {
		bst.Insert(v)
	}

	for _, v := range values {
		if !bst.Contains(v) {
			t.Errorf("BST should contain %d", v)
		}
	}

	if bst.Contains(10) {
		t.Errorf("BST should not contain %d", 10)
	}
}

func TestRemove(t *testing.T) {
	bst := New[int]()
	values := []int{5, 3, 7, 2, 4, 6, 8}

	for _, v := range values {
		bst.Insert(v)
	}

	bst.Remove(3)
	if bst.Contains(3) {
		t.Errorf("BST should not contain %d after removal", 3)
	}

	bst.Remove(5)
	if bst.Contains(5) {
		t.Errorf("BST should not contain %d after removal", 5)
	}

	bst.Remove(10) // Removing non-existing element should not cause error
}

func TestInOrderTraversal(t *testing.T) {
	bst := New[int]()
	values := []int{5, 3, 7, 2, 4, 6, 8}

	for _, v := range values {
		bst.Insert(v)
	}

	var traversed []int
	bst.InOrderTraversal(
		func(value int) {
			traversed = append(traversed, value)
		},
	)

	expected := []int{2, 3, 4, 5, 6, 7, 8}

	for i, v := range expected {
		if traversed[i] != v {
			t.Errorf("Expected %d, got %d", v, traversed[i])
		}
	}
}

func TestLenAndIsEmpty(t *testing.T) {
	bst := New[int]()

	if !bst.IsEmpty() {
		t.Error("BST should be empty")
	}

	if bst.Len() != 0 {
		t.Errorf("Expected length 0, got %d", bst.Len())
	}

	bst.Insert(1)

	if bst.IsEmpty() {
		t.Error("BST should not be empty after insertion")
	}

	if bst.Len() != 1 {
		t.Errorf("Expected length 1, got %d", bst.Len())
	}
}

func TestClear(t *testing.T) {
	bst := New[int]()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)

	bst.Clear()

	if !bst.IsEmpty() {
		t.Error("BST should be empty after Clear")
	}

	if bst.Len() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", bst.Len())
	}
}

func TestStrings(t *testing.T) {
	bst := New[string]()
	values := []string{"banana", "apple", "cherry", "date"}

	for _, v := range values {
		bst.Insert(v)
	}

	if !bst.Contains("apple") {
		t.Error("BST should contain 'apple'")
	}

	bst.Remove("banana")
	if bst.Contains("banana") {
		t.Error("BST should not contain 'banana' after removal")
	}

	var traversed []string
	bst.InOrderTraversal(
		func(value string) {
			traversed = append(traversed, value)
		},
	)

	expected := []string{"apple", "cherry", "date"}

	for i, v := range expected {
		if traversed[i] != v {
			t.Errorf("Expected %s, got %s", v, traversed[i])
		}
	}
}

func TestRemoveNodeWithTwoChildren(t *testing.T) {
	bst := New[int]()
	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range values {
		bst.Insert(v)
	}

	bst.Remove(30) // Node with two children

	if bst.Contains(30) {
		t.Error("BST should not contain 30 after removal")
	}

	var traversed []int
	bst.InOrderTraversal(
		func(value int) {
			traversed = append(traversed, value)
		},
	)

	expected := []int{20, 40, 50, 60, 70, 80}

	for i, v := range expected {
		if traversed[i] != v {
			t.Errorf("Expected %d, got %d", v, traversed[i])
		}
	}
}

func TestRemoveRoot(t *testing.T) {
	bst := New[int]()
	values := []int{10, 5, 15, 2, 7, 12, 17}

	for _, v := range values {
		bst.Insert(v)
	}

	bst.Remove(10) // Remove root node

	if bst.Contains(10) {
		t.Error("BST should not contain root node after removal")
	}

	if bst.Len() != 6 {
		t.Errorf("Expected length 6 after removing root, got %d", bst.Len())
	}
}

func TestRemoveLeaf(t *testing.T) {
	bst := New[int]()
	values := []int{10, 5, 15, 2, 7}

	for _, v := range values {
		bst.Insert(v)
	}

	bst.Remove(2) // Remove leaf node

	if bst.Contains(2) {
		t.Error("BST should not contain 2 after removal")
	}

	if bst.Len() != 4 {
		t.Errorf("Expected length 4 after removing leaf, got %d", bst.Len())
	}
}

// Test inserting duplicate values
func TestInsertDuplicate(t *testing.T) {
	bst := New[int]()
	bst.Insert(10)
	bst.Insert(10)

	if bst.Len() != 1 {
		t.Errorf("Expected length 1 after inserting duplicate, got %d", bst.Len())
	}

	count := 0
	bst.InOrderTraversal(
		func(value int) {
			count++
		},
	)

	if count != 1 {
		t.Errorf("Expected traversal count 1, got %d", count)
	}
}
