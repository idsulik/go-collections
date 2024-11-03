package bloomfilter

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter[T any] struct {
	bits    []bool
	numBits uint
	numHash uint
	count   uint
	hasher  hash.Hash
}

func NewBloomFilter[T any](expectedItems uint, falsePositiveProb float64) *BloomFilter[T] {
	if expectedItems == 0 {
		expectedItems = 1
	}
	if falsePositiveProb <= 0 {
		falsePositiveProb = 0.01
	}

	numBits := uint(math.Ceil(-float64(expectedItems) * math.Log(falsePositiveProb) / math.Pow(math.Log(2), 2)))
	numHash := uint(math.Ceil(float64(numBits) / float64(expectedItems) * math.Log(2)))

	return &BloomFilter[T]{
		bits:    make([]bool, numBits),
		numBits: numBits,
		numHash: numHash,
		hasher:  fnv.New64a(), // Using fnv.New64a() for better distribution
	}
}

// hashToUint converts a hash sum to uint
func hashToUint(sum []byte) uint {
	// Convert the first 8 bytes of the hash to uint64, then to uint
	if len(sum) < 8 {
		panic("Hash sum too short")
	}
	return uint(binary.BigEndian.Uint64(sum[:8]))
}

// getLocations generates multiple hash locations for an item
func (bf *BloomFilter[T]) getLocations(item T) []uint {
	locations := make([]uint, bf.numHash)
	itemStr := fmt.Sprintf("%v", item)

	// Calculate SHA-256 hash as a base for location generation
	hash := sha256.Sum256([]byte(itemStr))
	h1 := hashToUint(hash[:8])   // Use first 8 bytes for h1
	h2 := hashToUint(hash[8:16]) // Use next 8 bytes for h2

	// Generate all hash values using the formula: h1 + i*h2
	for i := uint(0); i < bf.numHash; i++ {
		locations[i] = (h1 + i*h2) % bf.numBits
	}

	return locations
}

// Add inserts an item into the Bloom Filter.
func (bf *BloomFilter[T]) Add(item T) {
	locations := bf.getLocations(item)
	for _, loc := range locations {
		bf.bits[loc] = true
	}
	bf.count++
}

// Contains tests whether an item might be in the set.
func (bf *BloomFilter[T]) Contains(item T) bool {
	locations := bf.getLocations(item)
	for _, loc := range locations {
		if !bf.bits[loc] {
			return false
		}
	}
	return true
}

// EstimatedFalsePositiveRate returns the estimated false positive rate.
func (bf *BloomFilter[T]) EstimatedFalsePositiveRate() float64 {
	if bf.count == 0 {
		return 0.0
	}
	exponent := -float64(bf.numHash) * float64(bf.count) / float64(bf.numBits)
	return math.Pow(1-math.Exp(exponent), float64(bf.numHash))
}

// Clear removes all items from the Bloom Filter.
func (bf *BloomFilter[T]) Clear() {
	bf.bits = make([]bool, bf.numBits)
	bf.count = 0
}

// Len returns the number of items added to the Bloom Filter.
func (bf *BloomFilter[T]) Len() int {
	return int(bf.count)
}

// IsEmpty returns true if no items have been added to the Bloom Filter.
func (bf *BloomFilter[T]) IsEmpty() bool {
	return bf.count == 0
}

// BitSize returns the size of the bit array.
func (bf *BloomFilter[T]) BitSize() uint {
	return bf.numBits
}

// NumberOfHashes returns the number of hash functions.
func (bf *BloomFilter[T]) NumberOfHashes() uint {
	return bf.numHash
}
