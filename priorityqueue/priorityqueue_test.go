package priorityqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	less := func(a, b int) bool {
		return a < b // Min-heap: smaller numbers have higher priority
	}
	pq := New(less)

	pq.Push(5)
	pq.Push(3)
	pq.Push(4)
	pq.Push(1)
	pq.Push(2)

	expectedOrder := []int{1, 2, 3, 4, 5}

	for i, expected := range expectedOrder {
		item, ok := pq.Pop()
		if !ok {
			t.Fatalf("Expected item %d but got none", expected)
		}
		if item != expected {
			t.Errorf("Test %d: Expected %d, got %d", i+1, expected, item)
		}
	}

	// The queue should now be empty
	if !pq.IsEmpty() {
		t.Error("Priority queue should be empty after popping all elements")
	}
}

func TestPeek(t *testing.T) {
	less := func(a, b int) bool {
		return a > b // Max-heap: larger numbers have higher priority
	}
	pq := New(less)

	pq.Push(10)
	pq.Push(30)
	pq.Push(20)

	item, ok := pq.Peek()
	if !ok {
		t.Fatal("Expected to peek an item but got none")
	}
	if item != 30 {
		t.Errorf("Expected top item to be 30, got %d", item)
	}

	// Ensure the item was not removed
	if pq.Len() != 3 {
		t.Errorf("Expected queue length to be 3, got %d", pq.Len())
	}
}

func TestIsEmptyAndLen(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	if !pq.IsEmpty() {
		t.Error("Newly created priority queue should be empty")
	}
	if pq.Len() != 0 {
		t.Errorf("Expected length 0, got %d", pq.Len())
	}

	pq.Push(1)
	if pq.IsEmpty() {
		t.Error("Priority queue should not be empty after pushing an element")
	}
	if pq.Len() != 1 {
		t.Errorf("Expected length 1, got %d", pq.Len())
	}
}

func TestClear(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	pq.Clear()

	if !pq.IsEmpty() {
		t.Error("Priority queue should be empty after Clear")
	}
	if pq.Len() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", pq.Len())
	}
}

type Task struct {
	name     string
	priority int
}

func TestCustomStruct(t *testing.T) {
	less := func(a, b Task) bool {
		return a.priority < b.priority
	}
	pq := New(less)

	pq.Push(Task{name: "Task 1", priority: 3})
	pq.Push(Task{name: "Task 2", priority: 1})
	pq.Push(Task{name: "Task 3", priority: 2})

	expectedOrder := []string{"Task 2", "Task 3", "Task 1"}

	for i, expectedName := range expectedOrder {
		task, ok := pq.Pop()
		if !ok {
			t.Fatalf("Expected task %s but got none", expectedName)
		}
		if task.name != expectedName {
			t.Errorf("Test %d: Expected %s, got %s", i+1, expectedName, task.name)
		}
	}
}

func TestPopEmpty(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	_, ok := pq.Pop()
	if ok {
		t.Error("Expected Pop to return false on empty queue")
	}
}

func TestPeekEmpty(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	_, ok := pq.Peek()
	if ok {
		t.Error("Expected Peek to return false on empty queue")
	}
}

func TestEqualPriority(t *testing.T) {
	type Item struct {
		value    string
		priority int
	}
	less := func(a, b Item) bool {
		return a.priority < b.priority
	}
	pq := New(less)

	pq.Push(Item{value: "Item 1", priority: 1})
	pq.Push(Item{value: "Item 2", priority: 1})
	pq.Push(Item{value: "Item 3", priority: 1})

	items := make(map[string]bool)
	for i := 0; i < 3; i++ {
		item, ok := pq.Pop()
		if !ok {
			t.Fatal("Expected an item but got none")
		}
		items[item.value] = true
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 unique items, got %d", len(items))
	}
}

func TestLargeDataSet(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	const numElements = 10000

	// Push elements in reverse order
	for i := numElements; i > 0; i-- {
		pq.Push(i)
	}

	// Pop elements and ensure they are in ascending order
	for i := 1; i <= numElements; i++ {
		item, ok := pq.Pop()
		if !ok {
			t.Fatalf("Expected item %d but got none", i)
		}
		if item != i {
			t.Errorf("Expected %d, got %d", i, item)
		}
	}
}
