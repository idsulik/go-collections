package lrucache

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Valid capacity", func(t *testing.T) {
		cache, err := New[string, int](5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if cache.Cap() != 5 {
			t.Errorf("Expected capacity 5, got %d", cache.Cap())
		}
		if !cache.IsEmpty() {
			t.Error("New cache should be empty")
		}
	})

	t.Run("Invalid capacity", func(t *testing.T) {
		_, err := New[string, int](0)
		if err == nil {
			t.Error("Expected error for zero capacity")
		}

		_, err = New[string, int](-1)
		if err == nil {
			t.Error("Expected error for negative capacity")
		}
	})
}

func TestBasicOperations(t *testing.T) {
	cache, _ := New[string, int](3)

	t.Run("Put and Get", func(t *testing.T) {
		cache.Put("key1", 1)
		cache.Put("key2", 2)

		if val, ok := cache.Get("key1"); !ok || val != 1 {
			t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
		}

		if val, ok := cache.Get("key2"); !ok || val != 2 {
			t.Errorf("Expected (2, true), got (%d, %v)", val, ok)
		}

		if _, ok := cache.Get("nonexistent"); ok {
			t.Error("Expected false for nonexistent key")
		}
	})

	t.Run("Update existing key", func(t *testing.T) {
		cache.Clear()
		cache.Put("key1", 1)
		cache.Put("key1", 10) // Update

		if val, ok := cache.Get("key1"); !ok || val != 10 {
			t.Errorf("Expected (10, true), got (%d, %v)", val, ok)
		}

		if cache.Len() != 1 {
			t.Errorf("Expected length 1, got %d", cache.Len())
		}
	})
}

func TestLRUEviction(t *testing.T) {
	cache, _ := New[string, int](2)

	// Fill cache to capacity
	cache.Put("key1", 1)
	cache.Put("key2", 2)

	// Access key1 to make it most recently used
	cache.Get("key1")

	// Add new item, should evict key2 (least recently used)
	cache.Put("key3", 3)

	if _, ok := cache.Get("key2"); ok {
		t.Error("key2 should have been evicted")
	}

	if _, ok := cache.Get("key1"); !ok {
		t.Error("key1 should still be in cache")
	}

	if _, ok := cache.Get("key3"); !ok {
		t.Error("key3 should be in cache")
	}
}

func TestPeek(t *testing.T) {
	cache, _ := New[string, int](2)
	cache.Put("key1", 1)
	cache.Put("key2", 2)

	// Peek should not affect LRU order
	if val, ok := cache.Peek("key1"); !ok || val != 1 {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}

	// Add new item, key1 should be evicted (since peek didn't move it)
	cache.Put("key3", 3)

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been evicted")
	}
}

func TestRemove(t *testing.T) {
	cache, _ := New[string, int](3)
	cache.Put("key1", 1)
	cache.Put("key2", 2)

	if !cache.Remove("key1") {
		t.Error("Remove should return true for existing key")
	}

	if cache.Remove("key1") {
		t.Error("Remove should return false for non-existing key")
	}

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should be removed")
	}

	if cache.Len() != 1 {
		t.Errorf("Expected length 1, got %d", cache.Len())
	}
}

func TestContains(t *testing.T) {
	cache, _ := New[string, int](2)
	cache.Put("key1", 1)

	if !cache.Contains("key1") {
		t.Error("Contains should return true for existing key")
	}

	if cache.Contains("nonexistent") {
		t.Error("Contains should return false for non-existing key")
	}
}

func TestCapacityAndLength(t *testing.T) {
	cache, _ := New[string, int](3)

	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}

	if cache.Len() != 0 {
		t.Errorf("Expected length 0, got %d", cache.Len())
	}

	cache.Put("key1", 1)
	cache.Put("key2", 2)

	if cache.Len() != 2 {
		t.Errorf("Expected length 2, got %d", cache.Len())
	}

	if cache.IsEmpty() {
		t.Error("Cache should not be empty")
	}

	if cache.IsFull() {
		t.Error("Cache should not be full")
	}

	cache.Put("key3", 3)

	if !cache.IsFull() {
		t.Error("Cache should be full")
	}
}

