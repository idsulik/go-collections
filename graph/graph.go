package graph

import (
	"github.com/idsulik/go-collections/v3/iterator"
)

// Graph represents the graph data structure.
type Graph[T comparable] struct {
	directed bool
	nodes    map[T]*node[T]
}

// node represents a node in the graph.
type node[T comparable] struct {
	value    T
	edges    map[T]*edge[T]
	incoming map[T]*edge[T] // For directed graphs
}

// edge represents an edge between two nodes.
type edge[T comparable] struct {
	from   *node[T]
	to     *node[T]
	weight float64
}

// New creates a new Graph. If directed is true, the graph is directed.
func New[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		directed: directed,
		nodes:    make(map[T]*node[T]),
	}
}

// AddNode adds a node to the graph and returns true if added, false if it already exists.
func (g *Graph[T]) AddNode(value T) bool {
	if _, exists := g.nodes[value]; exists {
		return false
	}

	g.nodes[value] = &node[T]{
		value:    value,
		edges:    make(map[T]*edge[T]),
		incoming: make(map[T]*edge[T]),
	}
	return true
}

// AddEdge adds an edge between two nodes with an optional weight and returns true if added, false if it already exists.
func (g *Graph[T]) AddEdge(from, to T, weight float64) bool {
	fromNode, fromExists := g.nodes[from]
	toNode, toExists := g.nodes[to]

	if !fromExists {
		fromNode = &node[T]{
			value:    from,
			edges:    make(map[T]*edge[T]),
			incoming: make(map[T]*edge[T]),
		}
		g.nodes[from] = fromNode
	}

	if !toExists {
		toNode = &node[T]{
			value:    to,
			edges:    make(map[T]*edge[T]),
			incoming: make(map[T]*edge[T]),
		}
		g.nodes[to] = toNode
	}

	if _, exists := fromNode.edges[to]; exists {
		return false // Edge already exists
	}

	newEdge := &edge[T]{
		from:   fromNode,
		to:     toNode,
		weight: weight,
	}

	fromNode.edges[to] = newEdge

	if !g.directed {
		// For undirected graphs, add the edge in both directions
		toNode.edges[from] = &edge[T]{
			from:   toNode,
			to:     fromNode,
			weight: weight,
		}
	} else {
		// For directed graphs, update incoming edges
		toNode.incoming[from] = newEdge
	}

	return true
}

// RemoveNode removes a node and all connected edges, returns true if removed, false if not found.
func (g *Graph[T]) RemoveNode(value T) bool {
	nodeToRemove, exists := g.nodes[value]
	if !exists {
		return false
	}

	// Remove all edges from other nodes to this node
	for _, n := range g.nodes {
		delete(n.edges, value)
		delete(n.incoming, value)
	}

	// For directed graphs, remove incoming edges
	if g.directed {
		for from := range nodeToRemove.incoming {
			delete(g.nodes[from].edges, value)
		}
	}

	delete(g.nodes, value)
	return true
}

// RemoveEdge removes an edge between two nodes and returns true if removed, false if not found.
func (g *Graph[T]) RemoveEdge(from, to T) bool {
	fromNode, fromExists := g.nodes[from]
	toNode, toExists := g.nodes[to]

	if !fromExists || !toExists {
		return false
	}

	if _, exists := fromNode.edges[to]; !exists {
		return false // Edge does not exist
	}

	delete(fromNode.edges, to)

	if !g.directed {
		delete(toNode.edges, from)
	} else {
		delete(toNode.incoming, from)
	}

	return true
}

// Neighbors returns a slice of nodes adjacent to the given node.
func (g *Graph[T]) Neighbors(value T) []T {
	node, exists := g.nodes[value]
	if !exists {
		return nil
	}

	neighbors := make([]T, 0, len(node.edges))
	for neighbor := range node.edges {
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}

// HasNode checks if a node exists in the graph.
func (g *Graph[T]) HasNode(value T) bool {
	_, exists := g.nodes[value]
	return exists
}

// HasEdge checks if an edge exists between two nodes.
func (g *Graph[T]) HasEdge(from, to T) bool {
	fromNode, fromExists := g.nodes[from]
	if !fromExists {
		return false
	}

	_, exists := fromNode.edges[to]
	return exists
}

// GetEdgeWeight returns the weight of the edge between two nodes.
func (g *Graph[T]) GetEdgeWeight(from, to T) (float64, bool) {
	if fromNode, fromExists := g.nodes[from]; fromExists {
		if edge, exists := fromNode.edges[to]; exists {
			return edge.weight, true
		}
	}
	return 0, false
}

// Traverse performs a breadth-first traversal starting from the given node.
func (g *Graph[T]) Traverse(start T, visit func(T)) {
	startNode, exists := g.nodes[start]
	if !exists {
		return
	}

	visited := make(map[T]bool)
	queue := []T{startNode.value}

	for len(queue) > 0 {
		currentValue := queue[0]
		queue = queue[1:]

		if visited[currentValue] {
			continue
		}

		visited[currentValue] = true
		visit(currentValue)

		currentNode := g.nodes[currentValue]
		for neighbor := range currentNode.edges {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

// Nodes returns a slice of all node values in the graph.
func (g *Graph[T]) Nodes() []T {
	nodes := make([]T, 0, len(g.nodes))
	for value := range g.nodes {
		nodes = append(nodes, value)
	}
	return nodes
}

// Edges returns a slice of all edges in the graph.
func (g *Graph[T]) Edges() [][2]T {
	var edges [][2]T
	seen := make(map[[2]T]bool)

	for fromValue, fromNode := range g.nodes {
		for toValue := range fromNode.edges {
			edgePair := [2]T{fromValue, toValue}
			if g.directed || !seen[edgePair] {
				edges = append(edges, edgePair)
				if !g.directed {
					seen[[2]T{toValue, fromValue}] = true
				}
			}
		}
	}

	return edges
}

func (g *Graph[T]) Iterator() iterator.Iterator[T] {
	nodes := g.Nodes()
	if len(nodes) == 0 {
		return &Iterator[T]{
			visited: make(map[T]bool),
			queue:   make([]T, 0),
			graph:   g,
		}
	}

	return NewIterator(g, nodes[0])
}

// ForEach applies a function to each node in the graph.
func (g *Graph[T]) ForEach(fn func(T)) {
	for value := range g.nodes {
		fn(value)
	}
}
