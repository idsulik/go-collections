package avltree

// Node represents a node in the AVL tree
type Node[T any] struct {
	Value       T
	Left, Right *Node[T]
	Height      int
}

// AVLTree represents an AVL tree data structure
type AVLTree[T any] struct {
	root    *Node[T]
	size    int
	compare func(a, b T) int
}

// New creates a new AVL tree
func New[T any](compare func(a, b T) int) *AVLTree[T] {
	return &AVLTree[T]{
		compare: compare,
	}
}

// getHeight returns the height of a node
func (t *AVLTree[T]) getHeight(node *Node[T]) int {
	if node == nil {
		return -1
	}
	return node.Height
}

// getBalance returns the balance factor of a node
func (t *AVLTree[T]) getBalance(node *Node[T]) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.Left) - t.getHeight(node.Right)
}

// updateHeight updates the height of a node
func (t *AVLTree[T]) updateHeight(node *Node[T]) {
	node.Height = max(t.getHeight(node.Left), t.getHeight(node.Right)) + 1
}

// rotateRight performs a right rotation
func (t *AVLTree[T]) rotateRight(y *Node[T]) *Node[T] {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	t.updateHeight(y)
	t.updateHeight(x)

	return x
}

// rotateLeft performs a left rotation
func (t *AVLTree[T]) rotateLeft(x *Node[T]) *Node[T] {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	t.updateHeight(x)
	t.updateHeight(y)

	return y
}

// Insert adds a new value to the AVL tree
func (t *AVLTree[T]) Insert(value T) {
	t.root = t.insert(t.root, value)
	t.size++
}

// insert recursively inserts a value and balances the tree
func (t *AVLTree[T]) insert(node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: value, Height: 0}
	}

	comp := t.compare(value, node.Value)
	if comp < 0 {
		node.Left = t.insert(node.Left, value)
	} else if comp > 0 {
		node.Right = t.insert(node.Right, value)
	} else {
		return node // Duplicate value, ignore
	}

	t.updateHeight(node)
	balance := t.getBalance(node)

	// Left Left Case
	if balance > 1 && t.compare(value, node.Left.Value) < 0 {
		return t.rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && t.compare(value, node.Right.Value) > 0 {
		return t.rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && t.compare(value, node.Left.Value) > 0 {
		node.Left = t.rotateLeft(node.Left)
		return t.rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && t.compare(value, node.Right.Value) < 0 {
		node.Right = t.rotateRight(node.Right)
		return t.rotateLeft(node)
	}

	return node
}

// Delete removes a value from the AVL tree
func (t *AVLTree[T]) Delete(value T) bool {
	if t.root == nil {
		return false
	}

	var deleted bool
	t.root, deleted = t.delete(t.root, value)
	if deleted {
		t.size--
	}
	return deleted
}

// delete recursively deletes a value and balances the tree
func (t *AVLTree[T]) delete(node *Node[T], value T) (*Node[T], bool) {
	if node == nil {
		return nil, false
	}

	comp := t.compare(value, node.Value)
	var deleted bool

	if comp < 0 {
		node.Left, deleted = t.delete(node.Left, value)
	} else if comp > 0 {
		node.Right, deleted = t.delete(node.Right, value)
	} else {
		// Node with only one child or no child
		if node.Left == nil {
			return node.Right, true
		} else if node.Right == nil {
			return node.Left, true
		}

		// Node with two children
		successor := t.findMin(node.Right)
		node.Value = successor.Value
		node.Right, deleted = t.delete(node.Right, successor.Value)
	}

	if !deleted {
		return node, false
	}

	t.updateHeight(node)
	balance := t.getBalance(node)

	// Left Left Case
	if balance > 1 && t.getBalance(node.Left) >= 0 {
		return t.rotateRight(node), true
	}

	// Left Right Case
	if balance > 1 && t.getBalance(node.Left) < 0 {
		node.Left = t.rotateLeft(node.Left)
		return t.rotateRight(node), true
	}

	// Right Right Case
	if balance < -1 && t.getBalance(node.Right) <= 0 {
		return t.rotateLeft(node), true
	}

	// Right Left Case
	if balance < -1 && t.getBalance(node.Right) > 0 {
		node.Right = t.rotateRight(node.Right)
		return t.rotateLeft(node), true
	}

	return node, true
}

// findMin returns the node with minimum value in the tree
func (t *AVLTree[T]) findMin(node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Search looks for a value in the tree
func (t *AVLTree[T]) Search(value T) bool {
	return t.search(t.root, value)
}

// search recursively searches for a value
func (t *AVLTree[T]) search(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	comp := t.compare(value, node.Value)
	if comp < 0 {
		return t.search(node.Left, value)
	} else if comp > 0 {
		return t.search(node.Right, value)
	}
	return true
}

// InOrderTraversal performs an in-order traversal of the tree
func (t *AVLTree[T]) InOrderTraversal(fn func(T)) {
	t.inOrder(t.root, fn)
}

// inOrder performs an in-order traversal from a given node
func (t *AVLTree[T]) inOrder(node *Node[T], fn func(T)) {
	if node != nil {
		t.inOrder(node.Left, fn)
		fn(node.Value)
		t.inOrder(node.Right, fn)
	}
}

// Clear removes all elements from the tree
func (t *AVLTree[T]) Clear() {
	t.root = nil
	t.size = 0
}

// Len returns the number of nodes in the tree
func (t *AVLTree[T]) Len() int {
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *AVLTree[T]) IsEmpty() bool {
	return t.size == 0
}

// Height returns the height of the tree
func (t *AVLTree[T]) Height() int {
	return t.getHeight(t.root) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
