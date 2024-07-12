# go-collections

`go-collections` is a Go library that provides implementations of common data structures including a double-ended queue (Deque), a linked list, a queue, and a stack. This package offers a simple and efficient way to use these structures in Go, with support for generic types.

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

## License

This project is licensed under the MIT License