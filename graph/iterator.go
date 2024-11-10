package graph

// Iterator implements iterator.Iterator for Graph using breadth-first traversal
type Iterator[T comparable] struct {
	visited map[T]bool
	queue   []T
	graph   *Graph[T]
	start   T
}

// NewIterator creates a new iterator for breadth-first traversal starting from the given node
func NewIterator[T comparable](g *Graph[T], start T) *Iterator[T] {
	it := &Iterator[T]{
		visited: make(map[T]bool),
		queue:   make([]T, 0),
		graph:   g,
		start:   start,
	}

	// Only add start node to queue if it exists in the graph
	if g.HasNode(start) {
		it.queue = append(it.queue, start)
	}

	return it
}

// HasNext returns true if there are more nodes to visit
func (it *Iterator[T]) HasNext() bool {
	// Skip nodes that were removed from the graph
	for len(it.queue) > 0 && !it.graph.HasNode(it.queue[0]) {
		it.queue = it.queue[1:]
	}
	return len(it.queue) > 0
}

// Next returns the next node in the breadth-first traversal
func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}

	current := it.queue[0]
	it.queue = it.queue[1:]
	it.visited[current] = true

	// Add unvisited neighbors that exist in the graph
	for _, neighbor := range it.graph.Neighbors(current) {
		if !it.visited[neighbor] && !it.isQueued(neighbor) && it.graph.HasNode(neighbor) {
			it.queue = append(it.queue, neighbor)
		}
	}

	return current, true
}

// Reset restarts the iteration from the original start node
func (it *Iterator[T]) Reset() {
	it.visited = make(map[T]bool)
	it.queue = it.queue[:0]

	// Restart from original start node if it exists
	if it.graph.HasNode(it.start) {
		it.queue = append(it.queue, it.start)
	}
}

// isQueued checks if a node is already in the queue to prevent duplicates
func (it *Iterator[T]) isQueued(node T) bool {
	for _, n := range it.queue {
		if n == node {
			return true
		}
	}
	return false
}
