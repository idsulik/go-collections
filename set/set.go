package set

import "math"

// Set represents a set of unique items.
type Set[T comparable] struct {
	items map[T]struct{}
	// Store NaN values separately since NaN != NaN but maps treat NaN keys as equal
	hasNaN bool
}

// New creates and returns a new, empty set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items:  make(map[T]struct{}),
		hasNaN: false,
	}
}

// isNaN checks if the value is a NaN float
func isNaN[T comparable](v T) bool {
	// Type assert to check if T is float32 or float64
	switch val := any(v).(type) {
	case float32:
		return math.IsNaN(float64(val))
	case float64:
		return math.IsNaN(val)
	default:
		return false
	}
}

// Add adds an item to the set.
func (s *Set[T]) Add(item T) {
	if isNaN(item) {
		s.hasNaN = true
		return
	}
	s.items[item] = struct{}{}
}

// Remove removes an item from the set.
func (s *Set[T]) Remove(item T) {
	if isNaN(item) {
		s.hasNaN = false
		return
	}
	delete(s.items, item)
}

// Has returns true if the set contains the specified item.
func (s *Set[T]) Has(item T) bool {
	if isNaN(item) {
		return s.hasNaN
	}
	_, ok := s.items[item]
	return ok
}

// Clear removes all items from the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
	s.hasNaN = false
}

// Len returns the number of items in the set.
func (s *Set[T]) Len() int {
	count := len(s.items)
	if s.hasNaN {
		count++
	}
	return count
}

// IsEmpty returns true if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0 && !s.hasNaN
}

// Elements returns a slice containing all items in the set.
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, s.Len())
	for item := range s.items {
		elements = append(elements, item)
	}
	// Add NaN if present
	if s.hasNaN {
		var nan T
		switch any(nan).(type) {
		case float32:
			elements = append(elements, any(float32(math.NaN())).(T))
		case float64:
			elements = append(elements, any(math.NaN()).(T))
		}
	}
	return elements
}

// AddAll adds multiple items to the set.
func (s *Set[T]) AddAll(items ...T) {
	for _, item := range items {
		s.Add(item)
	}
}

// RemoveAll removes multiple items from the set.
func (s *Set[T]) RemoveAll(items ...T) {
	for _, item := range items {
		s.Remove(item)
	}
}

// Diff returns a new set containing items that are in the receiver set but not in the other set.
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	out := New[T]()
	for item := range s.items {
		if !other.Has(item) {
			out.Add(item)
		}
	}
	s.handleNan(other, out)
	return out
}

// Intersect returns a new set containing items that are in both the receiver set and the other set.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	out := New[T]()
	for item := range s.items {
		if other.Has(item) {
			out.Add(item)
		}
	}
	s.handleNan(other, out)
	return out
}

// Union returns a new set containing items that are in either the receiver set or the other set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	out := New[T]()
	for item := range s.items {
		out.Add(item)
	}
	for item := range other.items {
		out.Add(item)
	}

	s.handleNan(other, out)
	return out
}

// IsSubset returns true if the receiver set is a subset of the other set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.hasNaN && !other.hasNaN {
		return false
	}
	for item := range s.items {
		if !other.Has(item) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if the receiver set is a superset of the other set.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equal returns true if the receiver set is equal to the other set.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	if s.hasNaN != other.hasNaN {
		return false
	}
	for item := range s.items {
		if !other.Has(item) {
			return false
		}
	}

	return true
}

func (s *Set[T]) handleNan(other *Set[T], out *Set[T]) {
	if s.hasNaN || other.hasNaN {
		var nan T
		switch any(nan).(type) {
		case float32:
			out.Add(any(float32(math.NaN())).(T))
		case float64:
			out.Add(any(math.NaN()).(T))
		}
	}
}

// Iterator returns a new iterator for the set.
func (s *Set[T]) Iterator() *Iterator[T] {
	return NewIterator(s.Elements())
}
