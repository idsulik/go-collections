package btree

import "github.com/idsulik/go-collections/v3/internal/cmp"

// BTree represents a B-Tree data structure.
// The degree (t) determines the range of keys a node can contain:
// - Each node can contain at most 2t-1 keys
// - Each node (except root) must contain at least t-1 keys
// - Each internal node can have at most 2t children
type BTree[T cmp.Ordered] struct {
	root   *node[T]
	degree int // minimum degree (t)
	size   int
}

// node represents a node in the B-Tree
type node[T cmp.Ordered] struct {
	keys     []T
	children []*node[T]
	leaf     bool
}

// New creates a new B-Tree with the specified minimum degree.
// The degree must be at least 2. A higher degree means more keys per node.
// Common values: 2-4 for in-memory trees, higher for disk-based trees.
func New[T cmp.Ordered](degree int) *BTree[T] {
	if degree < 2 {
		degree = 2
	}
	return &BTree[T]{
		root:   &node[T]{leaf: true},
		degree: degree,
	}
}

// Insert adds a value to the B-Tree.
// If the value already exists, it will not be added again.
func (t *BTree[T]) Insert(value T) {
	root := t.root

	// If root is full, split it
	if len(root.keys) == 2*t.degree-1 {
		newRoot := &node[T]{leaf: false}
		newRoot.children = append(newRoot.children, t.root)
		t.splitChild(newRoot, 0)
		t.root = newRoot
	}

	t.insertNonFull(t.root, value)
}

