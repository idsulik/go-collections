package timedeque

import (
	"testing"
	"time"
)

func TestTimedDequeBasicOperations(t *testing.T) {
	// Create a timed deque with a relatively long TTL to test basic operations
	td := New[int](time.Hour)

	// Test initial state
	if !td.IsEmpty() {
		t.Error("New TimedDeque should be empty")
	}

	if td.Len() != 0 {
		t.Errorf("Expected length 0, got %d", td.Len())
	}

	// Test PushBack and PeekFront
	td.PushBack(1)
	td.PushBack(2)

	if td.IsEmpty() {
		t.Error("TimedDeque should not be empty after pushing items")
	}

	if td.Len() != 2 {
		t.Errorf("Expected length 2, got %d", td.Len())
	}

	val, ok := td.PeekFront()
	if !ok || val != 1 {
		t.Errorf("PeekFront() = %d, %v; want 1, true", val, ok)
	}

	// Test PushFront and PeekBack
	td.PushFront(3)

	val, ok = td.PeekFront()
	if !ok || val != 3 {
		t.Errorf("PeekFront() after PushFront = %d, %v; want 3, true", val, ok)
	}

	val, ok = td.PeekBack()
	if !ok || val != 2 {
		t.Errorf("PeekBack() = %d, %v; want 2, true", val, ok)
	}

	// Test PopFront
	val, ok = td.PopFront()
	if !ok || val != 3 {
		t.Errorf("PopFront() = %d, %v; want 3, true", val, ok)
	}

	if td.Len() != 2 {
		t.Errorf("Expected length 2 after PopFront, got %d", td.Len())
	}

	// Test PopBack
	val, ok = td.PopBack()
	if !ok || val != 2 {
		t.Errorf("PopBack() = %d, %v; want 2, true", val, ok)
	}

	if td.Len() != 1 {
		t.Errorf("Expected length 1 after PopBack, got %d", td.Len())
	}

	// Test Clear
	td.Clear()
	if !td.IsEmpty() {
		t.Error("TimedDeque should be empty after Clear")
	}
}

func TestTimedDequeExpiry(t *testing.T) {
	// Create a timed deque with a short TTL
	shortTTL := 50 * time.Millisecond
	td := New[string](shortTTL)

	// Add items
	td.PushBack("item1")
	td.PushBack("item2")

	// Check initial state
	if td.Len() != 2 {
		t.Errorf("Expected initial length 2, got %d", td.Len())
	}

	// Wait for items to expire
	time.Sleep(shortTTL + 10*time.Millisecond)

	// Check if items expired correctly
	if !td.IsEmpty() {
		t.Errorf("Expected TimedDeque to be empty after TTL, but has %d items", td.Len())
	}

	// Test adding new items after expiry
	td.PushBack("new-item")

	if td.Len() != 1 {
		t.Errorf("Expected length 1 after adding new item, got %d", td.Len())
	}

	val, ok := td.PeekFront()
	if !ok || val != "new-item" {
		t.Errorf("PeekFront() = %s, %v; want new-item, true", val, ok)
	}
}

func TestTimedDequeMixedExpiry(t *testing.T) {
	// Create a timed deque with a medium TTL
	mediumTTL := 100 * time.Millisecond
	td := New[int](mediumTTL)

	// Add initial items
	td.PushBack(1)
	td.PushBack(2)

	// Wait for half the TTL
	time.Sleep(mediumTTL / 2)

	// Add more items
	td.PushBack(3)
	td.PushBack(4)

	// At this point, no items should have expired yet
	if td.Len() != 4 {
		t.Errorf("Expected all 4 items to still be valid, got %d", td.Len())
	}

	// Wait for the first batch to expire, but not the second
	time.Sleep(mediumTTL/2 + 10*time.Millisecond)

	// Now the first two items should have expired
	if td.Len() != 2 {
		t.Errorf("Expected 2 items to remain valid, got %d", td.Len())
	}

	val, ok := td.PeekFront()
	if !ok || val != 3 {
		t.Errorf("PeekFront() = %d, %v; want 3, true", val, ok)
	}

	// Wait for all items to expire
	time.Sleep(mediumTTL)

	if !td.IsEmpty() {
		t.Error("Expected all items to expire")
	}
}

func TestTimedDequeZeroTTL(t *testing.T) {
	// Zero or negative TTL means items never expire
	td := New[int](0)

	td.PushBack(1)
	td.PushBack(2)

	// Wait a bit
	time.Sleep(50 * time.Millisecond)

	// Items should still be there
	if td.Len() != 2 {
		t.Errorf("Expected items not to expire with zero TTL, got length %d", td.Len())
	}

	// Test with negative TTL
	td = New[int](-1 * time.Second)

	td.PushBack(1)

	// Wait a bit
	time.Sleep(50 * time.Millisecond)

	// Items should still be there
	if td.Len() != 1 {
		t.Errorf("Expected items not to expire with negative TTL, got length %d", td.Len())
	}
}

func TestTimedDequeChangeTTL(t *testing.T) {
	// Start with a long TTL
	td := New[int](time.Hour)

	td.PushBack(1)
	td.PushBack(2)

	// Change to a short TTL
	shortTTL := 50 * time.Millisecond
	td.SetTTL(shortTTL)

	// TTL should be updated
	if td.GetTTL() != shortTTL {
		t.Errorf("Expected TTL to be %v, got %v", shortTTL, td.GetTTL())
	}

	// Wait for items to expire with new TTL
	time.Sleep(shortTTL + 10*time.Millisecond)

	// Items should have expired with the new TTL
	if !td.IsEmpty() {
		t.Errorf("Expected items to expire after changing TTL, got length %d", td.Len())
	}
}

