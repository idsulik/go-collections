package stack

// Stack is an interface representing a LIFO (last-in, first-out) stack.
type Stack[T any] interface {
	// Push adds an item to the top of the stack.
	Push(item T)

	// Pop removes and returns the item from the top of the stack.
	// Returns false if the stack is empty.
	Pop() (T, bool)

	// Peek returns the item at the top of the stack without removing it.
	// Returns false if the stack is empty.
	Peek() (T, bool)

	// Len returns the number of items currently in the stack.
	Len() int

	// IsEmpty checks if the stack is empty.
	IsEmpty() bool

	// Clear removes all items from the stack, leaving it empty.
	Clear()
}
