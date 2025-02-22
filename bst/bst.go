package bst

import "github.com/idsulik/go-collections/v3/internal/cmp"

// BST represents the Binary Search Tree.
type BST[T cmp.Ordered] struct {
	root *node[T]
	size int
}

// node represents each node in the BST.
type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

// New creates a new empty Binary Search Tree.
func New[T cmp.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Insert adds a value into the BST.
func (bst *BST[T]) Insert(value T) {
	bst.root = bst.insert(bst.root, value)
}

func (bst *BST[T]) insert(n *node[T], value T) *node[T] {
	if n == nil {
		bst.size++
		return &node[T]{value: value}
	}

	cur := n
	for cur != nil {
		if value < cur.value {
			if cur.left == nil {
				bst.size++
				cur.left = &node[T]{value: value}
				break
			}
			cur = cur.left
		} else if value > cur.value {
			if cur.right == nil {
				bst.size++
				cur.right = &node[T]{value: value}
				break
			}
			cur = cur.right
		} else {
			// Value already exists
			break
		}
	}

	return n
}

// Contains checks if a value exists in the BST.
func (bst *BST[T]) Contains(value T) bool {
	return bst.contains(bst.root, value)
}

func (bst *BST[T]) contains(n *node[T], value T) bool {
	cur := n
	for cur != nil {
		if value < cur.value {
			cur = cur.left
		} else if value > cur.value {
			cur = cur.right
		} else {
			return true
		}
	}

	return false
}

// Remove deletes a value from the BST.
func (bst *BST[T]) Remove(value T) {
	var removed bool
	bst.root, removed = bst.remove(bst.root, value)
	if removed {
		bst.size--
	}
}

func (bst *BST[T]) remove(n *node[T], value T) (*node[T], bool) {
	if n == nil {
		return nil, false
	}

	var removed bool
	if value < n.value {
		n.left, removed = bst.remove(n.left, value)
	} else if value > n.value {
		n.right, removed = bst.remove(n.right, value)
	} else {
		// Node found, remove it
		removed = true
		if n.left == nil {
			return n.right, removed
		} else if n.right == nil {
			return n.left, removed
		} else {
			// Node with two children
			minRight := bst.min(n.right)
			n.value = minRight.value
			n.right, _ = bst.remove(n.right, n.value)
		}
	}

	return n, removed
}

func (bst *BST[T]) min(n *node[T]) *node[T] {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

// InOrderTraversal traverses the BST in order and applies the function fn to each node's value.
func (bst *BST[T]) InOrderTraversal(fn func(T)) {
	bst.inOrderTraversal(bst.root, fn)
}

func (bst *BST[T]) inOrderTraversal(n *node[T], fn func(T)) {
	if n != nil {
		bst.inOrderTraversal(n.left, fn)
		fn(n.value)
		bst.inOrderTraversal(n.right, fn)
	}
}

// Len returns the number of nodes in the BST.
func (bst *BST[T]) Len() int {
	return bst.size
}

// IsEmpty checks if the BST is empty.
func (bst *BST[T]) IsEmpty() bool {
	return bst.size == 0
}

// Clear removes all nodes from the BST.
func (bst *BST[T]) Clear() {
	bst.root = nil
	bst.size = 0
}
