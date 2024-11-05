package set

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

// TestNew checks the creation of a new Set.
func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Error("Expected new set to be non-nil")
	}
}

// TestAdd checks adding elements to the Set.
func TestAdd(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			if !s.Has(1) {
				t.Errorf("Expected 1 to be in the set")
			}
		},
	)

	t.Run(
		"float64 NaN handling", func(t *testing.T) {
			s := New[float64]()
			nan1 := math.NaN()
			nan2 := math.NaN()
			s.Add(nan1)
			if !s.Has(nan2) {
				t.Error("Expected NaN to be in the set")
			}
			if s.Len() != 1 {
				t.Errorf("Expected length 1, got %d", s.Len())
			}
		},
	)

	t.Run(
		"float32 NaN handling", func(t *testing.T) {
			s := New[float32]()
			nan1 := float32(math.NaN())
			nan2 := float32(math.NaN())
			s.Add(nan1)
			if !s.Has(nan2) {
				t.Error("Expected NaN to be in the set")
			}
			if s.Len() != 1 {
				t.Errorf("Expected length 1, got %d", s.Len())
			}
		},
	)
}

// TestRemove checks removing elements from the Set.
func TestRemove(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			s.Remove(1)
			if s.Has(1) {
				t.Errorf("Expected 1 to be removed from the set")
			}
		},
	)

	t.Run(
		"float64 NaN handling", func(t *testing.T) {
			s := New[float64]()
			nan1 := math.NaN()
			nan2 := math.NaN()
			s.Add(nan1)
			s.Remove(nan2)
			if s.Has(nan1) {
				t.Error("Expected NaN to be removed from the set")
			}
		},
	)
}

// TestHas checks if the Set has an element.
func TestHas(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			if !s.Has(1) {
				t.Errorf("Expected 1 to be in the set")
			}
		},
	)

	t.Run(
		"float64 NaN handling", func(t *testing.T) {
			s := New[float64]()
			nan1 := math.NaN()
			nan2 := math.NaN()
			s.Add(nan1)
			if !s.Has(nan2) {
				t.Error("Expected NaN to be in the set")
			}
		},
	)
}

// TestClear checks clearing the Set.
func TestClear(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			s.Add(2)
			s.Clear()
			if s.Len() != 0 {
				t.Errorf("Expected set to be empty, got size %d", s.Len())
			}
		},
	)

	t.Run(
		"with NaN values", func(t *testing.T) {
			s := New[float64]()
			s.Add(1.0)
			s.Add(math.NaN())
			s.Clear()
			if s.Len() != 0 {
				t.Errorf("Expected set to be empty, got size %d", s.Len())
			}
			if s.Has(math.NaN()) {
				t.Error("Expected NaN to be cleared from the set")
			}
		},
	)
}

// TestLen checks the size of the Set.
func TestLen(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			s.Add(2)
			if s.Len() != 2 {
				t.Errorf("Expected size 2, got %d", s.Len())
			}
		},
	)

	t.Run(
		"with NaN values", func(t *testing.T) {
			s := New[float64]()
			s.Add(1.0)
			s.Add(math.NaN())
			s.Add(math.NaN()) // Adding NaN twice shouldn't increase length
			if s.Len() != 2 {
				t.Errorf("Expected size 2, got %d", s.Len())
			}
		},
	)
}

// TestElements checks if Elements returns a correct slice of the set's elements.
func TestElements(t *testing.T) {
	t.Run(
		"regular integers", func(t *testing.T) {
			s := New[int]()
			s.Add(1)
			s.Add(2)
			expected := []int{1, 2}
			actual := s.Elements()
			sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Expected elements %v, got %v", expected, actual)
			}
		},
	)

	t.Run(
		"with NaN values", func(t *testing.T) {
			s := New[float64]()
			s.Add(1.0)
			s.Add(math.NaN())
			elements := s.Elements()
			if len(elements) != 2 {
				t.Errorf("Expected 2 elements, got %d", len(elements))
			}
			nanCount := 0
			regularCount := 0
			for _, e := range elements {
				if math.IsNaN(e) {
					nanCount++
				} else {
					regularCount++
				}
			}
			if nanCount != 1 || regularCount != 1 {
				t.Errorf("Expected 1 NaN and 1 regular value, got %d NaN and %d regular", nanCount, regularCount)
			}
		},
	)
}

// TestSet operations with NaN values
func TestSetOperationsWithNaN(t *testing.T) {
	t.Run(
		"Diff with NaN", func(t *testing.T) {
			s1 := New[float64]()
			s2 := New[float64]()
			s1.Add(1.0)
			s1.Add(math.NaN())
			s2.Add(1.0)

			diff := s1.Diff(s2)
			if diff.Len() != 1 || !diff.Has(math.NaN()) {
				t.Error("Expected diff to contain only NaN")
			}
		},
	)

	t.Run(
		"Intersect with NaN", func(t *testing.T) {
			s1 := New[float64]()
			s2 := New[float64]()
			s1.Add(math.NaN())
			s2.Add(math.NaN())

			intersect := s1.Intersect(s2)
			if intersect.Len() != 1 || !intersect.Has(math.NaN()) {
				t.Error("Expected intersection to contain NaN")
			}
		},
	)

	t.Run(
		"Union with NaN", func(t *testing.T) {
			s1 := New[float64]()
			s2 := New[float64]()
			s1.Add(1.0)
			s1.Add(math.NaN())
			s2.Add(2.0)

			union := s1.Union(s2)
			if union.Len() != 3 || !union.Has(math.NaN()) {
				t.Errorf("Expected union to contain 3 elements including NaN, got %d elements", union.Len())
			}
		},
	)

	t.Run(
		"IsSubset with NaN", func(t *testing.T) {
			s1 := New[float64]()
			s2 := New[float64]()
			s1.Add(math.NaN())
			s2.Add(math.NaN())
			s2.Add(1.0)

			if !s1.IsSubset(s2) {
				t.Error("Expected s1 to be subset of s2")
			}

			s1.Add(2.0)
			if s1.IsSubset(s2) {
				t.Error("Expected s1 not to be subset of s2")
			}
		},
	)

	t.Run(
		"Equal with NaN", func(t *testing.T) {
			s1 := New[float64]()
			s2 := New[float64]()
			s1.Add(1.0)
			s1.Add(math.NaN())
			s2.Add(1.0)
			s2.Add(math.NaN())

			if !s1.Equal(s2) {
				t.Error("Expected sets to be equal")
			}

			s2.Add(2.0)
			if s1.Equal(s2) {
				t.Error("Expected sets not to be equal")
			}
		},
	)
}

// TestInfinityHandling checks if the set properly handles infinity values
func TestInfinityHandling(t *testing.T) {
	s := New[float64]()
	posInf := math.Inf(1)
	negInf := math.Inf(-1)

	s.Add(posInf)
	s.Add(negInf)

	if !s.Has(posInf) {
		t.Error("Expected set to contain positive infinity")
	}

	if !s.Has(negInf) {
		t.Error("Expected set to contain negative infinity")
	}

	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}
}