func TestClear(t *testing.T) {
	cache, _ := New[string, int](3)
	cache.Put("key1", 1)
	cache.Put("key2", 2)

	cache.Clear()

	if !cache.IsEmpty() {
		t.Error("Cache should be empty after clear")
	}

	if cache.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", cache.Len())
	}

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should not exist after clear")
	}
}

func TestKeysAndValues(t *testing.T) {
	cache, _ := New[string, int](3)
	cache.Put("key1", 1)
	cache.Put("key2", 2)
	cache.Put("key3", 3)

	// Access key1 to make it most recently used
	cache.Get("key1")

	keys := cache.Keys()
	values := cache.Values()

	expectedKeys := []string{"key1", "key3", "key2"}
	expectedValues := []int{1, 3, 2}

	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}

	for i, key := range expectedKeys {
		if keys[i] != key {
			t.Errorf("Expected key[%d] = %s, got %s", i, key, keys[i])
		}
	}

	for i, value := range expectedValues {
		if values[i] != value {
			t.Errorf("Expected value[%d] = %d, got %d", i, value, values[i])
		}
	}
}

func TestOldestAndNewest(t *testing.T) {
	cache, _ := New[string, int](3)

	// Test empty cache
	if _, _, ok := cache.Oldest(); ok {
		t.Error("Oldest should return false for empty cache")
	}

	if _, _, ok := cache.Newest(); ok {
		t.Error("Newest should return false for empty cache")
	}

	cache.Put("key1", 1)
	cache.Put("key2", 2)
	cache.Put("key3", 3)

	// Access key1 to make it newest
	cache.Get("key1")

	if key, val, ok := cache.Newest(); !ok || key != "key1" || val != 1 {
		t.Errorf("Expected newest (key1, 1, true), got (%s, %d, %v)", key, val, ok)
	}

	if key, val, ok := cache.Oldest(); !ok || key != "key2" || val != 2 {
		t.Errorf("Expected oldest (key2, 2, true), got (%s, %d, %v)", key, val, ok)
	}
}

func TestResize(t *testing.T) {
	cache, _ := New[string, int](3)
	cache.Put("key1", 1)
	cache.Put("key2", 2)
	cache.Put("key3", 3)

	// Resize to smaller capacity
	err := cache.Resize(2)
	if err != nil {
		t.Errorf("Resize should not return error, got %v", err)
	}

	if cache.Cap() != 2 {
		t.Errorf("Expected capacity 2, got %d", cache.Cap())
	}

	if cache.Len() != 2 {
		t.Errorf("Expected length 2 after resize, got %d", cache.Len())
	}

	// key1 should be evicted (oldest)
	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been evicted during resize")
	}

	// Test invalid resize
	err = cache.Resize(0)
	if err == nil {
		t.Error("Resize to 0 should return error")
	}
}

func TestForEach(t *testing.T) {
	cache, _ := New[string, int](3)
	cache.Put("key1", 1)
	cache.Put("key2", 2)
	cache.Put("key3", 3)

	// Access key1 to change order
	cache.Get("key1")

	var keys []string
	var values []int

	cache.ForEach(func(key string, value int) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	expectedKeys := []string{"key1", "key3", "key2"}
	expectedValues := []int{1, 3, 2}

	for i, key := range expectedKeys {
		if keys[i] != key {
			t.Errorf("Expected key[%d] = %s, got %s", i, key, keys[i])
		}
	}

	for i, value := range expectedValues {
		if values[i] != value {
			t.Errorf("Expected value[%d] = %d, got %d", i, value, values[i])
		}
	}

	// Test early termination
	count := 0
	cache.ForEach(func(key string, value int) bool {
		count++
		return count < 2 // Stop after 2 iterations
	})

	if count != 2 {
		t.Errorf("Expected 2 iterations, got %d", count)
	}
}

func TestString(t *testing.T) {
	cache, _ := New[string, int](3)

	// Test empty cache
	if str := cache.String(); str != "LRUCache{}" {
		t.Errorf("Expected 'LRUCache{}', got '%s'", str)
	}

	cache.Put("key1", 1)
	cache.Put("key2", 2)

	str := cache.String()
	expected := "LRUCache{key2:2, key1:1}"
	if str != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str)
	}
}

