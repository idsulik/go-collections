package set

import (
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
	s := New[int]()
	s.Add(1)
	if !s.Has(1) {
		t.Errorf("Expected 1 to be in the set")
	}
}

// TestRemove checks removing elements from the Set.
func TestRemove(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Remove(1)
	if s.Has(1) {
		t.Errorf("Expected 1 to be removed from the set")
	}
}

// TestHas checks if the Set has an element.
func TestHas(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Has(1) {
		t.Errorf("Expected 1 to be in the set")
	}
}

// TestClear checks clearing the Set.
func TestClear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected set to be empty, got size %d", s.Len())
	}
}

// TestLen checks the size of the Set.
func TestLen(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	if s.Len() != 2 {
		t.Errorf("Expected size 2, got %d", s.Len())
	}
}

// TestIsEmpty checks if IsEmpty properly identifies an empty Set.
func TestIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty initially")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Errorf("Expected set not to be empty after adding an element")
	}
	s.Clear()
	if !s.IsEmpty() {
		t.Errorf("Expected set to be empty after clearing")
	}
}

// TestElements checks if Elements returns a correct slice of the set's elements.
func TestElements(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	expected := []int{1, 2}
	actual := s.Elements()
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected elements %v, got %v", expected, actual)
	}
}

// TestAddAll checks adding multiple elements to the Set.
func TestAddAll(t *testing.T) {
	s := New[int]()
	s.AddAll(1, 2, 3)
	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Errorf("Expected elements 1, 2, 3 to be in the set")
	}
}

// TestRemoveAll checks removing multiple elements from the Set.
func TestRemoveAll(t *testing.T) {
	s := New[int]()
	s.AddAll(1, 2, 3)
	s.RemoveAll(1, 2)
	if s.Has(1) || s.Has(2) || !s.Has(3) {
		t.Errorf("Expected elements 1 and 2 to be removed from the set")
	}
}

// TestDiff checks the difference between two sets.
func TestDiff(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	diff := s1.Diff(s2)
	expected := []int{1}
	actual := diff.Elements()
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected difference %v, got %v", expected, actual)
	}
}

// TestIntersect checks the intersection between two sets.
func TestIntersect(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	intersect := s1.Intersect(s2)
	expected := []int{2, 3}
	actual := intersect.Elements()
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected intersection %v, got %v", expected, actual)
	}
}

// TestUnion checks the union between two sets.
func TestUnion(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	union := s1.Union(s2)
	expected := []int{1, 2, 3, 4}
	actual := union.Elements()
	sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected union %v, got %v", expected, actual)
	}
}

// TestIsSubset checks if a set is a subset of another set.
func TestIsSubset(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2)
	s2.AddAll(1, 2, 3)
	if !s1.IsSubset(s2) {
		t.Errorf("Expected s1 to be a subset of s2")
	}
}

// TestIsSuperset checks if a set is a superset of another set.
func TestIsSuperset(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2, 3)
	s2.AddAll(1, 2)
	if !s1.IsSuperset(s2) {
		t.Errorf("Expected s1 to be a superset of s2")
	}
}

// TestEqual checks if two sets are equal.
func TestEqual(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()
	s1.AddAll(1, 2)
	s2.AddAll(1, 2)

	if !s1.Equal(s2) {
		t.Errorf("Expected s1 to be equal to s2")
	}

	s2.Add(3)
	if s1.Equal(s2) {
		t.Errorf("Expected s1 not to be equal to s2")
	}
	s2.Remove(3)

	s1.Add(3)
	if s1.Equal(s2) {
		t.Errorf("Expected s1 not to be equal to s2")
	}
}
