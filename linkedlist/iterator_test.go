package linkedlist

import (
	"testing"
)

func TestIterator_EmptyList(t *testing.T) {
	list := New[int]()
	it := NewIterator(list)

	t.Run(
		"HasNext should return false", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false for empty list")
			}
		},
	)

	t.Run(
		"Next should return zero value and false", func(t *testing.T) {
			value, ok := it.Next()
			if ok {
				t.Error("Next() should return false for empty list")
			}
			if value != 0 {
				t.Errorf("Next() should return zero value for empty list, got %v", value)
			}
		},
	)
}

func TestIterator_SingleElement(t *testing.T) {
	list := New[string]()
	list.AddBack("test")
	it := NewIterator(list)

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
	list := New[int]()
	expected := []int{1, 2, 3, 4, 5}

	for _, v := range expected {
		list.AddBack(v)
	}

	t.Run(
		"Should iterate over all elements in order", func(t *testing.T) {
			it := NewIterator(list)
			var actual []int

			for it.HasNext() {
				value, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				actual = append(actual, value)
			}

			if len(actual) != len(expected) {
				t.Errorf("Iterator returned wrong number of elements, got %d, want %d", len(actual), len(expected))
			}

			for i := range expected {
				if actual[i] != expected[i] {
					t.Errorf("Wrong value at position %d, got %d, want %d", i, actual[i], expected[i])
				}
			}
		},
	)
}

func TestIterator_Reset(t *testing.T) {
	list := New[int]()
	values := []int{1, 2, 3}
	for _, v := range values {
		list.AddBack(v)
	}

	t.Run(
		"Should reset to beginning of list", func(t *testing.T) {
			it := NewIterator(list)

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
			if value != values[0] {
				t.Errorf("After reset, got %d, want %d", value, values[0])
			}
		},
	)

	t.Run(
		"Should allow full iteration after reset", func(t *testing.T) {
			it := NewIterator(list)

			// Consume all elements
			for it.HasNext() {
				it.Next()
			}

			// Reset and count elements
			it.Reset()
			count := 0
			for it.HasNext() {
				_, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during second iteration")
				}
				count++
			}

			if count != len(values) {
				t.Errorf("Wrong number of elements after reset, got %d, want %d", count, len(values))
			}
		},
	)
}

func TestIterator_ModificationDuringIteration(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	it := NewIterator(list)

	t.Run(
		"Should reflect list state at creation", func(t *testing.T) {
			// Start iteration
			first, _ := it.Next()

			// Modify list during iteration
			list.AddBack(3)
			list.RemoveFront()

			// Continue iteration
			second, ok := it.Next()
			if !ok {
				t.Error("Next() should return true for second element")
			}
			if first != 1 || second != 2 {
				t.Errorf("Iterator values changed after list modification, got %d,%d, want 1,2", first, second)
			}
		},
	)
}

func TestIterator_CustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	list := New[Person]()
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}

	for _, p := range people {
		list.AddBack(p)
	}

	t.Run(
		"Should work with custom types", func(t *testing.T) {
			it := NewIterator(list)
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

func TestIterator_ConcurrentIteration(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 3; i++ {
		list.AddBack(i)
	}

	t.Run(
		"Multiple iterators should work independently", func(t *testing.T) {
			it1 := NewIterator(list)
			it2 := NewIterator(list)

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
