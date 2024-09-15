# go-collections
[![Go Report Card](https://goreportcard.com/badge/github.com/idsulik/go-collections)](https://goreportcard.com/report/github.com/idsulik/go-collections)
![Build Status](https://img.shields.io/github/actions/workflow/status/idsulik/go-collections/go.yaml?branch=main)
[![Version](https://img.shields.io/github/v/release/idsulik/go-collections)](https://github.com/idsulik/go-collections/releases)
[![License](https://img.shields.io/github/license/idsulik/go-collections)](https://github.com/idsulik/go-collections/blob/main/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/idsulik/go-collections)](https://pkg.go.dev/github.com/idsulik/go-collections)

`go-collections` is a Go package that provides implementations of common data structures including a double-ended queue (Deque), a linked list, a queue, a trie, and a stack. This package offers a simple and efficient way to use these structures in Go, with support for generic types.

## Installation

You can install the package using the Go module system:

```sh
go get github.com/idsulik/go-collections
```

## Usage

Here is a brief example of how to use the `Deque`:

```go
package main

import (
	"fmt"
	"github.com/idsulik/go-collections/deque"
)

func main() {
	d := deque.New 
	d.PushBack(1)
	d.PushFront(2)

	front, _ := d.PopFront()
	back, _ := d.PopBack()

	fmt.Println(front) // Output: 2
	fmt.Println(back)  // Output: 1
}
```

## Data Structures

---

# Set

A set represents a collection of unique items.

## Type

```go
type Set[T comparable]
```

## Constructor

```go
func New[T comparable]() *Set[T]
```

## Methods

- `Add(item T)`: Adds an item to the set.
- `Remove(item T)`: Removes an item from the set.
- `Has(item T) bool`: Returns `true` if the set contains the specified item.
- `Clear()`: Removes all items from the set.
- `Len() int`: Returns the number of items in the set.
- `IsEmpty() bool`: Returns `true` if the set is empty.
- `Elements() []T`: Returns a slice containing all items in the set.
- `AddAll(items ...T)`: Adds multiple items to the set.
- `RemoveAll(items ...T)`: Removes multiple items from the set.
- `Diff(other *Set[T]) *Set[T]`: Returns a new set containing items that are in the receiver set but not in the other set.
- `Intersect(other *Set[T]) *Set[T]`: Returns a new set containing items that are in both the receiver set and the other set.
- `Union(other *Set[T]) *Set[T]`: Returns a new set containing items that are in either the receiver set or the other set.
- `IsSubset(other *Set[T]) bool`: Returns `true` if the receiver set is a subset of the other set.
- `IsSuperset(other *Set[T]) bool`: Returns `true` if the receiver set is a superset of the other set.
- `Equal(other *Set[T]) bool`: Returns `true` if the receiver set is equal to the other set.

---

### Deque

A double-ended queue (Deque) allows adding and removing elements from both the front and the back.

#### Type `Deque[T any]`

- **Constructor:**
  ```go
  func New[T any](initialCapacity int) *Deque[T]
  ```

- **Methods:**
    - `PushFront(item T)`: Adds an item to the front of the deque.
    - `PushBack(item T)`: Adds an item to the back of the deque.
    - `PopFront() (T, bool)`: Removes and returns the item at the front of the deque.
    - `PopBack() (T, bool)`: Removes and returns the item at the back of the deque.
    - `PeekFront() (T, bool)`: Returns the item at the front of the deque without removing it.
    - `PeekBack() (T, bool)`: Returns the item at the back of the deque without removing it.
    - `Len() int`: Returns the number of items in the deque.
    - `IsEmpty() bool`: Checks if the deque is empty.
    - `Clear()`: Removes all items from the deque.

### LinkedList

A singly linked list where elements can be added or removed from both the front and the end.

#### Type `LinkedList[T any]`

- **Constructor:**
  ```go
  func New[T any]() *LinkedList[T]
  ```

- **Methods:**
    - `AddFront(value T)`: Adds a new node with the given value to the front of the list.
    - `AddBack(value T)`: Adds a new node with the given value to the end of the list.
    - `RemoveFront() (T, bool)`: Removes the node from the front of the list and returns its value.
    - `RemoveBack() (T, bool)`: Removes the node from the end of the list and returns its value.
    - `Iterate(fn func(T) bool)`: Iterates over the linked list and applies a function to each node's value.
    - `Size() int`: Returns the number of elements in the list.

### Queue

A FIFO (first-in, first-out) queue that supports basic queue operations.

#### Type `Queue[T any]`

- **Constructor:**
  ```go
  func New[T any](initialCapacity int) *Queue[T]
  ```

- **Methods:**
    - `Enqueue(item T)`: Adds an item to the end of the queue.
    - `Dequeue() (T, bool)`: Removes and returns the item at the front of the queue.
    - `Peek() (T, bool)`: Returns the item at the front of the queue without removing it.
    - `Len() int`: Returns the number of items currently in the queue.
    - `IsEmpty() bool`: Checks if the queue is empty.
    - `Clear()`: Removes all items from the queue.

### Stack

A LIFO (last-in, first-out) stack that supports standard stack operations.

#### Type `Stack[T any]`

- **Constructor:**
  ```go
  func New[T any](initialCapacity int) *Stack[T]
  ```

- **Methods:**
    - `Push(item T)`: Adds an item to the top of the stack.
    - `Pop() (T, bool)`: Removes and returns the item from the top of the stack.
    - `Peek() (T, bool)`: Returns the item at the top of the stack without removing it.
    - `Len() int`: Returns the number of items currently in the stack.
    - `IsEmpty() bool`: Checks if the stack is empty.
    - `Clear()`: Removes all items from the stack.
- Here's the README following the provided example format strictly:

### Trie

A Trie (prefix tree) data structure that supports insertion and search operations for words and prefixes.

#### Type `Trie`

- **Constructor:**
  ```go
  func New() *Trie
  ```

- **Methods:**
  - `Insert(word string)`: Adds a word to the Trie.
  - `Search(word string) bool`: Checks if the word exists in the Trie.
  - `StartsWith(prefix string) bool`: Checks if there is any word in the Trie that starts with the given prefix.

## License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.
