package queue

import (
	"testing"
)

func TestIterator_EmptyQueue(t *testing.T) {
	q := New[int](5)
	it := NewIterator(q)

	t.Run(
		"HasNext should return false for empty queue", func(t *testing.T) {
			if it.HasNext() {
				t.Error("HasNext() should return false for empty queue")
			}
		},
	)

	t.Run(
		"Next should return zero value and false", func(t *testing.T) {
			value, ok := it.Next()
			if ok {
				t.Error("Next() should return false for empty queue")
			}
			if value != 0 {
				t.Errorf("Next() should return zero value for empty queue, got %v", value)
			}
		},
	)
}

func TestIterator_SingleElement(t *testing.T) {
	q := New[string](5)
	q.Enqueue("test")
	it := NewIterator(q)

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
	q := New[int](5)
	expected := []int{1, 2, 3, 4, 5}

	for _, v := range expected {
		q.Enqueue(v)
	}

	t.Run(
		"Should iterate over all elements in order", func(t *testing.T) {
			it := NewIterator(q)
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
	q := New[int](5)
	values := []int{1, 2, 3}
	for _, v := range values {
		q.Enqueue(v)
	}

	t.Run(
		"Should reset to beginning of queue", func(t *testing.T) {
			it := NewIterator(q)

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
			it := NewIterator(q)

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

func TestQueueIterator_ModificationDuringIteration(t *testing.T) {
	q := New[int](5)
	q.Enqueue(1)
	q.Enqueue(2)

	it := NewIterator(q)

	t.Run(
		"Should maintain original snapshot during iteration and reset", func(t *testing.T) {
			// Start iteration
			first, _ := it.Next()

			// Modify queue during iteration
			q.Enqueue(3)
			q.Dequeue() // Removes 1

			// Continue iteration - should see original snapshot
			second, ok := it.Next()
			if !ok {
				t.Error("Next() should return true for second element")
			}
			if first != 1 || second != 2 {
				t.Errorf("Iterator values changed after queue modification, got %d,%d, want 1,2", first, second)
			}

			// Reset and verify we still see original snapshot
			it.Reset()
			first, _ = it.Next()
			second, _ = it.Next()
			if first != 1 || second != 2 {
				t.Errorf("Iterator values changed after reset, got %d,%d, want 1,2", first, second)
			}
		},
	)

	t.Run(
		"New iterator should get its own snapshot", func(t *testing.T) {
			anotherIt := NewIterator(q)

			var values []int
			for anotherIt.HasNext() {
				v, _ := anotherIt.Next()
				values = append(values, v)
			}

			if len(values) != 2 || values[0] != 2 || values[1] != 3 {
				t.Errorf("New iterator got wrong values, got %v, want [2,3]", values)
			}
		},
	)
}

func TestIterator_CustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := New[Person](5)
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}

	for _, p := range people {
		q.Enqueue(p)
	}

	t.Run(
		"Should work with custom types", func(t *testing.T) {
			it := NewIterator(q)
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
	q := New[int](5)
	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
	}

	t.Run(
		"Multiple iterators should work independently", func(t *testing.T) {
			it1 := NewIterator(q)
			it2 := NewIterator(q)

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

func TestIterator_CapacityExceeded(t *testing.T) {
	q := New[int](5)
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		q.Enqueue(v)
	}

	t.Run(
		"Should handle queue resizing", func(t *testing.T) {
			it := NewIterator(q)
			var actual []int

			for it.HasNext() {
				value, ok := it.Next()
				if !ok {
					t.Error("Next() returned false during iteration")
				}
				actual = append(actual, value)
			}

			if len(actual) != len(values) {
				t.Errorf("Wrong number of elements after queue resize, got %d, want %d", len(actual), len(values))
			}

			for i := range values {
				if actual[i] != values[i] {
					t.Errorf("Wrong value at position %d after queue resize, got %d, want %d", i, actual[i], values[i])
				}
			}
		},
	)
}
