package timedeque

import (
	"time"

	"github.com/idsulik/go-collections/v3/deque"
)

const defaultCapacity = 16

// TimedItem wraps an item with its insertion timestamp
type TimedItem[T any] struct {
	Value     T
	Timestamp time.Time
}

// TimedDeque extends the Deque with time-to-live functionality
type TimedDeque[T any] struct {
	deque *deque.Deque[TimedItem[T]]
	ttl   time.Duration
}

// New creates a new TimedDeque with the specified TTL
func New[T any](ttl time.Duration) *TimedDeque[T] {
	return &TimedDeque[T]{
		deque: deque.New[TimedItem[T]](defaultCapacity),
		ttl:   ttl,
	}
}

// NewWithCapacity creates a new TimedDeque with the specified TTL and capacity
func NewWithCapacity[T any](ttl time.Duration, capacity int) *TimedDeque[T] {
	return &TimedDeque[T]{
		deque: deque.New[TimedItem[T]](capacity),
		ttl:   ttl,
	}
}

// PushFront adds an item to the front of the deque with the current timestamp
func (td *TimedDeque[T]) PushFront(item T) {
	td.deque.PushFront(
		TimedItem[T]{
			Value:     item,
			Timestamp: time.Now(),
		},
	)
}

// PushBack adds an item to the back of the deque with the current timestamp
func (td *TimedDeque[T]) PushBack(item T) {
	td.deque.PushBack(
		TimedItem[T]{
			Value:     item,
			Timestamp: time.Now(),
		},
	)
}

// PopFront removes and returns the item at the front of the deque
// First removes any expired items from the front
func (td *TimedDeque[T]) PopFront() (T, bool) {
	td.removeExpiredFront()

	if td.deque.IsEmpty() {
		var zero T
		return zero, false
	}

	item, ok := td.deque.PopFront()
	if !ok {
		var zero T
		return zero, false
	}

	return item.Value, true
}

// PopBack removes and returns the item at the back of the deque
// First removes any expired items from the front
func (td *TimedDeque[T]) PopBack() (T, bool) {
	td.removeExpiredFront()

	if td.deque.IsEmpty() {
		var zero T
		return zero, false
	}

	item, ok := td.deque.PopBack()
	if !ok {
		var zero T
		return zero, false
	}

	return item.Value, true
}

// PeekFront returns the item at the front of the deque without removing it
// First removes any expired items from the front
func (td *TimedDeque[T]) PeekFront() (T, bool) {
	td.removeExpiredFront()

	if td.deque.IsEmpty() {
		var zero T
		return zero, false
	}

	item, ok := td.deque.PeekFront()
	if !ok {
		var zero T
		return zero, false
	}

	return item.Value, true
}

// PeekBack returns the item at the back of the deque without removing it
// First removes any expired items from the front
func (td *TimedDeque[T]) PeekBack() (T, bool) {
	td.removeExpiredFront()

	if td.deque.IsEmpty() {
		var zero T
		return zero, false
	}

	item, ok := td.deque.PeekBack()
	if !ok {
		var zero T
		return zero, false
	}

	return item.Value, true
}

// Len returns the number of items in the deque after removing expired items
func (td *TimedDeque[T]) Len() int {
	td.removeExpiredFront()
	return td.deque.Len()
}

// Cap returns the current capacity of the deque
func (td *TimedDeque[T]) Cap() int {
	return td.deque.Cap()
}

// IsEmpty checks if the deque is empty after removing expired items
func (td *TimedDeque[T]) IsEmpty() bool {
	td.removeExpiredFront()
	return td.deque.IsEmpty()
}

// Clear removes all items from the deque
func (td *TimedDeque[T]) Clear() {
	// Preserve the capacity of the underlying deque
	capacity := td.deque.Cap()
	td.deque = deque.New[TimedItem[T]](capacity)
}

// GetItems returns a slice containing all non-expired items in the deque
func (td *TimedDeque[T]) GetItems() []T {
	td.removeExpiredFront()
	timedItems := td.deque.GetItems()
	items := make([]T, len(timedItems))

	for i, timedItem := range timedItems {
		items[i] = timedItem.Value
	}

	return items
}

// Clone returns a deep copy of the TimedDeque
func (td *TimedDeque[T]) Clone() *TimedDeque[T] {
	return &TimedDeque[T]{
		deque: td.deque.Clone(),
		ttl:   td.ttl,
	}
}

// SetTTL updates the time-to-live duration and removes expired items
func (td *TimedDeque[T]) SetTTL(ttl time.Duration) {
	td.ttl = ttl
	td.removeExpiredFront()
}

// GetTTL returns the current time-to-live duration
func (td *TimedDeque[T]) GetTTL() time.Duration {
	return td.ttl
}

// IsExpired checks if an item with the given timestamp has expired
func (td *TimedDeque[T]) IsExpired(timestamp time.Time) bool {
	// If TTL is zero or negative, items never expire
	if td.ttl <= 0 {
		return false
	}
	return time.Since(timestamp) > td.ttl
}

// removeExpiredFront removes expired items from the front of the deque
func (td *TimedDeque[T]) removeExpiredFront() {
	// If TTL is zero or negative, items never expire
	if td.ttl <= 0 {
		return
	}

	for !td.deque.IsEmpty() {
		frontItem, ok := td.deque.PeekFront()
		if !ok {
			break
		}

		if time.Since(frontItem.Timestamp) > td.ttl {
			td.deque.PopFront()
		} else {
			break
		}
	}
}

// RemoveExpired removes all expired items from the deque
// This is more thorough than removeExpiredFront but has O(n) complexity
func (td *TimedDeque[T]) RemoveExpired() {
	if td.deque.IsEmpty() {
		return
	}

	// First remove from front (optimization)
	td.removeExpiredFront()

	// If ttl is 0, all items are kept forever
	if td.ttl <= 0 {
		return
	}

	// Check if there are any expired items in the middle or back
	// We'll rebuild the deque if necessary
	timedItems := td.deque.GetItems()
	now := time.Now()
	hasExpired := false

	for _, item := range timedItems {
		if now.Sub(item.Timestamp) > td.ttl {
			hasExpired = true
			break
		}
	}

	if !hasExpired {
		return
	}

	// Rebuild the deque without expired items
	newDeque := deque.New[TimedItem[T]](td.deque.Cap())
	for _, item := range timedItems {
		if now.Sub(item.Timestamp) <= td.ttl {
			newDeque.PushBack(item)
		}
	}

	td.deque = newDeque
}
