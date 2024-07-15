package set

// Set represents a set of unique items.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New creates and returns a new, empty set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add adds an item to the set.
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove removes an item from the set.
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Has returns true if the set contains the specified item.
func (s *Set[T]) Has(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Clear removes all items from the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Len returns the number of items in the set.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// IsEmpty returns true if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Elements returns a slice containing all items in the set.
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.items))
	for item := range s.items {
		elements = append(elements, item)
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
	return out
}

// IsSubset returns true if the receiver set is a subset of the other set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
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
	return s.IsSubset(other) && s.IsSuperset(other)
}
