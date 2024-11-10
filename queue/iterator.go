package queue

// Iterator implements iterator.Iterator for Queue
type Iterator[T any] struct {
	current int
	items   []T
	queue   *Queue[T]
}

func NewIterator[T any](q *Queue[T]) *Iterator[T] {
	it := &Iterator[T]{queue: q}
	it.Reset()
	return it
}

func (it *Iterator[T]) HasNext() bool {
	return it.current < len(it.items)
}

func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}

	value := it.items[it.current]
	it.current++
	return value, true
}

func (it *Iterator[T]) Reset() {
	it.current = 0
	// Take a snapshot of current queue items
	it.items = make([]T, it.queue.Len())
	copy(it.items, it.queue.d.GetItems())
}
