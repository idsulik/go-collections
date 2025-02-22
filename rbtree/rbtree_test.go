package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/idsulik/go-collections/v3/internal/cmp"
)

// verifyRedBlackProperties checks if the tree maintains Red-Black properties
func verifyRedBlackProperties[T any](t *RedBlackTree[T]) bool {
	if t.root == nil {
		return true
	}

	// Property 1: Root must be black
	if t.root.color != Black {
		return false
	}

	// Check other properties recursively
	blackHeight, valid := verifyNodeProperties(t.root, nil)
	return valid && blackHeight >= 0
}

// verifyNodeProperties checks Red-Black properties for a node and its subtrees
func verifyNodeProperties[T any](n *node[T], parent *node[T]) (int, bool) {
	if n == nil {
		return 0, true // Nil nodes are considered black
	}

	// Check parent pointer
	if n.parent != parent {
		return -1, false
	}

	// Property 2: No red node has a red child
	if n.color == Red && parent != nil && parent.color == Red {
		return -1, false
	}

	// Check left subtree
	leftBlackHeight, leftValid := verifyNodeProperties(n.left, n)
	if !leftValid {
		return -1, false
	}

	// Check right subtree
	rightBlackHeight, rightValid := verifyNodeProperties(n.right, n)
	if !rightValid {
		return -1, false
	}

	// Property 5: All paths must have same number of black nodes
	if leftBlackHeight != rightBlackHeight {
		return -1, false
	}

	// Calculate black height
	blackHeight := leftBlackHeight
	if n.color == Black {
		blackHeight++
	}

	return blackHeight, true
}

func TestNewRedBlackTree(t *testing.T) {
	tree := New[int](cmp.CompareInts)
	if tree == nil {
		t.Error("Expected non-nil tree")
	}
	if !tree.IsEmpty() {
		t.Error("Expected empty tree")
	}
	if tree.Len() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Len())
	}
}

func TestRedBlackTree_Insert(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{"Empty", []int{}},
		{"Single Value", []int{1}},
		{"Ascending Order", []int{1, 2, 3, 4, 5}},
		{"Descending Order", []int{5, 4, 3, 2, 1}},
		{"Random Order", []int{3, 1, 4, 5, 2}},
		{"Duplicates", []int{1, 2, 2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tree := New[int](cmp.CompareInts)
				uniqueValues := make(map[int]bool)

				for _, v := range tt.values {
					tree.Insert(v)
					uniqueValues[v] = true

					// Verify Red-Black properties after each insertion
					if !verifyRedBlackProperties(tree) {
						t.Error("Red-Black properties violated after insertion")
					}
				}

				// Check size
				expectedSize := len(uniqueValues)
				if tree.Len() != expectedSize {
					t.Errorf("Expected size %d, got %d", expectedSize, tree.Len())
				}

				// Verify all values are present
				for v := range uniqueValues {
					if !tree.Search(v) {
						t.Errorf("Value %d not found after insertion", v)
					}
				}
			},
		)
	}
}

func TestRedBlackTree_Delete(t *testing.T) {
	tests := []struct {
		name         string
		insertOrder  []int
		deleteOrder  []int
		expectedSize int
	}{
		{
			name:         "Delete Root",
			insertOrder:  []int{1},
			deleteOrder:  []int{1},
			expectedSize: 0,
		},
		{
			name:         "Delete Leaf",
			insertOrder:  []int{2, 1, 3},
			deleteOrder:  []int{1},
			expectedSize: 2,
		},
		{
			name:         "Delete Internal Node",
			insertOrder:  []int{2, 1, 3, 4},
			deleteOrder:  []int{3},
			expectedSize: 3,
		},
		{
			name:         "Delete All",
			insertOrder:  []int{1, 2, 3, 4, 5},
			deleteOrder:  []int{1, 2, 3, 4, 5},
			expectedSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tree := New[int](cmp.CompareInts)

				// Insert values
				for _, v := range tt.insertOrder {
					tree.Insert(v)
				}

				// Delete values
				for _, v := range tt.deleteOrder {
					if !tree.Delete(v) {
						t.Errorf("Failed to delete value %d", v)
					}

					// Verify Red-Black properties after each deletion
					if !verifyRedBlackProperties(tree) {
						t.Error("Red-Black properties violated after deletion")
					}
				}

				// Check final size
				if tree.Len() != tt.expectedSize {
					t.Errorf("Expected size %d, got %d", tt.expectedSize, tree.Len())
				}

				// Verify deleted values are gone
				for _, v := range tt.deleteOrder {
					if tree.Search(v) {
						t.Errorf("Value %d still present after deletion", v)
					}
				}
			},
		)
	}
}

