// Package rbtree implements a Red-Black Tree data structure
package rbtree

// color represents the color of a node in the Red-Black tree
type color bool

const (
	Black color = true
	Red   color = false
)

// Node represents a node in the Red-Black tree
type node[T any] struct {
	value  T
	color  color
	left   *node[T]
	right  *node[T]
	parent *node[T]
}

// RedBlackTree represents a Red-Black tree data structure
type RedBlackTree[T any] struct {
	root    *node[T]
	size    int
	compare func(a, b T) int
}

// New creates a new Red-Black tree
func New[T any](compare func(a, b T) int) *RedBlackTree[T] {
	return &RedBlackTree[T]{
		compare: compare,
	}
}

// Len returns the number of nodes in the tree
func (t *RedBlackTree[T]) Len() int {
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *RedBlackTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Clear removes all nodes from the tree
func (t *RedBlackTree[T]) Clear() {
	t.root = nil
	t.size = 0
}

// Insert adds a value to the tree if it doesn't already exist
func (t *RedBlackTree[T]) Insert(value T) {
	// First check if value already exists
	if t.Search(value) {
		return // Don't insert duplicates
	}

	newNode := &node[T]{
		value: value,
		color: Red,
	}

	if t.root == nil {
		t.root = newNode
		t.size++
		t.insertFixup(newNode)
		return
	}

	current := t.root
	var parent *node[T]

	for current != nil {
		parent = current
		cmp := t.compare(value, current.value)
		if cmp == 0 {
			return // Double-check for duplicates
		} else if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	newNode.parent = parent
	if t.compare(value, parent.value) < 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	t.size++
	t.insertFixup(newNode)
}

// insertFixup maintains Red-Black properties after insertion
func (t *RedBlackTree[T]) insertFixup(n *node[T]) {
	if n.parent == nil {
		n.color = Black
		return
	}

	for n.parent != nil && n.parent.color == Red {
		if n.parent == n.parent.parent.left {
			uncle := n.parent.parent.right
			if uncle != nil && uncle.color == Red {
				n.parent.color = Black
				uncle.color = Black
				n.parent.parent.color = Red
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					t.rotateLeft(n)
				}
				n.parent.color = Black
				n.parent.parent.color = Red
				t.rotateRight(n.parent.parent)
			}
		} else {
			uncle := n.parent.parent.left
			if uncle != nil && uncle.color == Red {
				n.parent.color = Black
				uncle.color = Black
				n.parent.parent.color = Red
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rotateRight(n)
				}
				n.parent.color = Black
				n.parent.parent.color = Red
				t.rotateLeft(n.parent.parent)
			}
		}
		if n == t.root {
			break
		}
	}
	t.root.color = Black
}

// rotateLeft performs a left rotation around the given node
func (t *RedBlackTree[T]) rotateLeft(x *node[T]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rotateRight performs a right rotation around the given node
func (t *RedBlackTree[T]) rotateRight(y *node[T]) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

// Search checks if a value exists in the tree
func (t *RedBlackTree[T]) Search(value T) bool {
	current := t.root
	for current != nil {
		cmp := t.compare(value, current.value)
		if cmp == 0 {
			return true
		} else if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

// InOrderTraversal visits all nodes in ascending order
func (t *RedBlackTree[T]) InOrderTraversal(fn func(T)) {
	var inorder func(*node[T])
	inorder = func(n *node[T]) {
		if n == nil {
			return
		}
		inorder(n.left)
		fn(n.value)
		inorder(n.right)
	}
	inorder(t.root)
}

// Delete removes a value from the tree
func (t *RedBlackTree[T]) Delete(value T) bool {
	node := t.findNode(value)
	if node == nil {
		return false
	}

	t.deleteNode(node)
	t.size--
	return true
}

// findNode finds the node containing the given value
func (t *RedBlackTree[T]) findNode(value T) *node[T] {
	current := t.root
	for current != nil {
		cmp := t.compare(value, current.value)
		if cmp == 0 {
			return current
		} else if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

// deleteNode removes the given node from the tree
func (t *RedBlackTree[T]) deleteNode(n *node[T]) {
	var x, y *node[T]

	if n.left == nil || n.right == nil {
		y = n
	} else {
		y = t.successor(n)
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	if x != nil {
		x.parent = y.parent
	}

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != n {
		n.value = y.value
	}

	if y.color == Black {
		t.deleteFixup(x, y.parent)
	}
}

// successor returns the next larger node
func (t *RedBlackTree[T]) successor(n *node[T]) *node[T] {
	if n.right != nil {
		return t.minimum(n.right)
	}
	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}
	return y
}

// minimum returns the node with the smallest value in the subtree
func (t *RedBlackTree[T]) minimum(n *node[T]) *node[T] {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

// deleteFixup maintains Red-Black properties after deletion
func (t *RedBlackTree[T]) deleteFixup(n *node[T], parent *node[T]) {
	for n != t.root && (n == nil || n.color == Black) {
		if n == parent.left {
			w := parent.right
			if w.color == Red {
				w.color = Black
				parent.color = Red
				t.rotateLeft(parent)
				w = parent.right
			}
			if (w.left == nil || w.left.color == Black) &&
				(w.right == nil || w.right.color == Black) {
				w.color = Red
				n = parent
				parent = n.parent
			} else {
				if w.right == nil || w.right.color == Black {
					if w.left != nil {
						w.left.color = Black
					}
					w.color = Red
					t.rotateRight(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = Black
				if w.right != nil {
					w.right.color = Black
				}
				t.rotateLeft(parent)
				n = t.root
				break
			}
		} else {
			w := parent.left
			if w.color == Red {
				w.color = Black
				parent.color = Red
				t.rotateRight(parent)
				w = parent.left
			}
			if (w.right == nil || w.right.color == Black) &&
				(w.left == nil || w.left.color == Black) {
				w.color = Red
				n = parent
				parent = n.parent
			} else {
				if w.left == nil || w.left.color == Black {
					if w.right != nil {
						w.right.color = Black
					}
					w.color = Red
					t.rotateLeft(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = Black
				if w.left != nil {
					w.left.color = Black
				}
				t.rotateRight(parent)
				n = t.root
				break
			}
		}
	}
	if n != nil {
		n.color = Black
	}
}

// Height returns the height of the tree
func (t *RedBlackTree[T]) Height() int {
	var height func(*node[T]) int
	height = func(n *node[T]) int {
		if n == nil {
			return -1
		}
		leftHeight := height(n.left)
		rightHeight := height(n.right)
		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}
	return height(t.root)
}