// insertNonFull inserts a value into a node that is not full
func (t *BTree[T]) insertNonFull(n *node[T], value T) {
	i := len(n.keys) - 1

	if n.leaf {
		// Check for duplicates first
		for j := 0; j < len(n.keys); j++ {
			if n.keys[j] == value {
				return // Duplicate found, don't insert
			}
		}

		// Insert into leaf node
		n.keys = append(n.keys, value) // Add space
		for i >= 0 && value < n.keys[i] {
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = value
		t.size++
	} else {
		// Find child to insert into
		for i >= 0 && value < n.keys[i] {
			i--
		}
		i++

		// Check if value already exists in current node
		if i > 0 && n.keys[i-1] == value {
			return
		}

		// Split child if full
		if len(n.children[i].keys) == 2*t.degree-1 {
			t.splitChild(n, i)
			if value > n.keys[i] {
				i++
			} else if value == n.keys[i] {
				return
			}
		}
		t.insertNonFull(n.children[i], value)
	}
}

// splitChild splits a full child of a node
func (t *BTree[T]) splitChild(parent *node[T], index int) {
	degree := t.degree
	fullChild := parent.children[index]
	newChild := &node[T]{leaf: fullChild.leaf}

	// Move the second half of keys to new child
	mid := degree - 1
	newChild.keys = make([]T, degree-1)
	copy(newChild.keys, fullChild.keys[degree:])

	// If not a leaf, move the second half of children
	if !fullChild.leaf {
		newChild.children = make([]*node[T], degree)
		copy(newChild.children, fullChild.children[degree:])
		fullChild.children = fullChild.children[:degree]
	}

	// Move middle key up to parent
	parent.keys = append(parent.keys, fullChild.keys[mid])
	copy(parent.keys[index+1:], parent.keys[index:])
	parent.keys[index] = fullChild.keys[mid]

	// Insert new child into parent
	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newChild

	// Truncate the original child
	fullChild.keys = fullChild.keys[:mid]
}

// Search checks if a value exists in the B-Tree
func (t *BTree[T]) Search(value T) bool {
	return t.search(t.root, value)
}

// search recursively searches for a value in the tree
func (t *BTree[T]) search(n *node[T], value T) bool {
	i := 0
	for i < len(n.keys) && value > n.keys[i] {
		i++
	}

	if i < len(n.keys) && value == n.keys[i] {
		return true
	}

	if n.leaf {
		return false
	}

	return t.search(n.children[i], value)
}

// Delete removes a value from the B-Tree
func (t *BTree[T]) Delete(value T) bool {
	if !t.Search(value) {
		return false
	}

	t.delete(t.root, value)

	// If root is empty after deletion, make its only child the new root
	if len(t.root.keys) == 0 && !t.root.leaf {
		t.root = t.root.children[0]
	}

	t.size--
	return true
}

// delete recursively deletes a value from the tree
func (t *BTree[T]) delete(n *node[T], value T) {
	i := 0
	for i < len(n.keys) && value > n.keys[i] {
		i++
	}

	if i < len(n.keys) && value == n.keys[i] {
		// Key found in this node
		if n.leaf {
			t.deleteFromLeaf(n, i)
		} else {
			t.deleteFromNonLeaf(n, i)
		}
	} else if !n.leaf {
		// Key might be in subtree
		isInSubtree := (i == len(n.keys))

		if len(n.children[i].keys) < t.degree {
			t.fill(n, i)
		}

		if isInSubtree && i > len(n.keys) {
			t.delete(n.children[i-1], value)
		} else {
			t.delete(n.children[i], value)
		}
	}
}

// deleteFromLeaf removes a key from a leaf node
func (t *BTree[T]) deleteFromLeaf(n *node[T], index int) {
	copy(n.keys[index:], n.keys[index+1:])
	n.keys = n.keys[:len(n.keys)-1]
}

// deleteFromNonLeaf removes a key from a non-leaf node
func (t *BTree[T]) deleteFromNonLeaf(n *node[T], index int) {
	key := n.keys[index]

	if len(n.children[index].keys) >= t.degree {
		// Get predecessor from left child
		predecessor := t.getPredecessor(n, index)
		n.keys[index] = predecessor
		t.delete(n.children[index], predecessor)
	} else if len(n.children[index+1].keys) >= t.degree {
		// Get successor from right child
		successor := t.getSuccessor(n, index)
		n.keys[index] = successor
		t.delete(n.children[index+1], successor)
	} else {
		// Merge with sibling
		t.merge(n, index)
		t.delete(n.children[index], key)
	}
}

// getPredecessor gets the predecessor key (rightmost in left subtree)
func (t *BTree[T]) getPredecessor(n *node[T], index int) T {
	curr := n.children[index]
	for !curr.leaf {
		curr = curr.children[len(curr.children)-1]
	}
	return curr.keys[len(curr.keys)-1]
}

// getSuccessor gets the successor key (leftmost in right subtree)
func (t *BTree[T]) getSuccessor(n *node[T], index int) T {
	curr := n.children[index+1]
	for !curr.leaf {
		curr = curr.children[0]
	}
	return curr.keys[0]
}

// fill ensures a child has at least t keys
func (t *BTree[T]) fill(n *node[T], index int) {
	// If previous sibling has at least t keys, borrow from it
	if index != 0 && len(n.children[index-1].keys) >= t.degree {
		t.borrowFromPrev(n, index)
	} else if index != len(n.children)-1 && len(n.children[index+1].keys) >= t.degree {
		// If next sibling has at least t keys, borrow from it
		t.borrowFromNext(n, index)
	} else {
		// Merge with sibling
		if index != len(n.children)-1 {
			t.merge(n, index)
		} else {
			t.merge(n, index-1)
		}
	}
}

// borrowFromPrev borrows a key from the previous sibling
func (t *BTree[T]) borrowFromPrev(n *node[T], childIndex int) {
	child := n.children[childIndex]
	sibling := n.children[childIndex-1]

	// Move a key from parent to child
	child.keys = append([]T{n.keys[childIndex-1]}, child.keys...)

	// Move a key from sibling to parent
	n.keys[childIndex-1] = sibling.keys[len(sibling.keys)-1]
	sibling.keys = sibling.keys[:len(sibling.keys)-1]

	// Move child pointer if not leaf
	if !child.leaf {
		child.children = append([]*node[T]{sibling.children[len(sibling.children)-1]}, child.children...)
		sibling.children = sibling.children[:len(sibling.children)-1]
	}
}

// borrowFromNext borrows a key from the next sibling
func (t *BTree[T]) borrowFromNext(n *node[T], childIndex int) {
	child := n.children[childIndex]
	sibling := n.children[childIndex+1]

	// Move a key from parent to child
	child.keys = append(child.keys, n.keys[childIndex])

	// Move a key from sibling to parent
	n.keys[childIndex] = sibling.keys[0]
	sibling.keys = sibling.keys[1:]

	// Move child pointer if not leaf
	if !child.leaf {
		child.children = append(child.children, sibling.children[0])
		sibling.children = sibling.children[1:]
	}
}

// merge merges a child with its sibling
func (t *BTree[T]) merge(n *node[T], index int) {
	child := n.children[index]
	sibling := n.children[index+1]

	// Pull key from this node and merge with right sibling
	child.keys = append(child.keys, n.keys[index])
	child.keys = append(child.keys, sibling.keys...)

	// Copy child pointers
	if !child.leaf {
		child.children = append(child.children, sibling.children...)
	}

	// Remove key from this node
	copy(n.keys[index:], n.keys[index+1:])
	n.keys = n.keys[:len(n.keys)-1]

	// Remove child pointer
	copy(n.children[index+1:], n.children[index+2:])
	n.children = n.children[:len(n.children)-1]
}

// InOrderTraversal traverses the tree in order and applies a function to each value
func (t *BTree[T]) InOrderTraversal(fn func(T)) {
	t.inOrderTraversal(t.root, fn)
}

// inOrderTraversal recursively traverses the tree in order
func (t *BTree[T]) inOrderTraversal(n *node[T], fn func(T)) {
	if n == nil {
		return
	}

	for i := 0; i < len(n.keys); i++ {
		if !n.leaf {
			t.inOrderTraversal(n.children[i], fn)
		}
		fn(n.keys[i])
	}

	if !n.leaf {
		t.inOrderTraversal(n.children[len(n.keys)], fn)
	}
}

// Min returns the minimum value in the tree
func (t *BTree[T]) Min() (T, bool) {
	var zero T
	if t.size == 0 {
		return zero, false
	}

	n := t.root
	for !n.leaf {
		n = n.children[0]
	}
	return n.keys[0], true
}

// Max returns the maximum value in the tree
func (t *BTree[T]) Max() (T, bool) {
	var zero T
	if t.size == 0 {
		return zero, false
	}

	n := t.root
	for !n.leaf {
		n = n.children[len(n.children)-1]
	}
	return n.keys[len(n.keys)-1], true
}

// Len returns the number of elements in the tree
func (t *BTree[T]) Len() int {
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *BTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Clear removes all elements from the tree
func (t *BTree[T]) Clear() {
	t.root = &node[T]{leaf: true}
	t.size = 0
}

// Height returns the height of the tree
func (t *BTree[T]) Height() int {
	return t.height(t.root)
}

// height recursively calculates the height of a node
func (t *BTree[T]) height(n *node[T]) int {
	if n == nil || n.leaf {
		return 0
	}
	return 1 + t.height(n.children[0])
}

// Degree returns the minimum degree of the tree
func (t *BTree[T]) Degree() int {
	return t.degree
}
