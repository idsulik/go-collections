package priorityqueue

import (
	"encoding/json"

	"github.com/idsulik/go-collections/v2/internal/cmp"
)

type PriorityQueue[T any] struct {
	items  []T
	less   func(a, b T) bool
	equals func(a, b T) bool
}

// Option is a function that configures a PriorityQueue.
type Option[T any] func(*PriorityQueue[T])

// WithLess sets a custom less function for the PriorityQueue.
func WithLess[T any](less func(a, b T) bool) Option[T] {
	return func(pq *PriorityQueue[T]) {
		pq.less = less
	}
}

// WithEquals sets a custom equals function for the PriorityQueue.
func WithEquals[T any](equals func(a, b T) bool) Option[T] {
	return func(pq *PriorityQueue[T]) {
		pq.equals = equals
	}
}

// New creates a new PriorityQueue with the provided comparison function.
func New[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: []T{},
		less:  less,
		equals: func(a, b T) bool {
			// Since we can't use == with generic types, we marshal both items
			// to JSON and compare the results
			jsonA, _ := json.Marshal(a)
			jsonB, _ := json.Marshal(b)
			return string(jsonA) == string(jsonB)
		},
	}
}

// NewOrdered creates a new PriorityQueue with Ordered elements.
func NewOrdered[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items:  []T{},
		less:   func(a, b T) bool { return a < b },
		equals: func(a, b T) bool { return a == b },
	}
}

func ApplyOptions[T any](pq *PriorityQueue[T], opts ...Option[T]) {
	for _, opt := range opts {
		opt(pq)
	}
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue[T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.up(len(pq.items) - 1)
}

// Pop removes and returns the highest priority item from the queue.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if len(pq.items) == 0 {
		var zero T
		return zero, false
	}
	top := pq.items[0]
	last := len(pq.items) - 1
	pq.items[0] = pq.items[last]
	pq.items = pq.items[:last]
	pq.down(0)
	return top, true
}

// Peek returns the highest priority item without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if len(pq.items) == 0 {
		var zero T
		return zero, false
	}
	return pq.items[0], true
}

// Len returns the number of items in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

// IsEmpty checks if the priority queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

// Clear removes all items from the priority queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.items = []T{}
}

// MarshalJSON implements json.Marshaler interface
func (pq *PriorityQueue[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(pq.items)
}

// UnmarshalJSON implements json.Unmarshaler interface
func (pq *PriorityQueue[T]) UnmarshalJSON(data []byte) error {
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	pq.items = items
	// Heapify the entire queue
	for i := len(pq.items)/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
	return nil
}

// Contains checks if an item exists in the queue
// Note: This is an O(n) operation
func (pq *PriorityQueue[T]) Contains(item T) bool {
	for _, v := range pq.items {
		if pq.equals(item, v) {
			return true
		}
	}
	return false
}

// PushIfAbsent adds an item to the queue only if it's not already present
// Returns true if the item was added, false if it was already present
func (pq *PriorityQueue[T]) PushIfAbsent(item T) bool {
	if pq.Contains(item) {
		return false
	}

	pq.Push(item)
	return true
}

// Remove removes the first occurrence of the specified item from the queue
// Returns true if the item was found and removed, false otherwise
func (pq *PriorityQueue[T]) Remove(item T) bool {
	for i, v := range pq.items {
		if pq.equals(item, v) {
			// Remove the item by swapping with the last element and removing the last
			last := len(pq.items) - 1
			pq.items[i] = pq.items[last]
			pq.items = pq.items[:last]

			// Restore heap property
			if i < last {
				pq.down(i)
				pq.up(i)
			}
			return true
		}
	}
	return false
}

// Keys returns a slice of all items in the queue, maintaining heap order
func (pq *PriorityQueue[T]) Keys() []T {
	result := make([]T, len(pq.items))
	copy(result, pq.items)
	return result
}

// Vals is an alias for Keys() for compatibility
func (pq *PriorityQueue[T]) Vals() []T {
	return pq.Keys()
}

// up restores the heap property by moving the item at index i up.
func (pq *PriorityQueue[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || !pq.less(pq.items[i], pq.items[parent]) {
			break
		}
		pq.items[i], pq.items[parent] = pq.items[parent], pq.items[i]
		i = parent
	}
}

// down restores the heap property by moving the item at index i down.
func (pq *PriorityQueue[T]) down(i int) {
	n := len(pq.items)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && pq.less(pq.items[left], pq.items[smallest]) {
			smallest = left
		}
		if right < n && pq.less(pq.items[right], pq.items[smallest]) {
			smallest = right
		}
		if smallest == i {
			break
		}
		pq.items[i], pq.items[smallest] = pq.items[smallest], pq.items[i]
		i = smallest
	}
}
