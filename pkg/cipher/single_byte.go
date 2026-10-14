package cipher

import (
	"fmt"
	"math"
)

// createKey creates a sequence of a certain byte. The byte and the length of
// the sequence are specified as parameters.
func createKey(b byte, length int) []byte {
	key := make([]byte, length)
	for i := range length {
		key[i] = b
	}
	return key
}

// BreakSingleByteXor attempts to decrypt a single-byte XOR cipher.
// It tests all 256 possible byte keys and uses the provided scorer function
// to evaluate each candidate. It returns the best plaintext candidate,
// the key used to decrypt it, the score of that key using the scorer function and
// any potential error. It relies on the fact that XOR decryption uses the same
// key as XOR encryption.
func BreakSingleByteXor(ciphertext []byte, scorer func([]byte) float64) ([]byte, byte, float64, error) {
	// Initialize variables for the comparison
	length := len(ciphertext)
	bestSol := math.Inf(1)
	var sol float64 = 0
	var err error
	var keyApplied, keyAppliedSol []byte
	var key byte
	for i := range 256 {
		// 1. Decipher the text using the key
		keyApplied, err = XorBytes(ciphertext, createKey(byte(i), length))
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to apply XOR to ciphertext: %w", err)
		}

		// 2. Compute the score for this canidadate text
		sol = scorer(keyApplied)
		// 3. Store current solution if it is the best right now
		if sol < bestSol {
			bestSol = sol
			key = byte(i)
			keyAppliedSol = keyApplied
		}
	}
	return keyAppliedSol, key, bestSol, nil
}
