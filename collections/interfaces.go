package collections

import "github.com/idsulik/go-collections/v3/iterator"

// Collection is a base interface for all collections.
type Collection[T any] interface {
	Len() int
	IsEmpty() bool
	Clear()
}

// Set represents a unique collection of elements.
type Set[T comparable] interface {
	Collection[T]
	Add(item T)
	Remove(item T)
	Has(item T) bool
	Iterator() iterator.Iterator[T]
}

type Stack[T any] interface {
	Collection[T]
	Push(item T)
	Pop() (T, bool)
	Peek() (T, bool)
}

type Deque[T any] interface {
	Collection[T]
	PushFront(item T)
	PushBack(item T)
	PopFront() (T, bool)
	PopBack() (T, bool)
}

type Queue[T any] interface {
	Collection[T]
	Enqueue(item T)
	Dequeue() (T, bool)
	Peek() (T, bool)
}

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Put(key K, value V)
	Remove(key K) bool
	Contains(key K) bool
	Len() int
	Cap() int
	IsEmpty() bool
	Clear()
}