func TestDifferentTypes(t *testing.T) {
	t.Run("Integer keys and string values", func(t *testing.T) {
		cache, _ := New[int, string](2)
		cache.Put(1, "one")
		cache.Put(2, "two")

		if val, ok := cache.Get(1); !ok || val != "one" {
			t.Errorf("Expected ('one', true), got ('%s', %v)", val, ok)
		}
	})

	t.Run("Custom struct types", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		cache, _ := New[string, Person](2)
		p1 := Person{"Alice", 30}
		p2 := Person{"Bob", 25}

		cache.Put("person1", p1)
		cache.Put("person2", p2)

		if val, ok := cache.Get("person1"); !ok || val.Name != "Alice" {
			t.Errorf("Expected Alice, got %s", val.Name)
		}
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("Single capacity cache", func(t *testing.T) {
		cache, _ := New[string, int](1)
		cache.Put("key1", 1)
		cache.Put("key2", 2) // Should evict key1

		if _, ok := cache.Get("key1"); ok {
			t.Error("key1 should have been evicted")
		}

		if val, ok := cache.Get("key2"); !ok || val != 2 {
			t.Errorf("Expected (2, true), got (%d, %v)", val, ok)
		}
	})

	t.Run("Rapid put/get operations", func(t *testing.T) {
		cache, _ := New[string, int](100)

		// Add many items
		for i := 0; i < 200; i++ {
			cache.Put(fmt.Sprintf("key%d", i), i)
		}

		// Cache should only have last 100 items
		if cache.Len() != 100 {
			t.Errorf("Expected length 100, got %d", cache.Len())
		}

		// First 100 items should be evicted
		for i := 0; i < 100; i++ {
			if _, ok := cache.Get(fmt.Sprintf("key%d", i)); ok {
				t.Errorf("key%d should have been evicted", i)
			}
		}

		// Last 100 items should be present
		for i := 100; i < 200; i++ {
			if val, ok := cache.Get(fmt.Sprintf("key%d", i)); !ok || val != i {
				t.Errorf("Expected (%d, true), got (%d, %v)", i, val, ok)
			}
		}
	})
}

func TestConcurrentAccess(t *testing.T) {
	// Note: This is a basic test. For true concurrent safety,
	// you would need to add synchronization to the LRU cache
	cache, _ := New[string, int](10)

	// Fill cache
	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	// Simulate concurrent access patterns
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i%10)
		cache.Get(key)
		cache.Put(key, i)
	}

	if cache.Len() != 10 {
		t.Errorf("Expected length 10, got %d", cache.Len())
	}
}

// Benchmark tests
func BenchmarkLRUCache(b *testing.B) {
	cache, _ := New[string, int](1000)

	b.Run("Put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Put(strconv.Itoa(i), i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		// Pre-populate cache
		for i := 0; i < 1000; i++ {
			cache.Put(strconv.Itoa(i), i)
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			cache.Get(strconv.Itoa(i % 1000))
		}
	})

	b.Run("Mixed operations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				cache.Put(strconv.Itoa(i), i)
			} else {
				cache.Get(strconv.Itoa(i / 2))
			}
		}
	})
}

func BenchmarkLRUCacheEviction(b *testing.B) {
	cache, _ := New[string, int](100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(strconv.Itoa(i), i)
	}
}
