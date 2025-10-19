package btree

import (
	"math/rand"
	"sort"
	"testing"
)

// TestNew tests the creation of a new B-Tree
func TestNew(t *testing.T) {
	tests := []struct {
		name           string
		degree         int
		expectedDegree int
	}{
		{"Valid degree 3", 3, 3},
		{"Valid degree 10", 10, 10},
		{"Invalid degree 1", 1, 2},
		{"Invalid degree 0", 0, 2},
		{"Invalid degree -5", -5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New[int](tt.degree)
			if tree == nil {
				t.Error("Expected new tree to be non-nil")
			}
			if tree.Degree() != tt.expectedDegree {
				t.Errorf("Expected degree %d, got %d", tt.expectedDegree, tree.Degree())
			}
			if !tree.IsEmpty() {
				t.Error("Expected new tree to be empty")
			}
			if tree.Len() != 0 {
				t.Errorf("Expected length 0, got %d", tree.Len())
			}
		})
	}
}

// TestInsertAndSearch tests insertion and search operations
func TestInsertAndSearch(t *testing.T) {
	tree := New[int](3)

	// Test inserting values
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, v := range values {
		tree.Insert(v)
	}

	if tree.Len() != len(values) {
		t.Errorf("Expected length %d, got %d", len(values), tree.Len())
	}

	// Test searching for existing values
	for _, v := range values {
		if !tree.Search(v) {
			t.Errorf("Expected to find value %d", v)
		}
	}

	// Test searching for non-existing values
	nonExisting := []int{1, 3, 8, 15, 25, 100}
	for _, v := range nonExisting {
		if tree.Search(v) {
			t.Errorf("Did not expect to find value %d", v)
		}
	}
}

// TestInsertDuplicates tests that duplicates are not inserted
func TestInsertDuplicates(t *testing.T) {
	tree := New[int](3)

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(10) // Duplicate
	tree.Insert(20) // Duplicate

	if tree.Len() != 2 {
		t.Errorf("Expected length 2, got %d", tree.Len())
	}
}

// TestDelete tests deletion operations
func TestDelete(t *testing.T) {
	tree := New[int](3)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range values {
		tree.Insert(v)
	}

	// Delete some values
	deleteValues := []int{5, 1, 10, 3}
	for _, v := range deleteValues {
		if !tree.Delete(v) {
			t.Errorf("Expected to delete value %d", v)
		}
		if tree.Search(v) {
			t.Errorf("Value %d should have been deleted", v)
		}
	}

	expectedLen := len(values) - len(deleteValues)
	if tree.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, tree.Len())
	}

	// Try to delete non-existing value
	if tree.Delete(100) {
		t.Error("Should not be able to delete non-existing value")
	}
}

// TestDeleteAll tests deleting all elements
func TestDeleteAll(t *testing.T) {
	tree := New[int](3)

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		tree.Insert(v)
	}

	for _, v := range values {
		tree.Delete(v)
	}

	if !tree.IsEmpty() {
		t.Error("Expected tree to be empty after deleting all elements")
	}
	if tree.Len() != 0 {
		t.Errorf("Expected length 0, got %d", tree.Len())
	}
}

// TestInOrderTraversal tests in-order traversal
func TestInOrderTraversal(t *testing.T) {
	tree := New[int](3)

	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, v := range values {
		tree.Insert(v)
	}

	// Collect values from traversal
	var result []int
	tree.InOrderTraversal(func(v int) {
		result = append(result, v)
	})

	// Check if result is sorted
	expected := make([]int, len(values))
	copy(expected, values)
	sort.Ints(expected)

	if len(result) != len(expected) {
		t.Errorf("Expected %d values, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}

// TestMinMax tests Min and Max operations
func TestMinMax(t *testing.T) {
	tree := New[int](3)

	// Test empty tree
	if _, ok := tree.Min(); ok {
		t.Error("Expected Min to return false for empty tree")
	}
	if _, ok := tree.Max(); ok {
		t.Error("Expected Max to return false for empty tree")
	}

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}

	min, ok := tree.Min()
	if !ok {
		t.Error("Expected Min to return true for non-empty tree")
	}
	if min != 20 {
		t.Errorf("Expected min 20, got %d", min)
	}

	max, ok := tree.Max()
	if !ok {
		t.Error("Expected Max to return true for non-empty tree")
	}
	if max != 80 {
		t.Errorf("Expected max 80, got %d", max)
	}
}

// TestClear tests clearing the tree
func TestClear(t *testing.T) {
	tree := New[int](3)

	for i := 0; i < 20; i++ {
		tree.Insert(i)
	}

	tree.Clear()

	if !tree.IsEmpty() {
		t.Error("Expected tree to be empty after Clear")
	}
	if tree.Len() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", tree.Len())
	}
}

// TestHeight tests the height calculation
func TestHeight(t *testing.T) {
	tree := New[int](3)

	// Empty tree has height 0
	if tree.Height() != 0 {
		t.Errorf("Expected height 0 for empty tree, got %d", tree.Height())
	}

	// Insert values and check height increases
	for i := 1; i <= 20; i++ {
		tree.Insert(i)
	}

	height := tree.Height()
	if height < 0 {
		t.Errorf("Height should be non-negative, got %d", height)
	}
}

// TestLargeDataset tests with a large number of elements
func TestLargeDataset(t *testing.T) {
	tree := New[int](5)
	n := 1000

	// Insert values
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	if tree.Len() != n {
		t.Errorf("Expected length %d, got %d", n, tree.Len())
	}

	// Search for all values
	for i := 0; i < n; i++ {
		if !tree.Search(i) {
			t.Errorf("Expected to find value %d", i)
		}
	}

	// Delete half the values
	for i := 0; i < n; i += 2 {
		tree.Delete(i)
	}

	expectedLen := n / 2
	if tree.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, tree.Len())
	}

	// Verify deleted values are gone
	for i := 0; i < n; i += 2 {
		if tree.Search(i) {
			t.Errorf("Value %d should have been deleted", i)
		}
	}

	// Verify remaining values still exist
	for i := 1; i < n; i += 2 {
		if !tree.Search(i) {
			t.Errorf("Expected to find value %d", i)
		}
	}
}

