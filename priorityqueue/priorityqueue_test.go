package priorityqueue

import (
	"encoding/json"
	"testing"
)

func TestPriorityQueueOptions(t *testing.T) {
	t.Run(
		"Custom Equals Function", func(t *testing.T) {
			type Person struct {
				ID   int
				Name string
				Age  int
			}

			pq := New(
				func(a, b Person) bool {
					return a.Age < b.Age
				},
			)

			// Add custom equals that only compares IDs
			ApplyOptions(
				pq, WithEquals(
					func(a, b Person) bool {
						return a.ID == b.ID
					},
				),
			)

			p1 := Person{ID: 1, Name: "Alice", Age: 30}
			p2 := Person{ID: 1, Name: "Alice Updated", Age: 31} // Same ID, different age
			p3 := Person{ID: 2, Name: "Bob", Age: 25}

			pq.Push(p1)

			// Should return false because ID already exists
			if pq.PushIfAbsent(p2) {
				t.Error("Should not allow push of person with same ID")
			}

			// Should allow push of person with different ID
			if !pq.PushIfAbsent(p3) {
				t.Error("Should allow push of person with different ID")
			}
		},
	)

	t.Run(
		"Ordered Type With Custom Equals", func(t *testing.T) {
			pq := NewOrdered[int]()

			// Override default equals
			ApplyOptions(
				pq, WithEquals(
					func(a, b int) bool {
						// Consider numbers equal if they have the same parity
						return (a % 2) == (b % 2)
					},
				),
			)

			pq.Push(1)
			pq.Push(3)

			// Should not add 5 as it's considered equal to 1 (both odd)
			if pq.PushIfAbsent(5) {
				t.Error("Should not add 5 as it's considered equal to existing odd number")
			}

			// Should add 2 as it's even
			if !pq.PushIfAbsent(2) {
				t.Error("Should add 2 as no even numbers exist")
			}
		},
	)

	t.Run(
		"Update Less Function", func(t *testing.T) {
			pq := NewOrdered[int]()

			// Change from min-heap to max-heap
			ApplyOptions(
				pq, WithLess(
					func(a, b int) bool {
						return a > b
					},
				),
			)

			nums := []int{1, 3, 2, 5, 4}
			for _, n := range nums {
				pq.Push(n)
			}

			// Should now pop in descending order
			expected := []int{5, 4, 3, 2, 1}
			for _, exp := range expected {
				if val, ok := pq.Pop(); !ok || val != exp {
					t.Errorf("Expected %d, got %d", exp, val)
				}
			}
		},
	)

	t.Run(
		"Multiple Options", func(t *testing.T) {
			pq := NewOrdered[int]()

			ApplyOptions(
				pq,
				WithLess(
					func(a, b int) bool {
						return a > b // max-heap
					},
				),
				WithEquals(
					func(a, b int) bool {
						return a/10 == b/10 // equal if same tens digit
					},
				),
			)

			pq.Push(11)

			// Should not add 15 as it's in the same tens group as 11
			if pq.PushIfAbsent(15) {
				t.Error("Should not add 15 as it's in same tens group as 11")
			}

			// Should add 21 as it's in different tens group
			if !pq.PushIfAbsent(21) {
				t.Error("Should add 21 as it's in different tens group")
			}

			// First pop should be 21 due to max-heap property
			if val, ok := pq.Pop(); !ok || val != 21 {
				t.Errorf("Expected 21, got %d", val)
			}
		},
	)
}

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

