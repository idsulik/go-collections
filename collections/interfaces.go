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