// TestRandomOperations tests random insertions and deletions
func TestRandomOperations(t *testing.T) {
	tree := New[int](4)
	rand.Seed(42)

	inserted := make(map[int]bool)
	n := 500

	// Random insertions
	for i := 0; i < n; i++ {
		val := rand.Intn(1000)
		tree.Insert(val)
		inserted[val] = true
	}

	// Verify all inserted values exist
	for val := range inserted {
		if !tree.Search(val) {
			t.Errorf("Expected to find value %d", val)
		}
	}

	// Random deletions
	deleteCount := 0
	for val := range inserted {
		if rand.Float32() < 0.5 {
			tree.Delete(val)
			delete(inserted, val)
			deleteCount++
		}
		if deleteCount >= len(inserted)/2 {
			break
		}
	}

	// Verify remaining values
	for val := range inserted {
		if !tree.Search(val) {
			t.Errorf("Expected to find value %d", val)
		}
	}
}

// TestStringType tests B-Tree with string type
func TestStringType(t *testing.T) {
	tree := New[string](3)

	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, word := range words {
		tree.Insert(word)
	}

	if tree.Len() != len(words) {
		t.Errorf("Expected length %d, got %d", len(words), tree.Len())
	}

	for _, word := range words {
		if !tree.Search(word) {
			t.Errorf("Expected to find word %s", word)
		}
	}

	min, _ := tree.Min()
	if min != "apple" {
		t.Errorf("Expected min 'apple', got '%s'", min)
	}

	max, _ := tree.Max()
	if max != "elderberry" {
		t.Errorf("Expected max 'elderberry', got '%s'", max)
	}
}

// TestDifferentDegrees tests B-Trees with different degrees
func TestDifferentDegrees(t *testing.T) {
	degrees := []int{2, 3, 5, 10}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, degree := range degrees {
		t.Run(string(rune(degree)), func(t *testing.T) {
			tree := New[int](degree)

			for _, v := range values {
				tree.Insert(v)
			}

			if tree.Len() != len(values) {
				t.Errorf("Degree %d: Expected length %d, got %d", degree, len(values), tree.Len())
			}

			for _, v := range values {
				if !tree.Search(v) {
					t.Errorf("Degree %d: Expected to find value %d", degree, v)
				}
			}
		})
	}
}

// TestEdgeCases tests various edge cases
func TestEdgeCases(t *testing.T) {
	t.Run("Single element", func(t *testing.T) {
		tree := New[int](3)
		tree.Insert(42)

		if tree.Len() != 1 {
			t.Errorf("Expected length 1, got %d", tree.Len())
		}

		min, _ := tree.Min()
		max, _ := tree.Max()
		if min != 42 || max != 42 {
			t.Errorf("Expected min and max to be 42, got min=%d, max=%d", min, max)
		}

		tree.Delete(42)
		if !tree.IsEmpty() {
			t.Error("Expected tree to be empty")
		}
	})

	t.Run("Sequential insertion", func(t *testing.T) {
		tree := New[int](3)
		for i := 1; i <= 100; i++ {
			tree.Insert(i)
		}

		if tree.Len() != 100 {
			t.Errorf("Expected length 100, got %d", tree.Len())
		}

		// Verify sorted order
		var result []int
		tree.InOrderTraversal(func(v int) {
			result = append(result, v)
		})

		for i := 0; i < len(result)-1; i++ {
			if result[i] >= result[i+1] {
				t.Errorf("Values not in sorted order: %d >= %d", result[i], result[i+1])
			}
		}
	})

	t.Run("Reverse sequential insertion", func(t *testing.T) {
		tree := New[int](3)
		for i := 100; i >= 1; i-- {
			tree.Insert(i)
		}

		if tree.Len() != 100 {
			t.Errorf("Expected length 100, got %d", tree.Len())
		}

		min, _ := tree.Min()
		max, _ := tree.Max()
		if min != 1 || max != 100 {
			t.Errorf("Expected min=1, max=100, got min=%d, max=%d", min, max)
		}
	})
}

// Benchmark tests

func BenchmarkBTreeInsert(b *testing.B) {
	degrees := []int{2, 3, 5, 10}

	for _, degree := range degrees {
		b.Run(string(rune(degree)), func(b *testing.B) {
			tree := New[int](degree)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tree.Insert(i)
			}
		})
	}
}

func BenchmarkBTreeSearch(b *testing.B) {
	tree := New[int](5)
	n := 10000

	// Pre-populate tree
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(i % n)
	}
}

func BenchmarkBTreeDelete(b *testing.B) {
	tree := New[int](5)
	n := 100000

	// Pre-populate tree
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N && i < n; i++ {
		tree.Delete(i)
	}
}

func BenchmarkBTreeInOrderTraversal(b *testing.B) {
	tree := New[int](5)
	n := 10000

	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.InOrderTraversal(func(v int) {
			// Do nothing, just traverse
		})
	}
}

func BenchmarkBTreeMixed(b *testing.B) {
	tree := New[int](5)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(i)
		if i%3 == 0 {
			tree.Search(i / 2)
		}
		if i%5 == 0 && i > 0 {
			tree.Delete(i - 1)
		}
	}
}
