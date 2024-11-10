package graph

// Iterator implements iterator.Iterator for Graph using breadth-first traversal
type Iterator[T comparable] struct {
	visited map[T]bool
	queue   []T
	graph   *Graph[T]
}

func NewIterator[T comparable](g *Graph[T], start T) *Iterator[T] {
	it := &Iterator[T]{
		visited: make(map[T]bool),
		queue:   make([]T, 0),
		graph:   g,
	}
	it.queue = append(it.queue, start)
	return it
}

func (it *Iterator[T]) HasNext() bool {
	return len(it.queue) > 0
}

func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}

	current := it.queue[0]
	it.queue = it.queue[1:]

	if !it.visited[current] {
		it.visited[current] = true
		// Add unvisited neighbors
		for _, neighbor := range it.graph.Neighbors(current) {
			if !it.visited[neighbor] {
				it.queue = append(it.queue, neighbor)
			}
		}
	}

	return current, true
}

func (it *Iterator[T]) Reset() {
	it.visited = make(map[T]bool)
	it.queue = it.queue[:0]
	// Start over with the first node
	if nodes := it.graph.Nodes(); len(nodes) > 0 {
		it.queue = append(it.queue, nodes[0])
	}
}
