package iterator

// Iterator is a generic interface for iterating over collections
type Iterator[T any] interface {
	// HasNext returns true if there are more elements to iterate over
	HasNext() bool

	// Next returns the next element in the iteration.
	// Second return value is false if there are no more elements.
	Next() (T, bool)

	// Reset restarts the iteration from the beginning
	Reset()
}

// Iterable is an interface for collections that can provide an iterator
type Iterable[T any] interface {
	// Iterator returns a new iterator for the collection
	Iterator() Iterator[T]
}
