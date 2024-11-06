package segmenttree

import (
	"math"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	// Test case 1: Range Sum
	t.Run(
		"Range Sum", func(t *testing.T) {
			arr := []int{1, 3, 5, 7, 9, 11}
			st := NewSegmentTree(arr, 0, func(a, b int) int { return a + b })

			tests := []struct {
				left, right int
				want        int
			}{
				{0, 2, 9},  // 1 + 3 + 5
				{1, 4, 24}, // 3 + 5 + 7 + 9
				{0, 5, 36}, // sum of all elements
				{3, 3, 7},  // single element
			}

			for _, tt := range tests {
				got := st.Query(tt.left, tt.right)
				if got != tt.want {
					t.Errorf("Query(%d, %d) = %d; want %d", tt.left, tt.right, got, tt.want)
				}
			}

			// Test update
			st.Update(2, 6) // Change 5 to 6
			if got := st.Query(0, 2); got != 10 {
				t.Errorf("After update, Query(0, 2) = %d; want 10", got)
			}
		},
	)

	// Test case 2: Range Minimum
	t.Run(
		"Range Minimum", func(t *testing.T) {
			arr := []float64{3.5, 1.2, 4.8, 2.3, 5.4}
			st := NewSegmentTree(arr, math.Inf(1), func(a, b float64) float64 { return math.Min(a, b) })

			tests := []struct {
				left, right int
				want        float64
			}{
				{0, 2, 1.2},
				{2, 4, 2.3},
				{0, 4, 1.2},
			}

			for _, tt := range tests {
				got := st.Query(tt.left, tt.right)
				if got != tt.want {
					t.Errorf("Query(%d, %d) = %f; want %f", tt.left, tt.right, got, tt.want)
				}
			}
		},
	)
}

func BenchmarkSegmentTree(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	st := NewSegmentTree(arr, 0, func(a, b int) int { return a + b })

	b.Run(
		"Query", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				st.Query(100, 9900)
			}
		},
	)

	b.Run(
		"Update", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				st.Update(5000, i)
			}
		},
	)
}
