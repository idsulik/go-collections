package lrucache

import (
	"errors"
	"fmt"
)

// node represents a node in the doubly linked list
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// LRUCache represents a Least Recently Used cache with fixed capacity
type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*node[K, V]
	head     *node[K, V] // dummy head node
	tail     *node[K, V] // dummy tail node
}

// New creates a new LRU cache with the specified capacity
func New[K comparable, V any](capacity int) (*LRUCache[K, V], error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be positive")
	}

	lru := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*node[K, V]),
	}

	// Create dummy head and tail nodes
	lru.head = &node[K, V]{}
	lru.tail = &node[K, V]{}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru, nil
}

// Get retrieves a value from the cache and marks it as recently used
func (lru *LRUCache[K, V]) Get(key K) (V, bool) {
	var zero V

	if node, exists := lru.cache[key]; exists {
		// Move to front (most recently used)
		lru.moveToFront(node)
		return node.value, true
	}

	return zero, false
}

// Put adds or updates a key-value pair in the cache
func (lru *LRUCache[K, V]) Put(key K, value V) {
	if node, exists := lru.cache[key]; exists {
		// Update existing node
		node.value = value
		lru.moveToFront(node)
		return
	}

	// Create new node
	newNode := &node[K, V]{
		key:   key,
		value: value,
	}

	// Add to cache and front of list
	lru.cache[key] = newNode
	lru.addToFront(newNode)

	// Check capacity and evict if necessary
	if len(lru.cache) > lru.capacity {
		lru.evictLRU()
	}
}

// Remove removes a key from the cache
func (lru *LRUCache[K, V]) Remove(key K) bool {
	if node, exists := lru.cache[key]; exists {
		lru.removeNode(node)
		delete(lru.cache, key)
		return true
	}
	return false
}

// Peek retrieves a value without marking it as recently used
func (lru *LRUCache[K, V]) Peek(key K) (V, bool) {
	var zero V

	if node, exists := lru.cache[key]; exists {
		return node.value, true
	}

	return zero, false
}

// Contains checks if a key exists in the cache without affecting its position
func (lru *LRUCache[K, V]) Contains(key K) bool {
	_, exists := lru.cache[key]
	return exists
}

// Len returns the current number of items in the cache
func (lru *LRUCache[K, V]) Len() int {
	return len(lru.cache)
}

// Cap returns the capacity of the cache
func (lru *LRUCache[K, V]) Cap() int {
	return lru.capacity
}

// IsEmpty returns true if the cache is empty
func (lru *LRUCache[K, V]) IsEmpty() bool {
	return len(lru.cache) == 0
}

// IsFull returns true if the cache is at capacity
func (lru *LRUCache[K, V]) IsFull() bool {
	return len(lru.cache) == lru.capacity
}

// Clear removes all items from the cache
func (lru *LRUCache[K, V]) Clear() {
	lru.cache = make(map[K]*node[K, V])
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
}

// Keys returns a slice of all keys in the cache, ordered from most to least recently used
func (lru *LRUCache[K, V]) Keys() []K {
	keys := make([]K, 0, len(lru.cache))
	current := lru.head.next

	for current != lru.tail {
		keys = append(keys, current.key)
		current = current.next
	}

	return keys
}

// Values returns a slice of all values in the cache, ordered from most to least recently used
func (lru *LRUCache[K, V]) Values() []V {
	values := make([]V, 0, len(lru.cache))
	current := lru.head.next

	for current != lru.tail {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

// Oldest returns the least recently used key-value pair without removing it
func (lru *LRUCache[K, V]) Oldest() (K, V, bool) {
	var zeroK K
	var zeroV V

	if lru.IsEmpty() {
		return zeroK, zeroV, false
	}

	oldest := lru.tail.prev
	return oldest.key, oldest.value, true
}

// Newest returns the most recently used key-value pair without removing it
func (lru *LRUCache[K, V]) Newest() (K, V, bool) {
	var zeroK K
	var zeroV V

	if lru.IsEmpty() {
		return zeroK, zeroV, false
	}

	newest := lru.head.next
	return newest.key, newest.value, true
}

// Resize changes the capacity of the cache, evicting items if necessary
func (lru *LRUCache[K, V]) Resize(newCapacity int) error {
	if newCapacity <= 0 {
		return errors.New("capacity must be positive")
	}

	lru.capacity = newCapacity

	// Evict items if new capacity is smaller
	for len(lru.cache) > lru.capacity {
		lru.evictLRU()
	}

	return nil
}

// ForEach iterates over all key-value pairs in the cache from most to least recently used
func (lru *LRUCache[K, V]) ForEach(fn func(key K, value V) bool) {
	current := lru.head.next

	for current != lru.tail {
		if !fn(current.key, current.value) {
			break
		}
		current = current.next
	}
}

// String returns a string representation of the cache
func (lru *LRUCache[K, V]) String() string {
	if lru.IsEmpty() {
		return "LRUCache{}"
	}

	result := "LRUCache{"
	current := lru.head.next
	first := true

	for current != lru.tail {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v:%v", current.key, current.value)
		current = current.next
		first = false
	}

	result += "}"
	return result
}

// moveToFront moves a node to the front of the list (most recently used)
func (lru *LRUCache[K, V]) moveToFront(node *node[K, V]) {
	lru.removeNode(node)
	lru.addToFront(node)
}

// addToFront adds a node to the front of the list
func (lru *LRUCache[K, V]) addToFront(node *node[K, V]) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode removes a node from the list
func (lru *LRUCache[K, V]) removeNode(node *node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// evictLRU removes the least recently used item
func (lru *LRUCache[K, V]) evictLRU() {
	if lru.IsEmpty() {
		return
	}

	oldest := lru.tail.prev
	lru.removeNode(oldest)
	delete(lru.cache, oldest.key)
}
