# go-collections

[![Go Report Card](https://goreportcard.com/badge/github.com/idsulik/go-collections)](https://goreportcard.com/report/github.com/idsulik/go-collections)
![Build Status](https://img.shields.io/github/actions/workflow/status/idsulik/go-collections/go.yaml?branch=main)
[![Version](https://img.shields.io/github/v/release/idsulik/go-collections)](https://github.com/idsulik/go-collections/releases)
[![License](https://img.shields.io/github/license/idsulik/go-collections)](https://github.com/idsulik/go-collections/blob/main/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/idsulik/go-collections)](https://pkg.go.dev/github.com/idsulik/go-collections)

`go-collections` is a Go package that provides implementations of common data structures, including a double-ended queue (Deque), various stack implementations, a linked list, a queue, a trie, a priority queue, a binary search tree, a skip list, and a graph. This package offers a simple and efficient way to use these structures in Go, with support for generic types.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Data Structures](#data-structures)
  - [Set](#set)
  - [Deque](#deque)
  - [TimedDeque](#timed-deque)
  - [LinkedList](#linkedlist)
  - [Queue](#queue)
  - [Stack Interface](#stack-interface)
  - [ArrayStack](#arraystack)
  - [LinkedListStack](#linkedliststack)
  - [Trie](#trie)
  - [Priority Queue](#priority-queue)
  - [Binary Search Tree](#binary-search-tree)
  - [Skip List](#skip-list)
  - [Graph](#graph)
  - [BloomFilter](#bloom-filter)
  - [RingBuffer (Circular Buffer)](#ring-buffer)
  - [SegmentTree](#segment-tree)
  - [DisjointSet (UnionFind)](#disjoint-set)
  - [AVL Tree](#avl-tree)
  - [RedBlack Tree](#redblack-tree)
  - [Iterator Interface](#iterator-interface)
4. [Performance Comparison](#performance-comparison)
5. [Contributing](#contributing)
6. [License](#license)

## [Installation](#installation)

You can install the package using the Go module system:

```sh
go get github.com/idsulik/go-collections/v3
```

## [Usage](#usage)

Here is a brief example of how to use the `Deque`:

```go
package main

import (
  "fmt"
  "github.com/idsulik/go-collections/v3/deque"
)

func main() {
  d := deque.New[int](0)
  d.PushBack(1)
  d.PushFront(2)

  front, _ := d.PopFront()
  back, _ := d.PopBack()

  fmt.Println(front) // Output: 2
  fmt.Println(back)  // Output: 1
}
```

## [Data Structures](#data-structures)

### [Set](#set)

A set represents a collection of unique items.

#### Type `Set[T comparable]`

- **Constructor:**

  ```go
  func New[T comparable]() *Set[T]
  ```

- **Methods:**

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

### [Deque](#deque)

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
  - `GetItems() []T`: Returns a slice of all items in the deque in order.
  - `Clone() *Deque[T]`: Returns a deep copy of the deque.
  - `ForEach(fn func(T))`: Applies a function to each item in the deque.

---

### [TimedDeque](#timed-deque)

A TimedDeque extends the Deque with time-to-live (TTL) functionality, automatically removing expired items. It's useful for caching, session management, and timed operations.

#### Type `TimedDeque[T any]`

- **Constructor:**

  ```go
  func New[T any](ttl time.Duration) *TimedDeque[T]
  ```

  ```go
  func NewWithCapacity[T any](ttl time.Duration, capacity int) *TimedDeque[T]
  ```

- **Methods:**

  - `PushFront(item T)`: Adds an item to the front with the current timestamp.
  - `PushBack(item T)`: Adds an item to the back with the current timestamp.
  - `PopFront() (T, bool)`: Removes expired items, then returns the front item.
  - `PopBack() (T, bool)`: Removes expired items, then returns the back item.
  - `PeekFront() (T, bool)`: Removes expired items, then returns the front item without removing it.
  - `PeekBack() (T, bool)`: Removes expired items, then returns the back item without removing it.
  - `Len() int`: Returns the number of non-expired items.
  - `Cap() int`: Returns the capacity of the underlying deque.
  - `IsEmpty() bool`: Checks if the deque is empty after removing expired items.
  - `Clear()`: Removes all items.
  - `GetItems() []T`: Returns a slice of all non-expired items.
  - `Clone() *TimedDeque[T]`: Returns a deep copy.
  - `SetTTL(ttl time.Duration)`: Updates the TTL and removes expired items.
  - `GetTTL() time.Duration`: Returns the current TTL.
  - `RemoveExpired()`: Removes all expired items from the deque.

#### Performance Characteristics:

- Time Complexity: Similar to Deque, with additional O(n) for thorough expiration checks
- Space Complexity: O(n) for n items

### [LinkedList](#linkedlist)

A singly linked list where elements can be added or removed from both the front and the end.

#### Type `LinkedList[T any]`

- **Constructor:**

  ```go
  func New[T any]() *LinkedList[T]
  ```

- **Methods:**

  - `AddFront(value T)`: Adds a new node with the given value to the front of the list.
  - `AddBack(value T)`: Adds a new node with the given value to the end of the list.
  - `PeekFront() (T, bool)`: Returns the value of the node at the front without removing it.
  - `PeekBack() (T, bool)`: Returns the value of the node at the end without removing it.
  - `RemoveFront() (T, bool)`: Removes the node from the front and returns its value.
  - `RemoveBack() (T, bool)`: Removes the node from the end and returns its value.
  - `Iterate(fn func(T) bool)`: Iterates over the list and applies a function to each node's value until the function returns false or the end is reached.
  - `IsEmpty() bool`: Checks if the list is empty.
  - `Len() int`: Returns the number of elements in the list.
  - `Clear()`: Removes all elements from the list.
  - `Iterator() iterator.Iterator[T]`: Returns an iterator for the list.
  - `ForEach(fn func(T))`: Applies a function to each element in the list.

---

### [Queue](#queue)

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
  - `Iterator() iterator.Iterator[T]`: Returns an iterator for the queue.
  - `ForEach(fn func(T))`: Applies a function to each item in the queue.
  - `GetItems() []T`: Returns a slice of all items in the queue.

---

### [Stack Interface](#stack-interface)

An interface representing a LIFO (last-in, first-out) stack.

#### Type `Stack[T any]`

- **Methods:**

  - `Push(item T)`: Adds an item to the top of the stack.
  - `Pop() (T, bool)`: Removes and returns the item from the top of the stack.
  - `Peek() (T, bool)`: Returns the item at the top without removing it.
  - `Len() int`: Returns the number of items in the stack.
  - `IsEmpty() bool`: Checks if the stack is empty.
  - `Clear()`: Removes all items from the stack.

---

### [ArrayStack](#arraystack)

An array-based stack implementation using a slice.

#### Type `ArrayStack[T any]`

- **Constructor:**

  ```go
  func New[T any](initialCapacity int) *ArrayStack[T]
  ```

- **Implements:** `Stack[T]`

- **Methods:**

  - `Push(item T)`: Adds an item to the top of the stack.
  - `Pop() (T, bool)`: Removes and returns the item from the top.
  - `Peek() (T, bool)`: Returns the item at the top without removing it.
  - `Len() int`: Returns the number of items currently in the stack.
  - `IsEmpty() bool`: Checks if the stack is empty.
  - `Clear()`: Removes all items, leaving the stack empty.

---

### [LinkedListStack](#linkedliststack)

A linked list-based stack implementation.

#### Type `LinkedListStack[T any]`

- **Constructor:**

  ```go
  func New[T any]() *LinkedListStack[T]
  ```

- **Implements:** `Stack[T]`

- **Methods:**

  - `Push(item T)`: Adds an item to the top of the stack.
  - `Pop() (T, bool)`: Removes and returns the item from the top.
  - `Peek() (T, bool)`: Returns the item at the top without removing it.
  - `Len() int`: Returns the number of items currently in the stack.
  - `IsEmpty() bool`: Checks if the stack is empty.
  - `Clear()`: Removes all items from the stack.

---

### [Trie](#trie)

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

---

### [Priority Queue](#priorityqueue)

A priority queue allows for efficient retrieval and removal of the highest (or lowest) priority element. It's commonly used in algorithms like Dijkstra's shortest path and task scheduling.

#### Type `PriorityQueue[T any]`

- **Constructor:**

  ```go
  func New[T any](less func(a, b T) bool) *PriorityQueue[T]
  ```

  - `less`: A comparison function that determines the priority of elements. If `less(a, b)` returns `true`, then `a` has higher priority than `b`.

- **Methods:**

  - `Push(item T)`: Adds an item to the priority queue.
  - `Pop() (T, bool)`: Removes and returns the highest priority item.
  - `Peek() (T, bool)`: Returns the highest priority item without removing it.
  - `Len() int`: Returns the number of items in the priority queue.
  - `IsEmpty() bool`: Checks if the priority queue is empty.
  - `Clear()`: Removes all items from the priority queue.
  - `PopFunc(fn func(T) bool) (T, bool)`: Removes and returns the first item that satisfies the function.
  - `LenFunc(fn func(T) bool) int`: Returns the number of items that satisfy the function.
  - `Contains(item T) bool`: Checks if an item exists in the queue.
  - `PushIfAbsent(item T) bool`: Adds an item only if it's not already present.
  - `Remove(item T) bool`: Removes the first occurrence of the item.
  - `RemoveFunc(fn func(T) bool) bool`: Removes the first item that satisfies the function.
  - `Clone() *PriorityQueue[T]`: Returns a deep copy of the priority queue.
  - `Keys() []T`: Returns a slice of all items in the queue.
  - `Vals() []T`: Alias for Keys().

---

### [Binary Search Tree](#binarysearchtree)

A Binary Search Tree (BST) maintains elements in sorted order, allowing for efficient insertion, deletion, and lookup operations. Each node has at most two children, with left child values less than the parent and right child values greater.

#### Type `BST[T Ordered]`

- **Constructor:**

  ```go
  func New[T Ordered]() *BST[T]
  ```

  - `T Ordered`: A type constraint that ensures the elements can be compared using `<` and `>` operators. Supported types include integers, floats, and strings.

- **Methods:**

  - `Insert(value T)`: Inserts a value into the BST.
  - `Remove(value T)`: Removes a value from the BST.
  - `Contains(value T) bool`: Checks if a value exists in the BST.
  - `InOrderTraversal(fn func(T))`: Traverses the BST in order and applies a function to each node's value.
  - `Len() int`: Returns the number of nodes in the BST.
  - `IsEmpty() bool`: Checks if the BST is empty.
  - `Clear()`: Removes all nodes from the BST.

---

### [Skip List](#skiplist)

A Skip List is a probabilistic data structure that allows fast search, insertion, and deletion operations within an ordered sequence of elements. It achieves efficiency by maintaining multiple levels of linked lists, where each higher level skips over a larger number of elements, allowing operations to be performed in O(log n) average time.

#### Type `SkipList[T Ordered]`

- **Constructor:**

  ```go
  func New[T Ordered](maxLevel int, p float64) *SkipList[T]
  ```

  - `maxLevel`: The maximum level of the skip list (controls the space vs. time trade-off).
  - `p`: The probability factor used to determine the level of new nodes (usually set to 0.5).

- **Methods:**

  - `Insert(value T)`: Inserts a value into the skip list.
  - `Delete(value T)`: Deletes a value from the skip list.
  - `Search(value T) bool`: Searches for a value in the skip list.
  - `Len() int`: Returns the number of elements in the skip list.
  - `IsEmpty() bool`: Checks if the skip list is empty.
  - `Clear()`: Removes all elements from the skip list.

---
### [Graph](#graph)

Represents networks of nodes and edges, suitable for various algorithms like search, shortest path, and spanning trees.

#### Type `Graph[T comparable]`

- **Constructor:**

  ```go
  func New[T comparable](directed bool) *Graph[T]
  ```

  - `directed`: Specifies whether the graph is directed or undirected.

- **Methods:**

  - `AddNode(value T)`: Adds a node to the graph.
  - `AddEdge(from, to T, weight float64)`: Adds an edge between two nodes with an optional weight.
  - `RemoveNode(value T)`: Removes a node and all connected edges.
  - `RemoveEdge(from, to T)`: Removes an edge between two nodes.
  - `Neighbors(value T) []T`: Returns adjacent nodes.
  - `HasNode(value T) bool`: Checks if a node exists.
  - `HasEdge(from, to T) bool`: Checks if an edge exists.
  - `GetEdgeWeight(from, to T) (float64, bool)`: Retrieves the weight of the edge between two nodes.
  - `Traverse(start T, visit func(T))`: Traverses the graph from a starting node using Breadth-First Search.
  - `Nodes() []T`: Returns a slice of all node values in the graph.
  - `Edges() [][2]T`: Returns a slice of all edges in the graph.
  - `Iterator() iterator.Iterator[T]`: Returns an iterator for the graph.
  - `ForEach(fn func(T))`: Applies a function to each node in the graph.

---
### [Bloom Filter](#bloom-filter)

A Bloom Filter is a space-efficient probabilistic data structure used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not. Elements can be added to the set, but not removed.

#### Type `BloomFilter[T any]`

- **Constructor:**

  ```go
  func NewBloomFilter[T any](expectedItems uint, falsePositiveProb float64) *BloomFilter[T]
  ```

  - `expectedItems`: Expected number of items to be added to the filter
  - `falsePositiveProb`: Desired false positive probability (between 0 and 1)

- **Methods:**

  - `Add(item T)`: Adds an item to the Bloom Filter.
  - `Contains(item T) bool`: Tests whether an item might be in the set.
  - `EstimatedFalsePositiveRate() float64`: Returns the current estimated false positive rate.
  - `Clear()`: Removes all items from the Bloom Filter.
  - `Len() int`: Returns the number of items added to the Bloom Filter.
  - `IsEmpty() bool`: Returns true if no items have been added.
  - `BitSize() uint`: Returns the size of the underlying bit array.
  - `NumberOfHashes() uint`: Returns the number of hash functions being used.

#### Performance Characteristics:

- Space Complexity: O(m), where m is the size of the bit array
- Time Complexity:
  - Add: O(k), where k is the number of hash functions
  - Contains: O(k), where k is the number of hash functions
- False Positive Probability: (1 - e^(-kn/m))^k
  - k: number of hash functions
  - n: number of inserted elements
  - m: size of bit array

#### Use Cases:

- Duplicate detection
- Cache filtering
- URL shorteners
- Spell checkers
- Network routing
- Database query optimization

---
### [Ring Buffer](#ring-buffer)

A Ring Buffer (also known as a Circular Buffer) is a fixed-size buffer that wraps around to the beginning when it reaches the end. It's particularly useful for streaming data processing, implementing queues with size limits, and managing buffers in embedded systems.

#### Type `RingBuffer[T any]`

- **Constructor:**

```go
  func New[T any](capacity int) *RingBuffer[T]
  ```

- **Methods:**

- `Write(item T) bool`: Adds an item to the buffer. Returns false if the buffer is full.
- `Read() (T, bool)`: Removes and returns the oldest item from the buffer.
- `Peek() (T, bool)`: Returns the oldest item without removing it.
- `IsFull() bool`: Returns true if the buffer is at capacity.
- `IsEmpty() bool`: Returns true if the buffer contains no items.
- `Cap() int`: Returns the total capacity of the buffer.
- `Len() int`: Returns the current number of items in the buffer.
- `Clear()`: Removes all items from the buffer.

#### Performance Characteristics:

- Space Complexity: O(n), where n is the buffer capacity
- Time Complexity:
  - Write: O(1)
  - Read: O(1)
  - Peek: O(1)
  - Clear: O(1)

---
### [Segment Tree](#segment-tree)

A Segment Tree is a versatile data structure that supports various range query operations (sum, minimum, maximum, GCD, etc.) with efficient updates. It provides O(log n) complexity for both range queries and point updates.

#### Type `SegmentTree[T any]`

- **Constructor:**

  ```go
  func NewSegmentTree[T any](arr []T, identity T, combine Operation[T]) *SegmentTree[T]
  ```

  - `arr`: Initial array of elements
  - `identity`: Identity element for the operation (e.g., 0 for sum, Inf for min)
  - `combine`: Function that defines how to combine elements (e.g., addition for sum queries)

- **Methods:**

  - `Update(index int, value T)`: Updates the value at the given index
  - `Query(left, right int) T`: Returns the result of the operation for the range [left, right]

#### Performance Characteristics:

- Construction: O(n)
- Range Query: O(log n)
- Point Update: O(log n)
- Space Complexity: O(n)

---
### [Disjoint Set](#disjoint-set)

A Disjoint Set (also known as Union-Find) is a data structure that keeps track of elements partitioned into non-overlapping subsets. It provides near-constant-time operations to merge sets and determine if two elements belong to the same set.

#### Type `DisjointSet[T comparable]`

- **Constructor:**

  ```go
  func New[T comparable]() *DisjointSet[T]
  ```

- **Methods:**

  - `MakeSet(x T)`: Creates a new set containing a single element.
  - `Find(x T) T`: Returns the representative element of the set containing x.
  - `Union(x, y T)`: Merges the sets containing elements x and y.
  - `Connected(x, y T) bool`: Returns true if elements x and y are in the same set.
  - `Clear()`: Removes all elements from the disjoint set.
  - `Len() int`: Returns the number of elements in the disjoint set.
  - `IsEmpty() bool`: Returns true if the disjoint set contains no elements.
  - `GetSets() map[T][]T`: Returns a map of representatives to their set members.

#### Performance Characteristics:

- MakeSet: O(1)
- Find: O(α(n)) amortized (nearly constant)
- Union: O(α(n)) amortized (nearly constant)
- Connected: O(α(n)) amortized (nearly constant)

Where α(n) is the inverse Ackermann function, which grows extremely slowly and is effectively constant for all practical values of n.

---

### [AVL Tree](#avl-tree)

An AVL tree is a self-balancing binary search tree where the heights of the two child subtrees of any node differ by at most one. This balancing ensures O(log n) time complexity for insertions, deletions, and lookups.

#### Type `AVLTree[T any]`

- **Constructor:**

  ```go
  func New[T any](compare func(a, b T) int) *AVLTree[T]
  ```

  - `compare`: A comparison function that returns:
    - -1 if a < b
    - 0 if a == b
    - 1 if a > b

- **Methods:**

  - `Insert(value T)`: Adds a value to the tree while maintaining AVL balance
  - `Delete(value T) bool`: Removes a value from the tree while maintaining AVL balance
  - `Search(value T) bool`: Checks if a value exists in the tree
  - `InOrderTraversal(fn func(T))`: Visits all nodes in ascending order
  - `Clear()`: Removes all elements from the tree
  - `Len() int`: Returns the number of nodes in the tree
  - `IsEmpty() bool`: Checks if the tree is empty
  - `Height() int`: Returns the height of the tree

#### Performance Characteristics:

| Operation | Average Case | Worst Case |
|-----------|--------------|------------|
| Space     | O(n)         | O(n)       |
| Search    | O(log n)     | O(log n)   |
| Insert    | O(log n)     | O(log n)   |
| Delete    | O(log n)     | O(log n)   |

---
### [Red-Black Tree](#redblack-tree)
A Red-Black Tree is a self-balancing binary search tree where each node is colored either red or black, following specific rules that maintain balance. This ensures O(log n) time complexity for insertions, deletions, and lookups.

#### Type `RedBlackTree[T any]`
- **Constructor:**
```go
func New[T any](compare func(a, b T) int) *RedBlackTree[T]
```
- *`compare`*: A comparison function that returns:
  - -1 if a < b
  - 0 if a == b
  - 1 if a > b

- **Methods:**
  - `Insert(value T)`: Adds a value to the tree while maintaining Red-Black properties
  - `Delete(value T) bool`: Removes a value from the tree while maintaining Red-Black properties
  - `Search(value T) bool`: Checks if a value exists in the tree
  - `InOrderTraversal(fn func(T))`: Visits all nodes in ascending order
  - `Clear()`: Removes all elements from the tree
  - `Len() int`: Returns the number of nodes in the tree
  - `IsEmpty() bool`: Checks if the tree is empty
  - `Height() int`: Returns the height of the tree

#### Performance Characteristics:
| Operation | Average Case | Worst Case |
|-----------|--------------|------------|
| Space     | O(n)         | O(n)       |
| Search    | O(log n)     | O(log n)   |
| Insert    | O(log n)     | O(log n)   |
| Delete    | O(log n)     | O(log n)   |

Where n is the number of nodes in the tree.

#### Balance Property:
The Red-Black tree maintains the following invariants:
1. Every node is either red or black
2. The root is always black
3. All leaves (NIL nodes) are black
4. If a node is red, both its children are black
5. Every path from root to leaf contains the same number of black nodes

#### Use Cases:
1. **Database Systems:**
  - Index structures implementation
  - Multi-version concurrency control

2. **Memory Management:**
  - Virtual memory segment trees
  - Memory allocation tracking

3. **File Systems:**
  - Directory organization
  - File system journaling

4. **Process Scheduling:**
  - Task prioritization
  - Real-time scheduling

5. **Programming Languages:**
  - Symbol table implementation
  - Garbage collection algorithms
---

### [Iterator Interface](#iterator-interface)

An interface for iterating over collections in a standardized way.

#### Type `Iterator[T any]`

- **Methods:**

  - `HasNext() bool`: Returns true if there are more elements to iterate over.
  - `Next() (T, bool)`: Returns the next element in the iteration and false if there are no more elements.
  - `Reset()`: Restarts the iteration from the beginning.

#### Type `Iterable[T any]`

- **Methods:**

  - `Iterator() Iterator[T]`: Returns a new iterator for the collection.

## [Performance Comparison](#performance-comparison)

| Data Structure  | Access   | Search   | Insertion | Deletion | Space                    |
|-----------------|----------|----------|-----------|----------|--------------------------|
| Array           | O(1)     | O(n)     | O(n)      | O(n)     | O(n)                     |
| Set             | O(1)     | O(1)     | O(1)      | O(1)     | O(n)                     |
| Queue           | O(1)     | O(n)     | O(1)*     | O(1)*    | O(n)                     |
| Priority Queue  | O(1)     | O(1)     | O(log n)  | O(log n) | O(n)                     |
| BST (balanced)  | O(log n) | O(log n) | O(log n)  | O(log n) | O(n)                     |
| Trie            | O(m)     | O(m)     | O(m)      | O(m)     | O(ALPHABET_SIZE * m * n) |
| Graph           | O(1)     | O(V+E)   | O(E)      | O(E)     | O(V+E)                   |
| BloomFilter     | N/A      | O(k)     | O(k)      | N/A      | O(m)                     |
| Disjoint Set    | O(α(n))  | O(α(n))  | O(α(n))   | O(α(n))  | O(n)                     |
| RingBuffer      | O(1)     | O(n)     | O(1)      | O(1)     | O(n)                     |
| SkipList        | O(1)     | O(log n) | O(log n)  | O(log n) | O(n log n)               |
| LinkedList      | O(n)     | O(n)     | O(1)      | O(1)**   | O(n)                     |
| SegmentTree     | O(1)     | O(log n) | O(log n)  | O(log n) | O(n)                     |
| ArrayStack      | O(1)     | O(n)     | O(1)*     | O(1)*    | O(n)                     |
| LinkedListStack | O(1)     | O(n)     | O(1)      | O(1)     | O(n)                     |
| Deque           | O(1)     | O(n)     | O(1)*     | O(1)*    | O(n)                     |
| TimedDeque      | O(1)     | O(n)     | O(1)*     | O(1)*    | O(n)                     |                    |
| AVL Tree        | O(log n) | O(log n) | O(log n)  | O(log n) | O(n)                     |
| Red-Black Tree  | O(log n) | O(log n) | O(log n)  | O(log n) | O(n)                     |

Where:
- n is the number of elements
- m is the length of the string/key
- k is the number of hash functions
- V is the number of vertices
- E is the number of edges
- α(n) is the inverse Ackermann function (effectively constant)
- * indicates amortized time complexity
- ** for LinkedList, deletion is O(1) at front/back but O(n) for arbitrary position

## [Contributing](#contributing)

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## [License](#license)

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## [Acknowledgments](#acknowledgments)

- Thanks to all contributors who have helped shape this library
- Inspired by various Go community projects and standard library patterns



