package ringbuffer

import (
	"testing"
)

func TestRingBuffer(t *testing.T) {
	t.Run(
		"New buffer creation", func(t *testing.T) {
			rb := New[int](5)
			if rb.Cap() != 5 {
				t.Errorf("Expected capacity of 5, got %d", rb.Cap())
			}
			if !rb.IsEmpty() {
				t.Error("New buffer should be empty")
			}
		},
	)

	t.Run(
		"Writing and reading", func(t *testing.T) {
			rb := New[int](3)

			// Write test
			if !rb.Write(1) {
				t.Error("Write should succeed on empty buffer")
			}
			if rb.Len() != 1 {
				t.Errorf("Expected length 1, got %d", rb.Len())
			}

			// Read test
			val, ok := rb.Read()
			if !ok || val != 1 {
				t.Errorf("Expected to read 1, got %d", val)
			}
			if !rb.IsEmpty() {
				t.Error("Buffer should be empty after reading")
			}
		},
	)

	t.Run(
		"Buffer full behavior", func(t *testing.T) {
			rb := New[int](2)

			rb.Write(1)
			rb.Write(2)

			if !rb.IsFull() {
				t.Error("Buffer should be full")
			}

			if rb.Write(3) {
				t.Error("Write should fail when buffer is full")
			}
		},
	)

	t.Run(
		"Peek operation", func(t *testing.T) {
			rb := New[int](2)
			rb.Write(1)

			val, ok := rb.Peek()
			if !ok || val != 1 {
				t.Errorf("Expected to peek 1, got %d", val)
			}
			if rb.Len() != 1 {
				t.Error("Peek should not remove items")
			}
		},
	)

	t.Run(
		"Clear operation", func(t *testing.T) {
			rb := New[int](2)
			rb.Write(1)
			rb.Write(2)

			rb.Clear()
			if !rb.IsEmpty() {
				t.Error("Buffer should be empty after clear")
			}
			if rb.Len() != 0 {
				t.Errorf("Expected length 0 after clear, got %d", rb.Len())
			}
		},
	)

	t.Run(
		"Circular behavior", func(t *testing.T) {
			rb := New[int](3)

			// Fill the buffer
			rb.Write(1)
			rb.Write(2)
			rb.Write(3)

			// Read two items
			rb.Read()
			rb.Read()

			// Write two more
			rb.Write(4)
			rb.Write(5)

			// Check the sequence
			val1, _ := rb.Read()
			val2, _ := rb.Read()
			val3, _ := rb.Read()

			if val1 != 3 || val2 != 4 || val3 != 5 {
				t.Error("Circular buffer not maintaining correct order")
			}
		},
	)
}
