package graph

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	g := New[int](false)
	if !g.AddNode(1) {
		t.Error("Expected to add node 1")
	}
	if !g.AddNode(2) {
		t.Error("Expected to add node 2")
	}
	if g.AddNode(1) {
		t.Error("Node 1 should not be added again")
	}

	if !g.HasNode(1) || !g.HasNode(2) {
		t.Error("Graph should contain nodes 1 and 2")
	}

	if len(g.Nodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(g.Nodes()))
	}
}

func TestAddEdge(t *testing.T) {
	g := New[int](false)
	if !g.AddEdge(1, 2, 1.0) {
		t.Error("Expected to add edge between 1 and 2")
	}
	if g.AddEdge(1, 2, 1.0) {
		t.Error("Edge between 1 and 2 should not be added again")
	}

	if !g.HasNode(1) || !g.HasNode(2) {
		t.Error("Graph should contain nodes 1 and 2 after adding edge")
	}

	if !g.HasEdge(1, 2) || !g.HasEdge(2, 1) {
		t.Error("Graph should have edge between 1 and 2")
	}

	weight, exists := g.GetEdgeWeight(1, 2)
	if !exists || weight != 1.0 {
		t.Error("Edge weight between 1 and 2 should be 1.0")
	}
}

func TestRemoveNode(t *testing.T) {
	g := New[int](false)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)
	if !g.RemoveNode(2) {
		t.Error("Expected to remove node 2")
	}
	if g.RemoveNode(2) {
		t.Error("Node 2 should not be removed again")
	}

	if g.HasNode(2) {
		t.Error("Node 2 should have been removed")
	}

	if g.HasEdge(1, 2) || g.HasEdge(2, 3) {
		t.Error("Edges connected to node 2 should have been removed")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := New[int](false)
	g.AddEdge(1, 2, 1.0)
	if !g.RemoveEdge(1, 2) {
		t.Error("Expected to remove edge between 1 and 2")
	}
	if g.RemoveEdge(1, 2) {
		t.Error("Edge between 1 and 2 should not be removed again")
	}

	if g.HasEdge(1, 2) || g.HasEdge(2, 1) {
		t.Error("Edge between 1 and 2 should have been removed")
	}
}

func TestNeighbors(t *testing.T) {
	g := New[int](false)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 1.0)

	neighbors := g.Neighbors(1)
	expected := map[int]bool{2: true, 3: true}

	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors, got %d", len(neighbors))
	}

	for _, n := range neighbors {
		if !expected[n] {
			t.Errorf("Unexpected neighbor: %d", n)
		}
	}
}

func TestDirected(t *testing.T) {
	g := New[int](true)
	if !g.AddEdge(1, 2, 1.0) {
		t.Error("Expected to add directed edge from 1 to 2")
	}

	if !g.HasEdge(1, 2) {
		t.Error("Directed graph should have edge from 1 to 2")
	}

	if g.HasEdge(2, 1) {
		t.Error("Directed graph should not have edge from 2 to 1")
	}

	neighbors := g.Neighbors(1)
	if len(neighbors) != 1 || neighbors[0] != 2 {
		t.Error("Neighbors of 1 should be [2]")
	}
}

func TestTraverse(t *testing.T) {
	g := New[int](false)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)
	g.AddEdge(3, 4, 1.0)
	g.AddEdge(4, 5, 1.0)

	visited := make(map[int]bool)
	g.Traverse(
		1, func(value int) {
			visited[value] = true
		},
	)

	if len(visited) != 5 {
		t.Errorf("Expected to visit 5 nodes, visited %d", len(visited))
	}

	for i := 1; i <= 5; i++ {
		if !visited[i] {
			t.Errorf("Node %d was not visited", i)
		}
	}
}

func TestEdges(t *testing.T) {
	g := New[int](false)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)
	g.AddEdge(3, 1, 3.0)

	edges := g.Edges()
	expectedEdges := map[[2]int]bool{
		{1, 2}: true,
		{2, 3}: true,
		{3, 1}: true,
	}

	if len(edges) != len(expectedEdges) {
		t.Errorf("Expected %d edges, got %d", len(expectedEdges), len(edges))
	}

	for _, edge := range edges {
		// Check both possible directions for undirected edges
		if !expectedEdges[[2]int{edge[0], edge[1]}] && !expectedEdges[[2]int{edge[1], edge[0]}] {
			t.Errorf("Unexpected edge: %v", edge)
		}
	}
}
