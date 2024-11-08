package avltree

import (
	"testing"
)

func compareInts(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestAVLTree(t *testing.T) {
	t.Run(
		"Basic Operations", func(t *testing.T) {
			tree := New[int](compareInts)

			// Test Insert and Search
			tree.Insert(5)
			tree.Insert(3)
			tree.Insert(7)

			if !tree.Search(5) {
				t.Error("Expected to find 5 in tree")
			}
			if tree.Search(1) {
				t.Error("Did not expect to find 1 in tree")
			}

			// Test Size
			if size := tree.Size(); size != 3 {
				t.Errorf("Expected size 3, got %d", size)
			}
		},
	)

	t.Run(
		"Balance After Insert", func(t *testing.T) {
			tree := New[int](compareInts)

			// Left-Left case
			tree.Insert(30)
			tree.Insert(20)
			tree.Insert(10)

			if tree.getHeight(tree.root.Left) != 0 || tree.getHeight(tree.root.Right) != 0 {
				t.Error("Tree not properly balanced after left-left case")
			}

			tree.Clear()

			// Right-Right case
			tree.Insert(10)
			tree.Insert(20)
			tree.Insert(30)

			if tree.getHeight(tree.root.Left) != 0 || tree.getHeight(tree.root.Right) != 0 {
				t.Error("Tree not properly balanced after right-right case")
			}
		},
	)

	t.Run(
		"Delete Operations", func(t *testing.T) {
			tree := New[int](compareInts)

			// Insert some values
			values := []int{10, 5, 15, 3, 7, 12, 17}
			for _, v := range values {
				tree.Insert(v)
			}

			// Test deletion
			if !tree.Delete(5) {
				t.Error("Delete should return true for existing value")
			}
			if tree.Delete(100) {
				t.Error("Delete should return false for non-existing value")
			}

			// Verify size after deletion
			if size := tree.Size(); size != 6 {
				t.Errorf("Expected size 6 after deletion, got %d", size)
			}

			// Verify the value is actually deleted
			if tree.Search(5) {
				t.Error("Found deleted value 5 in tree")
			}
		},
	)

	t.Run(
		"InOrder Traversal", func(t *testing.T) {
			tree := New[int](compareInts)
			values := []int{5, 3, 7, 1, 4, 6, 8}
			expected := []int{1, 3, 4, 5, 6, 7, 8}

			for _, v := range values {
				tree.Insert(v)
			}

			result := make([]int, 0)
			tree.InOrderTraversal(
				func(v int) {
					result = append(result, v)
				},
			)

			if len(result) != len(expected) {
				t.Errorf("Expected traversal length %d, got %d", len(expected), len(result))
			}

			for i := range result {
				if result[i] != expected[i] {
					t.Errorf("At position %d, expected %d, got %d", i, expected[i], result[i])
				}
			}
		},
	)

	t.Run(
		"Height Calculation", func(t *testing.T) {
			tree := New[int](compareInts)

			if h := tree.Height(); h != 0 {
				t.Errorf("Expected height 0 for empty tree, got %d", h)
			}

			tree.Insert(1)
			if h := tree.Height(); h != 1 {
				t.Errorf("Expected height 1 for single node, got %d", h)
			}

			tree.Insert(2)
			tree.Insert(3)
			if h := tree.Height(); h != 2 {
				t.Errorf("Expected height 2 after balancing, got %d", h)
			}
		},
	)

	t.Run(
		"Clear and IsEmpty", func(t *testing.T) {
			tree := New[int](compareInts)

			if !tree.IsEmpty() {
				t.Error("New tree should be empty")
			}

			tree.Insert(1)
			tree.Insert(2)

			if tree.IsEmpty() {
				t.Error("Tree should not be empty after insertions")
			}

			tree.Clear()

			if !tree.IsEmpty() {
				t.Error("Tree should be empty after Clear()")
			}

			if size := tree.Size(); size != 0 {
				t.Errorf("Expected size 0 after Clear(), got %d", size)
			}
		},
	)

	t.Run(
		"Complex Balancing", func(t *testing.T) {
			tree := New[int](compareInts)
			values := []int{10, 20, 30, 40, 50, 25}

			for _, v := range values {
				tree.Insert(v)
				if !isBalanced(tree.root) {
					t.Errorf("Tree became unbalanced after inserting %d", v)
				}
			}

			// Delete some values and check balance
			deleteValues := []int{30, 40}
			for _, v := range deleteValues {
				tree.Delete(v)
				if !isBalanced(tree.root) {
					t.Errorf("Tree became unbalanced after deleting %d", v)
				}
			}
		},
	)
}

// Helper function to check if the tree is balanced
func isBalanced(node *Node[int]) bool {
	if node == nil {
		return true
	}

	balance := getNodeHeight(node.Left) - getNodeHeight(node.Right)
	if balance < -1 || balance > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}

func getNodeHeight(node *Node[int]) int {
	if node == nil {
		return -1
	}
	return node.Height
}