func TestLenFunc(t *testing.T) {
	pq := NewOrdered[int]()

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	count := pq.LenFunc(
		func(v int) bool {
			return v%2 == 0
		},
	)
	if count != 1 {
		t.Errorf("Expected 1 even number, got %d", count)
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

func TestPopCustomStruct(t *testing.T) {
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

func TestPopFunc(t *testing.T) {
	pq := NewOrdered[int]()

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	item, ok := pq.PopFunc(
		func(v int) bool {
			return v == 2
		},
	)

	if !ok {
		t.Error("Expected to find and remove 2")
	}
	if item != 2 {
		t.Errorf("Expected 2, got %d", item)
	}

	item, ok = pq.PopFunc(
		func(v int) bool {
			return v == 4
		},
	)
	if ok {
		t.Error("Did not expect to find 4")
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

func TestMarshalUnmarshal(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	data, err := pq.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Add some items
	items := []int{5, 3, 4, 1, 2}
	for _, item := range items {
		pq.Push(item)
	}

	// Marshal
	data, err = json.Marshal(pq)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Create new queue and unmarshal
	newPQ := New(less)
	err = json.Unmarshal(data, newPQ)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify contents
	for i := 1; i <= 5; i++ {
		item, ok := newPQ.Pop()
		if !ok || item != i {
			t.Errorf("Expected %d, got %d", i, item)
		}
	}
}

func TestMarshalNil(t *testing.T) {
	var pq *PriorityQueue[int]
	//goland:noinspection GoDfaNilDereference
	_, err := pq.MarshalJSON()
	if err == nil {
		t.Error("Expected unmarshal into nil PriorityQueue to fail")
	}
}

func TestUnmarshalNil(t *testing.T) {
	var pq *PriorityQueue[int]
	//goland:noinspection GoDfaNilDereference
	err := pq.UnmarshalJSON([]byte(`[1,2,3]`))
	if err == nil {
		t.Error("Expected unmarshal into nil PriorityQueue to fail")
	}
}

func TestGetFunc(t *testing.T) {
	pq := NewOrdered[int]()

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	item := pq.GetFunc(func(v int) bool { return v == 2 })
	if item != 2 {
		t.Errorf("Expected 2, got %d", item)
	}

	item = pq.GetFunc(func(v int) bool { return v == 4 })
	if item != 0 {
		t.Errorf("Expected 0, got %d", item)
	}
}

func TestContains(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	if !pq.Contains(2) {
		t.Error("Expected to find 2 in queue")
	}

	if pq.Contains(4) {
		t.Error("Did not expect to find 4 in queue")
	}
}

func TestContainsFunc(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	if !pq.ContainsFunc(func(v int) bool { return v == 2 }) {
		t.Error("Expected to find 2 in queue")
	}

	if pq.ContainsFunc(func(v int) bool { return v == 4 }) {
		t.Error("Did not expect to find 4 in queue")
	}
}

func TestPushIfAbsent(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	// First push should succeed
	if !pq.PushIfAbsent(1) {
		t.Error("First push should succeed")
	}

	// Second push of same value should fail
	if pq.PushIfAbsent(1) {
		t.Error("Second push should fail")
	}

	if pq.Len() != 1 {
		t.Errorf("Expected length 1, got %d", pq.Len())
	}
}

func TestRemoveFunc(t *testing.T) {
	pq := NewOrdered[int]()

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	// Remove existing item
	if !pq.RemoveFunc(func(v int) bool { return v == 2 }) {
		t.Error("Expected to remove 2")
	}

	// Try to remove non-existent item
	if pq.RemoveFunc(func(v int) bool { return v == 4 }) {
		t.Error("Should not be able to remove non-existent item")
	}

	// Verify remaining items
	expected := []int{1, 3}
	for _, exp := range expected {
		item, ok := pq.Pop()
		if !ok || item != exp {
			t.Errorf("Expected %d, got %d", exp, item)
		}
	}
}

func TestRemove(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	// Remove existing item
	if !pq.Remove(2) {
		t.Error("Expected to remove 2")
	}

	// Try to remove non-existent item
	if pq.Remove(4) {
		t.Error("Should not be able to remove non-existent item")
	}

	// Verify remaining items
	expected := []int{1, 3}
	for _, exp := range expected {
		item, ok := pq.Pop()
		if !ok || item != exp {
			t.Errorf("Expected %d, got %d", exp, item)
		}
	}
}

func TestKeysAndVals(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	items := []int{5, 3, 4, 1, 2}
	for _, item := range items {
		pq.Push(item)
	}

	keys := pq.Keys()
	if len(keys) != 5 {
		t.Errorf("Expected 5 keys, got %d", len(keys))
	}

	vals := pq.Vals()
	if len(vals) != 5 {
		t.Errorf("Expected 5 values, got %d", len(vals))
	}

	// Verify that modifying the returned slices doesn't affect the queue
	keys[0] = 100
	if pq.items[0] == 100 {
		t.Error("Modifying returned keys should not affect queue")
	}
}

func TestClone(t *testing.T) {
	less := func(a, b int) bool {
		return a < b
	}
	pq := New(less)

	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	clone := pq.Clone()

	// Verify that the clone has the same contents
	for i := 1; i <= 3; i++ {
		if !clone.Contains(i) {
			t.Errorf("Expected to find %d in clone", i)
		}
	}

	// Modify the original queue
	pq.Pop()

	// Verify that the clone is unaffected
	if clone.Len() != 3 {
		t.Errorf("Expected length 3, got %d", clone.Len())
	}
}
