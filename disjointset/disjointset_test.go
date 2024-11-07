package disjointset

import (
	"testing"
)

func TestDisjointSet(t *testing.T) {
	t.Run(
		"New DisjointSet", func(t *testing.T) {
			ds := New[int]()
			if !ds.IsEmpty() {
				t.Error("New DisjointSet should be empty")
			}
		},
	)

	t.Run(
		"MakeSet", func(t *testing.T) {
			ds := New[int]()
			ds.MakeSet(1)
			if ds.Find(1) != 1 {
				t.Error("MakeSet should create a set with the element as its own representative")
			}
		},
	)

	t.Run(
		"Union and Find", func(t *testing.T) {
			ds := New[int]()
			ds.MakeSet(1)
			ds.MakeSet(2)
			ds.MakeSet(3)

			ds.Union(1, 2)
			if !ds.Connected(1, 2) {
				t.Error("Elements 1 and 2 should be connected after Union")
			}

			ds.Union(2, 3)
			if !ds.Connected(1, 3) {
				t.Error("Elements 1 and 3 should be connected after Union")
			}
		},
	)

	t.Run(
		"Connected", func(t *testing.T) {
			ds := New[string]()
			ds.MakeSet("A")
			ds.MakeSet("B")
			ds.MakeSet("C")

			if ds.Connected("A", "B") {
				t.Error("Elements should not be connected before Union")
			}

			ds.Union("A", "B")
			if !ds.Connected("A", "B") {
				t.Error("Elements should be connected after Union")
			}
		},
	)

	t.Run(
		"Clear", func(t *testing.T) {
			ds := New[int]()
			ds.MakeSet(1)
			ds.MakeSet(2)
			ds.Union(1, 2)

			ds.Clear()
			if !ds.IsEmpty() {
				t.Error("DisjointSet should be empty after Clear")
			}
		},
	)

	t.Run(
		"GetSets", func(t *testing.T) {
			ds := New[int]()
			ds.MakeSet(1)
			ds.MakeSet(2)
			ds.MakeSet(3)
			ds.MakeSet(4)

			ds.Union(1, 2)
			ds.Union(3, 4)

			sets := ds.GetSets()
			if len(sets) != 2 {
				t.Error("Should have exactly 2 distinct sets")
			}

			for _, set := range sets {
				if len(set) != 2 {
					t.Error("Each set should contain exactly 2 elements")
				}
			}
		},
	)

	t.Run(
		"Path Compression", func(t *testing.T) {
			ds := New[int]()
			ds.MakeSet(1)
			ds.MakeSet(2)
			ds.MakeSet(3)

			ds.Union(1, 2)
			ds.Union(2, 3)

			// After finding 3, the path should be compressed
			root := ds.Find(3)
			if ds.parent[3] != root {
				t.Error("Path compression should make 3 point directly to the root")
			}
		},
	)
}
