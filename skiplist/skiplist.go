package skiplist

import (
	"math/rand"
	"time"

	"github.com/idsulik/go-collections/internal/cmp"
)

// SkipList represents the Skip List.
type SkipList[T cmp.Ordered] struct {
	maxLevel int
	level    int
	p        float64
	header   *node[T]
	length   int
	randSrc  *rand.Rand
}

type node[T cmp.Ordered] struct {
	value T
	next  []*node[T]
}

// New creates a new empty Skip List.
func New[T cmp.Ordered](maxLevel int, p float64) *SkipList[T] {
	return &SkipList[T]{
		maxLevel: maxLevel,
		level:    1,
		p:        p,
		header:   &node[T]{next: make([]*node[T], maxLevel)},
		randSrc:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Insert adds a value into the Skip List.
func (sl *SkipList[T]) Insert(value T) {
	update := make([]*node[T], sl.maxLevel)
	current := sl.header

	// Find positions to update
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	// Check if value already exists
	if current.next[0] != nil && current.next[0].value == value {
		return // Do not insert duplicates
	}

	// Generate a random level for the new node
	newLevel := sl.randomLevel()
	if newLevel > sl.level {
		for i := sl.level; i < newLevel; i++ {
			update[i] = sl.header
		}
		sl.level = newLevel
	}

	// Create new node
	newNode := &node[T]{
		value: value,
		next:  make([]*node[T], newLevel),
	}

	// Insert node and update pointers
	for i := 0; i < newLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}

	sl.length++
}

// Search checks if a value exists in the Skip List.
func (sl *SkipList[T]) Search(value T) bool {
	current := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
	}
	current = current.next[0]
	return current != nil && current.value == value
}

// Delete removes a value from the Skip List.
func (sl *SkipList[T]) Delete(value T) {
	update := make([]*node[T], sl.maxLevel)
	current := sl.header

	// Find positions to update
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	current = current.next[0]
	if current != nil && current.value == value {
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != current {
				break
			}
			update[i].next[i] = current.next[i]
		}

		// Adjust the level if necessary
		for sl.level > 1 && sl.header.next[sl.level-1] == nil {
			sl.level--
		}
		sl.length--
	}
}

// Len returns the number of elements in the Skip List.
func (sl *SkipList[T]) Len() int {
	return sl.length
}

// IsEmpty checks if the Skip List is empty.
func (sl *SkipList[T]) IsEmpty() bool {
	return sl.length == 0
}

// Clear removes all elements from the Skip List.
func (sl *SkipList[T]) Clear() {
	sl.header = &node[T]{next: make([]*node[T], sl.maxLevel)}
	sl.level = 1
	sl.length = 0
}

// randomLevel generates a random level for a new node.
func (sl *SkipList[T]) randomLevel() int {
	level := 1
	for sl.randSrc.Float64() < sl.p && level < sl.maxLevel {
		level++
	}
	return level
}
