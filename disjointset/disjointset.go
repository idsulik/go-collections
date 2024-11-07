package disjointset

// DisjointSet represents a disjoint set data structure
type DisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

// New creates a new DisjointSet instance
func New[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

// MakeSet creates a new set containing a single element
func (ds *DisjointSet[T]) MakeSet(x T) {
	if _, exists := ds.parent[x]; !exists {
		ds.parent[x] = x
		ds.rank[x] = 0
	}
}

// Find returns the representative element of the set containing x
// Uses path compression for optimization
func (ds *DisjointSet[T]) Find(x T) T {
	if _, exists := ds.parent[x]; !exists {
		return x
	}

	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// Union merges the sets containing elements x and y
// Uses union by rank for optimization
func (ds *DisjointSet[T]) Union(x, y T) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX == rootY {
		return
	}

	// Union by rank
	if ds.rank[rootX] < ds.rank[rootY] {
		ds.parent[rootX] = rootY
	} else if ds.rank[rootX] > ds.rank[rootY] {
		ds.parent[rootY] = rootX
	} else {
		ds.parent[rootY] = rootX
		ds.rank[rootX]++
	}
}

// Connected returns true if elements x and y are in the same set
func (ds *DisjointSet[T]) Connected(x, y T) bool {
	return ds.Find(x) == ds.Find(y)
}

// Clear removes all elements from the disjoint set
func (ds *DisjointSet[T]) Clear() {
	ds.parent = make(map[T]T)
	ds.rank = make(map[T]int)
}

// Len returns the number of elements in the disjoint set
func (ds *DisjointSet[T]) Len() int {
	return len(ds.parent)
}

// IsEmpty returns true if the disjoint set contains no elements
func (ds *DisjointSet[T]) IsEmpty() bool {
	return len(ds.parent) == 0
}

// GetSets returns a map of representatives to their set members
func (ds *DisjointSet[T]) GetSets() map[T][]T {
	sets := make(map[T][]T)
	for element := range ds.parent {
		root := ds.Find(element)
		sets[root] = append(sets[root], element)
	}
	return sets
}
