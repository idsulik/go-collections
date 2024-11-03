package bloomfilter

import (
	"fmt"
	"testing"
)

func TestBloomFilter_Basic(t *testing.T) {
	tests := []struct {
		name            string
		expectedItems   uint
		falsePositive   float64
		itemsToAdd      []string
		itemsToCheck    []string
		shouldContain   []bool
		expectedMinBits uint
	}{
		{
			name:          "Basic operation",
			expectedItems: 100,
			falsePositive: 0.01,
			itemsToAdd:    []string{"apple", "banana", "cherry"},
			itemsToCheck:  []string{"apple", "banana", "cherry", "date"},
			shouldContain: []bool{true, true, true, false},
		},
		{
			name:          "Empty filter",
			expectedItems: 100,
			falsePositive: 0.01,
			itemsToAdd:    []string{},
			itemsToCheck:  []string{"apple"},
			shouldContain: []bool{false},
		},
		{
			name:          "Single item",
			expectedItems: 100,
			falsePositive: 0.01,
			itemsToAdd:    []string{"apple"},
			itemsToCheck:  []string{"apple", "banana"},
			shouldContain: []bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bf := NewBloomFilter[string](tt.expectedItems, tt.falsePositive)

				// Add items
				for _, item := range tt.itemsToAdd {
					bf.Add(item)
				}

				// Check size matches expected items
				if bf.Len() != len(tt.itemsToAdd) {
					t.Errorf("Expected length %d, got %d", len(tt.itemsToAdd), bf.Len())
				}

				// Check contains
				for i, item := range tt.itemsToCheck {
					if bf.Contains(item) != tt.shouldContain[i] {
						t.Errorf(
							"Contains(%s) = %v, want %v",
							item, bf.Contains(item), tt.shouldContain[i],
						)
					}
				}
			},
		)
	}
}

func TestBloomFilter_DifferentTypes(t *testing.T) {
	t.Run(
		"Integer type", func(t *testing.T) {
			bf := NewBloomFilter[int](100, 0.01)
			numbers := []int{1, 2, 3, 4, 5}

			for _, n := range numbers {
				bf.Add(n)
			}

			for _, n := range numbers {
				if !bf.Contains(n) {
					t.Errorf("Should contain %d", n)
				}
			}

			if bf.Contains(6) {
				t.Error("Should not contain 6")
			}
		},
	)

	t.Run(
		"Custom struct type", func(t *testing.T) {
			type Person struct {
				Name string
				Age  int
			}

			bf := NewBloomFilter[Person](100, 0.01)
			p1 := Person{"Alice", 30}
			p2 := Person{"Bob", 25}

			bf.Add(p1)
			bf.Add(p2)

			if !bf.Contains(p1) {
				t.Error("Should contain person 1")
			}
			if !bf.Contains(p2) {
				t.Error("Should contain person 2")
			}
			if bf.Contains(Person{"Charlie", 35}) {
				t.Error("Should not contain person 3")
			}
		},
	)
}

func TestBloomFilter_EdgeCases(t *testing.T) {
	t.Run(
		"Zero expected items", func(t *testing.T) {
			bf := NewBloomFilter[string](0, 0.01)
			if bf == nil {
				t.Error("Should create filter even with zero expected items")
			}
			bf.Add("test")
			if !bf.Contains("test") {
				t.Error("Should still function with zero expected items")
			}
		},
	)

	t.Run(
		"Zero false positive rate", func(t *testing.T) {
			bf := NewBloomFilter[string](100, 0)
			if bf == nil {
				t.Error("Should create filter even with zero false positive rate")
			}
			bf.Add("test")
			if !bf.Contains("test") {
				t.Error("Should still function with zero false positive rate")
			}
		},
	)
}

func TestBloomFilter_Operations(t *testing.T) {
	t.Run(
		"Clear operation", func(t *testing.T) {
			bf := NewBloomFilter[string](100, 0.01)
			bf.Add("test")

			if !bf.Contains("test") {
				t.Error("Should contain 'test' before clear")
			}

			bf.Clear()

			if bf.Contains("test") {
				t.Error("Should not contain 'test' after clear")
			}

			if !bf.IsEmpty() {
				t.Error("Should be empty after clear")
			}

			if bf.Len() != 0 {
				t.Error("Length should be 0 after clear")
			}
		},
	)
}

func TestBloomFilter_FalsePositiveRate(t *testing.T) {
	expectedItems := uint(1000)
	targetFPR := 0.01
	bf := NewBloomFilter[int](expectedItems, targetFPR)

	// Add expectedItems number of items
	for i := 0; i < int(expectedItems); i++ {
		bf.Add(i)
	}

	// Test false positive rate
	falsePositives := 0
	trials := 10000
	for i := int(expectedItems); i < int(expectedItems)+trials; i++ {
		if bf.Contains(i) {
			falsePositives++
		}
	}

	actualFPR := float64(falsePositives) / float64(trials)
	estimatedFPR := bf.EstimatedFalsePositiveRate()

	// Allow for some variance in the actual false positive rate
	maxAcceptableFPR := targetFPR * 2
	if actualFPR > maxAcceptableFPR {
		t.Errorf(
			"False positive rate too high: got %f, want <= %f",
			actualFPR, maxAcceptableFPR,
		)
	}

	// Check if estimated FPR is reasonably close to actual FPR
	if estimatedFPR < actualFPR/2 || estimatedFPR > actualFPR*2 {
		t.Errorf(
			"Estimated FPR %f significantly different from actual FPR %f",
			estimatedFPR, actualFPR,
		)
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	bf := NewBloomFilter[string](1000, 0.01)

	b.Run(
		"Add", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bf.Add(fmt.Sprintf("item%d", i))
			}
		},
	)

	b.Run(
		"Contains", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bf.Contains(fmt.Sprintf("item%d", i))
			}
		},
	)
}