func TestRedBlackTree_InOrderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{
			name:     "Empty Tree",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "Single Node",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Multiple Nodes",
			values:   []int{5, 3, 7, 1, 9, 4, 6, 8, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tree := New[int](cmp.CompareInts)
				for _, v := range tt.values {
					tree.Insert(v)
				}

				var result []int
				tree.InOrderTraversal(
					func(v int) {
						result = append(result, v)
					},
				)

				if len(result) != len(tt.expected) {
					t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				}

				for i := range result {
					if result[i] != tt.expected[i] {
						t.Errorf("Expected %v at index %d, got %v", tt.expected[i], i, result[i])
					}
				}
			},
		)
	}
}

func TestRedBlackTree_Height(t *testing.T) {
	tests := []struct {
		name           string
		values         []int
		expectedHeight int
	}{
		{"Empty Tree", []int{}, -1},
		{"Single Node", []int{1}, 0},
		{"Two Nodes", []int{1, 2}, 1},
		{"Three Nodes", []int{2, 1, 3}, 1},
		{"Multiple Nodes", []int{5, 3, 7, 1, 9, 4, 6, 8, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tree := New[int](cmp.CompareInts)
				for _, v := range tt.values {
					tree.Insert(v)
				}

				height := tree.Height()
				if height != tt.expectedHeight {
					t.Errorf("Expected height %d, got %d", tt.expectedHeight, height)
				}
			},
		)
	}
}

func TestRedBlackTree_Clear(t *testing.T) {
	tree := New[int](cmp.CompareInts)
	values := []int{5, 3, 7, 1, 9}

	for _, v := range values {
		tree.Insert(v)
	}

	tree.Clear()

	if !tree.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}
	if tree.Len() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", tree.Len())
	}
	if tree.Height() != -1 {
		t.Errorf("Expected height -1 after Clear(), got %d", tree.Height())
	}

	// Verify no values remain
	for _, v := range values {
		if tree.Search(v) {
			t.Errorf("Value %d still present after Clear()", v)
		}
	}
}

func TestRedBlackTree_RandomOperations(t *testing.T) {
	tree := New[int](cmp.CompareInts)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	operations := 1000
	maxValue := 100
	values := make(map[int]bool)

	for i := 0; i < operations; i++ {
		value := rng.Intn(maxValue)
		if rng.Float32() < 0.7 { // 70% insertions
			tree.Insert(value)
			values[value] = true
		} else { // 30% deletions
			tree.Delete(value)
			delete(values, value)
		}

		// Verify Red-Black properties
		if !verifyRedBlackProperties(tree) {
			t.Errorf("Red-Black properties violated after operation %d", i)
		}

		// Verify size matches unique values
		if tree.Len() != len(values) {
			t.Errorf(
				"Len mismatch at operation %d: expected %d, got %d",
				i, len(values), tree.Len(),
			)
		}

		// Verify all values are present
		for v := range values {
			if !tree.Search(v) {
				t.Errorf("Value %d missing at operation %d", v, i)
			}
		}
	}
}

func BenchmarkRedBlackTree(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small", 100},
		{"Medium", 1000},
		{"Large", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(
			fmt.Sprintf("Insert_%s", bm.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tree := New[int](cmp.CompareInts)
					for j := 0; j < bm.size; j++ {
						tree.Insert(j)
					}
				}
			},
		)

		b.Run(
			fmt.Sprintf("Search_%s", bm.name), func(b *testing.B) {
				tree := New[int](cmp.CompareInts)
				for i := 0; i < bm.size; i++ {
					tree.Insert(i)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					tree.Search(rand.Intn(bm.size))
				}
			},
		)

		b.Run(
			fmt.Sprintf("Delete_%s", bm.name), func(b *testing.B) {
				values := make([]int, bm.size)
				for i := range values {
					values[i] = i
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					tree := New[int](cmp.CompareInts)
					for _, v := range values {
						tree.Insert(v)
					}
					b.StartTimer()
					for _, v := range values {
						tree.Delete(v)
					}
				}
			},
		)
	}
}
