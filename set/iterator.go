package set

import (
	"github.com/idsulik/go-collections/v3/iterator"
)

type Iterator[T any] struct {
	current int
	items   []T
}

func NewIterator[T any](items []T) iterator.Iterator[T] {
	return &Iterator[T]{items: items}
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
}
