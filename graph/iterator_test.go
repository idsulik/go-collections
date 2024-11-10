package graph

import (
	"sort"
	"testing"
)

func TestIterator_EmptyGraph(t *testing.T) {
	g := New[string](false)
	it := NewIterator(g, "")

	t.Run(
		"HasNext should return false for empty graph", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false for empty graph")
			}
		},
	)

	t.Run(
		"Next should return zero value and false", func(t *testing.T) {
			value, ok := it.Next()
			if ok {
				t.Error("Next() should return false for empty graph")
			}
			if value != "" {
				t.Errorf("Next() should return zero value for empty graph, got %v", value)
			}
		},
	)
}

func TestIterator_SingleNode(t *testing.T) {
	g := New[string](false)
	g.AddNode("A")
	it := NewIterator(g, "A")

	t.Run(
		"HasNext should return true initially", func(t *testing.T) {
			if !it.HasNext() {
				t.Error("HasNext() should return true when there is a node")
			}
		},
	)

	t.Run(
		"Next should return node and true", func(t *testing.T) {
			value, ok := it.Next()
			if !ok {
				t.Error("Next() should return true for existing node")
			}
			if value != "A" {
				t.Errorf("Next() returned wrong value, got %v, want 'A'", value)
			}
		},
	)

	t.Run(
		"HasNext should return false after visiting single node", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false after visiting single node")
			}
		},
	)
}

func TestIterator_LinearGraph(t *testing.T) {
	g := New[int](false)
	// Create a linear graph: 1 - 2 - 3 - 4
	edges := [][2]int{{1, 2}, {2, 3}, {3, 4}}
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], 1.0)
	}

	t.Run(
		"Should visit all nodes in BFS order", func(t *testing.T) {
			it := NewIterator(g, 1)
			var visited []int

			for it.HasNext() {
				value, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				visited = append(visited, value)
			}

			if len(visited) != 4 {
				t.Errorf("Expected 4 nodes to be visited, got %d", len(visited))
			}

			// Check BFS order: 1, 2, 3, 4
			expected := []int{1, 2, 3, 4}
			for i, v := range expected {
				if visited[i] != v {
					t.Errorf("Wrong BFS order at position %d: got %d, want %d", i, visited[i], v)
				}
			}
		},
	)
}

func TestIterator_CyclicGraph(t *testing.T) {
	g := New[string](false)
	// Create a cyclic graph: A - B - C - A
	edges := [][2]string{{"A", "B"}, {"B", "C"}, {"C", "A"}}
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], 1.0)
	}

	t.Run(
		"Should handle cycles without infinite loop", func(t *testing.T) {
			it := NewIterator(g, "A")
			visited := make(map[string]bool)

			for it.HasNext() {
				value, _ := it.Next()
				visited[value] = true
			}

			expected := []string{"A", "B", "C"}
			for _, v := range expected {
				if !visited[v] {
					t.Errorf("Node %s was not visited", v)
				}
			}

			if len(visited) != 3 {
				t.Errorf("Expected 3 unique nodes to be visited, got %d", len(visited))
			}
		},
	)
}

func TestIterator_DisconnectedGraph(t *testing.T) {
	g := New[string](false)
	// Create two disconnected components: (A-B) (C-D)
	g.AddEdge("A", "B", 1.0)
	g.AddEdge("C", "D", 1.0)

	t.Run(
		"Should only visit connected component from start node", func(t *testing.T) {
			it := NewIterator(g, "A")
			visited := make(map[string]bool)

			for it.HasNext() {
				value, _ := it.Next()
				visited[value] = true
			}

			// Should only visit A and B
			if !visited["A"] || !visited["B"] {
				t.Error("Failed to visit nodes in connected component")
			}
			if visited["C"] || visited["D"] {
				t.Error("Visited nodes in disconnected component")
			}
		},
	)
}

func TestIterator_Reset(t *testing.T) {
	g := New[int](false)
	edges := [][2]int{{1, 2}, {2, 3}, {3, 4}}
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], 1.0)
	}

	t.Run(
		"Should allow complete retraversal after reset", func(t *testing.T) {
			it := NewIterator(g, 1)

			// First traversal
			firstVisit := make([]int, 0)
			for it.HasNext() {
				value, _ := it.Next()
				firstVisit = append(firstVisit, value)
			}

			// Reset and second traversal
			it.Reset()
			secondVisit := make([]int, 0)
			for it.HasNext() {
				value, _ := it.Next()
				secondVisit = append(secondVisit, value)
			}

			// Compare both traversals
			if len(firstVisit) != len(secondVisit) {
				t.Errorf(
					"Different number of nodes visited after reset: first %d, second %d",
					len(firstVisit), len(secondVisit),
				)
			}

			// Sort both slices to compare (since BFS order might vary after reset)
			sort.Ints(firstVisit)
			sort.Ints(secondVisit)
			for i := range firstVisit {
				if firstVisit[i] != secondVisit[i] {
					t.Errorf(
						"Different nodes visited after reset at position %d: first %d, second %d",
						i, firstVisit[i], secondVisit[i],
					)
				}
			}
		},
	)
}

func TestIterator_DirectedGraph(t *testing.T) {
	g := New[int](true) // Create directed graph
	// Create a directed path: 1 -> 2 -> 3
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)

	t.Run(
		"Should respect edge direction", func(t *testing.T) {
			it := NewIterator(g, 1)
			var visited []int

			for it.HasNext() {
				value, _ := it.Next()
				visited = append(visited, value)
			}

			// Check order: should be able to reach all nodes from 1
			expected := []int{1, 2, 3}
			if len(visited) != len(expected) {
				t.Errorf("Expected %d nodes to be visited, got %d", len(expected), len(visited))
			}

			// Start from node 3 (should not reach 1 or 2)
			it = NewIterator(g, 3)
			visited = nil
			for it.HasNext() {
				value, _ := it.Next()
				visited = append(visited, value)
			}

			if len(visited) != 1 || visited[0] != 3 {
				t.Error("Should only visit starting node in directed graph with no outgoing edges")
			}
		},
	)
}

func TestIterator_ModificationDuringIteration(t *testing.T) {
	g := New[string](false)
	g.AddEdge("A", "B", 1.0)
	g.AddEdge("B", "C", 1.0)

	t.Run(
		"Should handle graph modifications during iteration", func(t *testing.T) {
			it := NewIterator(g, "A")

			// Get first node
			first, _ := it.Next()
			if first != "A" {
				t.Errorf("Expected first node to be 'A', got %s", first)
			}

			// Modify graph during iteration
			g.AddEdge("B", "D", 1.0)
			g.RemoveNode("C")

			// Continue iteration
			visited := map[string]bool{first: true}
			for it.HasNext() {
				value, _ := it.Next()
				visited[value] = true
			}

			// Should still visit B and D, but not C
			if !visited["B"] {
				t.Error("Failed to visit node B")
			}
			if !visited["D"] {
				t.Error("Failed to visit new node D")
			}
			if visited["C"] {
				t.Error("Visited removed node C")
			}
		},
	)
}