func TestTimedDequeGetItems(t *testing.T) {
	td := New[string](time.Hour)

	// Add some items
	items := []string{"item1", "item2", "item3"}
	for _, item := range items {
		td.PushBack(item)
	}

	// Get all items
	gotItems := td.GetItems()

	if len(gotItems) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(gotItems))
	}

	for i, item := range items {
		if gotItems[i] != item {
			t.Errorf("At index %d, expected %s, got %s", i, item, gotItems[i])
		}
	}
}

func TestTimedDequeRemoveExpired(t *testing.T) {
	// Create a timed deque with a short TTL
	shortTTL := 50 * time.Millisecond
	td := New[int](shortTTL)

	// Add initial items
	td.PushBack(1)
	td.PushBack(2)

	// Wait for partial expiry
	time.Sleep(shortTTL / 2)

	// Add more items
	td.PushBack(3)
	td.PushFront(0) // This goes at the front but should not expire yet

	// Wait a bit more so the first items expire but not the new ones
	time.Sleep(shortTTL/2 + 10*time.Millisecond)

	// Call explicit RemoveExpired
	td.RemoveExpired()

	// Only the last two items should remain
	if td.Len() != 2 {
		t.Errorf("Expected 2 items after RemoveExpired, got %d", td.Len())
	}

	items := td.GetItems()
	if len(items) != 2 {
		t.Errorf("Expected 2 items in the slice, got %d", len(items))
	} else {
		// The items should be 0 and 3 (in that order)
		if items[0] != 0 || items[1] != 3 {
			t.Errorf("Expected items [0, 3], got %v", items)
		}
	}
}

func TestTimedDequeClone(t *testing.T) {
	// Create a timed deque with some items
	td := New[int](time.Hour)
	td.PushBack(1)
	td.PushBack(2)

	// Clone it
	clone := td.Clone()

	// Verify the clone has the same items
	if clone.Len() != td.Len() {
		t.Errorf("Clone length %d differs from original %d", clone.Len(), td.Len())
	}

	// Verify the clone has the same TTL
	if clone.GetTTL() != td.GetTTL() {
		t.Errorf("Clone TTL %v differs from original %v", clone.GetTTL(), td.GetTTL())
	}

	// Modify the original
	td.PushBack(3)
	td.SetTTL(time.Minute)

	// Verify the clone is independent
	if clone.Len() == td.Len() {
		t.Error("Clone should be independent of original")
	}

	if clone.GetTTL() == td.GetTTL() {
		t.Error("Clone TTL should be independent of original")
	}
}

func TestTimedDequeCapacity(t *testing.T) {
	// Test with custom capacity
	td := NewWithCapacity[int](time.Hour, 100)

	if cap := td.Cap(); cap < 100 {
		t.Errorf("Expected capacity at least 100, got %d", cap)
	}

	// Add a lot of items
	for i := 0; i < 50; i++ {
		td.PushBack(i)
	}

	// Clear and verify capacity is preserved
	td.Clear()

	if td.Len() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", td.Len())
	}

	if cap := td.Cap(); cap < 100 {
		t.Errorf("Expected capacity at least 100 after Clear, got %d", cap)
	}
}

func TestTimedDequeEmptyOperations(t *testing.T) {
	td := New[string](time.Hour)

	// Operations on empty deque
	val, ok := td.PopFront()
	if ok || val != "" {
		t.Errorf("PopFront() on empty deque = %s, %v; want \"\", false", val, ok)
	}

	val, ok = td.PopBack()
	if ok || val != "" {
		t.Errorf("PopBack() on empty deque = %s, %v; want \"\", false", val, ok)
	}

	val, ok = td.PeekFront()
	if ok || val != "" {
		t.Errorf("PeekFront() on empty deque = %s, %v; want \"\", false", val, ok)
	}

	val, ok = td.PeekBack()
	if ok || val != "" {
		t.Errorf("PeekBack() on empty deque = %s, %v; want \"\", false", val, ok)
	}
}

func TestTimedDequeCustomTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	td := New[Person](time.Hour)

	p1 := Person{"Alice", 30}
	p2 := Person{"Bob", 25}

	td.PushBack(p1)
	td.PushBack(p2)

	person, ok := td.PeekFront()
	if !ok || person.Name != "Alice" {
		t.Errorf("PeekFront() = %+v, %v; want %+v, true", person, ok, p1)
	}

	// Pop and verify
	person, ok = td.PopFront()
	if !ok || person.Name != "Alice" {
		t.Errorf("PopFront() = %+v, %v; want %+v, true", person, ok, p1)
	}

	person, ok = td.PopFront()
	if !ok || person.Name != "Bob" {
		t.Errorf("Second PopFront() = %+v, %v; want %+v, true", person, ok, p2)
	}
}

func BenchmarkTimedDequePushPop(b *testing.B) {
	td := New[int](time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		td.PushBack(i)
		td.PopFront()
	}
}

func BenchmarkTimedDequeRemoveExpired(b *testing.B) {
	td := New[int](time.Millisecond * 100)

	// Add a bunch of items
	for i := 0; i < 1000; i++ {
		td.PushBack(i)
	}

	// Wait for some to expire
	time.Sleep(time.Millisecond * 50)

	// Add more items
	for i := 0; i < 1000; i++ {
		td.PushBack(i + 1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		td.RemoveExpired()
	}
}
