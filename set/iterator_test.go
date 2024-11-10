package set

import (
	"testing"
)

func TestIterator_EmptySet(t *testing.T) {
	it := NewIterator([]int{})

	t.Run(
		"HasNext should return false for empty set", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false for empty set")
			}
		},
	)

	t.Run(
		"Next should return zero value and false", func(t *testing.T) {
			value, ok := it.Next()
			if ok {
				t.Error("Next() should return false for empty set")
			}
			if value != 0 {
				t.Errorf("Next() should return zero value for empty set, got %v", value)
			}
		},
	)
}

func TestIterator_SingleElement(t *testing.T) {
	it := NewIterator([]string{"test"})

	t.Run(
		"HasNext should return true initially", func(t *testing.T) {
			if !it.HasNext() {
				t.Error("HasNext() should return true when there is an element")
			}
		},
	)

	t.Run(
		"Next should return element and true", func(t *testing.T) {
			value, ok := it.Next()
			if !ok {
				t.Error("Next() should return true for existing element")
			}
			if value != "test" {
				t.Errorf("Next() returned wrong value, got %v, want 'test'", value)
			}
		},
	)

	t.Run(
		"HasNext should return false after iteration", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false after iterating over single element")
			}
		},
	)
}

func TestIterator_MultipleElements(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	it := NewIterator(items)

	t.Run(
		"Should iterate over all elements in order", func(t *testing.T) {
			var actual []int

			for it.HasNext() {
				value, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				actual = append(actual, value)
			}

			if len(actual) != len(items) {
				t.Errorf("Iterator returned wrong number of elements, got %d, want %d", len(actual), len(items))
			}

			for i := range items {
				if actual[i] != items[i] {
					t.Errorf("Wrong value at position %d, got %d, want %d", i, actual[i], items[i])
				}
			}
		},
	)
}

func TestIterator_Reset(t *testing.T) {
	items := []int{1, 2, 3}
	it := NewIterator(items)

	t.Run(
		"Should reset to beginning", func(t *testing.T) {
			// Consume some elements
			it.Next()
			it.Next()

			// Reset iterator
			it.Reset()

			// Verify we're back at the start
			value, ok := it.Next()
			if !ok {
				t.Error("Next() should return true after reset")
			}
			if value != items[0] {
				t.Errorf("After reset, got %d, want %d", value, items[0])
			}
		},
	)

	t.Run(
		"Should allow full iteration after reset", func(t *testing.T) {
			it.Reset()
			count := 0
			for it.HasNext() {
				value, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				if value != items[count] {
					t.Errorf("Wrong value at position %d, got %d, want %d", count, value, items[count])
				}
				count++
			}

			if count != len(items) {
				t.Errorf("Wrong number of elements after reset, got %d, want %d", count, len(items))
			}
		},
	)

	t.Run(
		"Should handle multiple resets", func(t *testing.T) {
			it.Reset()
			first1, _ := it.Next()
			it.Reset()
			first2, _ := it.Next()

			if first1 != first2 {
				t.Errorf("Different first values after multiple resets: got %d and %d", first1, first2)
			}
		},
	)
}

func TestIterator_CustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}

	t.Run(
		"Should work with custom types", func(t *testing.T) {
			it := NewIterator(people)
			index := 0

			for it.HasNext() {
				person, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				if person != people[index] {
					t.Errorf("Wrong person at index %d, got %v, want %v", index, person, people[index])
				}
				index++
			}
		},
	)
}

func TestIterator_NilSlice(t *testing.T) {
	var items []int
	it := NewIterator(items)

	t.Run(
		"Should handle nil slice", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false for nil slice")
			}

			value, ok := it.Next()
			if ok {
				t.Error("Next() should return false for nil slice")
			}
			if value != 0 {
				t.Errorf("Next() should return zero value for nil slice, got %v", value)
			}
		},
	)
}

func TestIterator_ConcurrentIteration(t *testing.T) {
	items := []int{1, 2, 3}

	t.Run(
		"Multiple iterators should work independently", func(t *testing.T) {
			it1 := NewIterator(items)
			it2 := NewIterator(items)

			// Advance first iterator
			it1.Next()

			// Second iterator should start from beginning
			value, ok := it2.Next()
			if !ok {
				t.Error("Next() should return true for first element of second iterator")
			}
			if value != 1 {
				t.Errorf("Second iterator got wrong value, got %d, want 1", value)
			}
		},
	)
}

func TestIterator_BoundaryConditions(t *testing.T) {
	items := []int{42}
	it := NewIterator(items)

	t.Run(
		"Should handle boundary conditions", func(t *testing.T) {
			// Call Next() at the boundary
			value, ok := it.Next()
			if !ok || value != 42 {
				t.Errorf("First Next() failed, got %d, %v", value, ok)
			}

			// Should be at the end now
			if it.HasNext() {
				t.Error("HasNext() should be false after last element")
			}

			// Call Next() past the end
			value, ok = it.Next()
			if ok {
				t.Error("Next() should return false when past end")
			}
			if value != 0 {
				t.Errorf("Next() should return zero value when past end, got %d", value)
			}

			// Reset and verify we can iterate again
			it.Reset()
			if !it.HasNext() {
				t.Error("HasNext() should be true after reset")
			}
		},
	)
}
