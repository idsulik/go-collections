package priorityqueue

type PriorityQueue[T any] struct {
	items []T
	less  func(a, b T) bool
}

// New creates a new PriorityQueue with the provided comparison function.
func New[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: []T{},
		less:  less,
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
