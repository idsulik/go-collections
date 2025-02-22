package linkedlist

import (
	"github.com/idsulik/go-collections/v2/iterator"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Iterator() iterator.Iterator[T] {
	return NewIterator(l)
}

// ForEach applies a function to each element in the list.
func (l *LinkedList[T]) ForEach(fn func(T)) {
	current := l.head
	for current != nil {
		fn(current.Value)
		current = current.Next
	}
}

// AddFront adds a new node with the given value to the front of the list.
func (l *LinkedList[T]) AddFront(value T) {
	newNode := &Node[T]{Value: value, Next: l.head}
	if l.head == nil {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

// PeekFront returns the value of the node at the front of the list without removing it.
func (l *LinkedList[T]) PeekFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.Value, true
}

// AddBack adds a new node with the given value to the end of the list.
func (l *LinkedList[T]) AddBack(value T) {
	newNode := &Node[T]{Value: value}
	if l.tail != nil {
		l.tail.Next = newNode
	}
	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
	l.size++
}

// PeekBack returns the value of the node at the end of the list without removing it.
func (l *LinkedList[T]) PeekBack() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.Value, true
}

// RemoveFront removes the node from the front of the list and returns its value.
func (l *LinkedList[T]) RemoveFront() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	value := l.head.Value
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return value, true
}

// RemoveBack removes the node from the end of the list and returns its value.
func (l *LinkedList[T]) RemoveBack() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	if l.head == l.tail {
		value := l.head.Value
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}
	current := l.head
	for current.Next != l.tail {
		current = current.Next
	}
	value := l.tail.Value
	l.tail = current
	l.tail.Next = nil
	l.size--
	return value, true
}

// IsEmpty checks if the list is empty.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Size returns the number of elements in the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Iterate iterates over the linked list and applies a function to each node's value
// until the end of the list or the function returns false.
func (l *LinkedList[T]) Iterate(fn func(T) bool) {
	current := l.head
	for current != nil {
		if !fn(current.Value) {
			break
		}
		current = current.Next
	}
}
