package cipher

import (
	"cmp"
	"cryptopals/pkg/analysis"
	"fmt"
	"slices"
)

type KeyResult struct {
	Size     int
	Distance float64
}

// FindKeySize determines the most probable key sizes for a repeating-key XOR cipher.
// It computes the normalized Hamming distance between multiple blocks of the ciphertext.
// Returns a slice of the best keysizes (nKeys will be the amount of results returned).
func FindKeySize(corpus []byte, max int, min int, nKeys int) ([]KeyResult, error) {
	// Slice of KeyResult in which we will input the score for every keysize
	results := make([]KeyResult, 0, max-min+1)
	for i := min; i <= max; i++ {
		// Prevent out-of-bounds panics if the corpus is too short for the current key size
		if len(corpus) < 4*i {
			break
		}

		// Extract 4 consecutive blocks of size i
		b0 := corpus[:i]
		b1 := corpus[i : 2*i]
		b2 := corpus[2*i : 3*i]
		b3 := corpus[3*i : 4*i]

		// Define the pairs to compare: (0,1), (2,3), (0,2), (1,3)
		pairs := [][2][]byte{
			{b0, b1},
			{b2, b3},
			{b0, b2},
			{b1, b3},
		}

		var totalHamming int

		// Compute Hamming distances for all defined pairs
		for _, pair := range pairs {
			dist, err := analysis.Hamming(pair[0], pair[1])
			if err != nil {
				return nil, fmt.Errorf("error computing hamming distance for keysize %d: %w", i, err)
			}
			totalHamming += dist
		}

		// Normalize the distance: average distance across pairs, divided by the key size
		normalizedScore := float64(totalHamming) / (4.0 * float64(i))
		results = append(results, KeyResult{Size: i, Distance: normalizedScore})
	}
	// Sort the keysizes from the smallest to the highest hamming distance
	slices.SortFunc(results, func(a, b KeyResult) int {
		return cmp.Compare(a.Distance, b.Distance)
	})
	return results[:nKeys], nil
}
