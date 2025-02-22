package linkedlist

import (
	"github.com/idsulik/go-collections/v2/iterator"
)

// Iterator implements iterator.Iterator for LinkedList
type Iterator[T any] struct {
	current *Node[T]
	list    *LinkedList[T]
}

func NewIterator[T any](list *LinkedList[T]) iterator.Iterator[T] {
	return &Iterator[T]{list: list, current: list.head}
}

func (it *Iterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zero T
		return zero, false
	}
	value := it.current.Value
	it.current = it.current.Next
	return value, true
}

func (it *Iterator[T]) Reset() {
	it.current = it.list.head
}
