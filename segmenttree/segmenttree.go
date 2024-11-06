package segmenttree

// Operation represents a function type for segment tree operations
type Operation[T any] func(T, T) T

// SegmentTree provides efficient range query operations
type SegmentTree[T any] struct {
	tree     []T
	size     int
	identity T
	combine  Operation[T]
}

// NewSegmentTree creates a new segment tree from given array
func NewSegmentTree[T any](arr []T, identity T, combine Operation[T]) *SegmentTree[T] {
	n := len(arr)
	tree := make([]T, 4*n) // 4*n is enough to store the segment tree
	st := &SegmentTree[T]{
		tree:     tree,
		size:     n,
		identity: identity,
		combine:  combine,
	}
	st.build(arr, 0, 0, n-1)
	return st
}

// build constructs the segment tree
func (st *SegmentTree[T]) build(arr []T, node int, start, end int) T {
	if start == end {
		st.tree[node] = arr[start]
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftVal := st.build(arr, 2*node+1, start, mid)
	rightVal := st.build(arr, 2*node+2, mid+1, end)
	st.tree[node] = st.combine(leftVal, rightVal)
	return st.tree[node]
}

// Update updates a value at given index
func (st *SegmentTree[T]) Update(index int, value T) {
	st.updateRecursive(0, 0, st.size-1, index, value)
}

func (st *SegmentTree[T]) updateRecursive(node int, start, end, index int, value T) {
	if start == end {
		st.tree[node] = value
		return
	}

	mid := (start + end) / 2
	if index <= mid {
		st.updateRecursive(2*node+1, start, mid, index, value)
	} else {
		st.updateRecursive(2*node+2, mid+1, end, index, value)
	}
	st.tree[node] = st.combine(st.tree[2*node+1], st.tree[2*node+2])
}

// Query returns the result of the operation in range [left, right]
func (st *SegmentTree[T]) Query(left, right int) T {
	return st.queryRecursive(0, 0, st.size-1, left, right)
}

func (st *SegmentTree[T]) queryRecursive(node int, start, end, left, right int) T {
	if right < start || left > end {
		return st.identity
	}
	if left <= start && end <= right {
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftVal := st.queryRecursive(2*node+1, start, mid, left, right)
	rightVal := st.queryRecursive(2*node+2, mid+1, end, left, right)
	return st.combine(leftVal, rightVal)
}
