package set1

import (
	"bufio"
	"cryptopals/config"
	"cryptopals/pkg/analysis"
	"cryptopals/pkg/cipher"
	"encoding/hex"
	"math"
	"os"
	"testing"
)

const (
	CORPUS_4 string = "corpus_4.txt"
)

// setupFileArray is a test helper that reads all the 60-character hexadecimal strings from
// corpus_4.txt in testdata and returns a slice containing their conversion to bytes.
func setupFileArray(t *testing.T) [][]byte {
	t.Helper()

	file, err := os.Open(config.DataPath(CORPUS_4))
	if err != nil {
		t.Fatalf("failed to open corpus_4: %v", err)
	}
	defer file.Close()

	messages := make([][]byte, 0, 400)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		msg, err := hex.DecodeString(scanner.Text())
		if err != nil {
			t.Fatalf("failed to decode hex input: %v", err)
		}
		messages = append(messages, msg)
	}
	if scanner.Err() != nil {
		t.Fatalf("failed to scan through corpus: %v", scanner.Err())
	}
	return messages
}

func TestChallenge04_DetectByteXOR(t *testing.T) {
	// 1. Setup enviroment
	messages := setupFileArray(t)
	freqMap := setupFrequencyMap(t)

	// 2. Define the scoring closure to inject into the decrypter
	scorer := func(candidate []byte) float64 {
		return analysis.Score(candidate, freqMap)
	}

	// 3. Analyze each message and keep the best one
	bestScore := math.Inf(1)
	var decryptedText, encryptedBytes []byte
	for _, msg := range messages {
		plaintext, _, score, err := cipher.BreakSingleByteXor(msg, scorer)
		if err != nil {
			t.Fatalf("unexpected error breaking single-byte XOR: %v", err)
		}
		if score < bestScore {
			decryptedText = plaintext
			encryptedBytes = msg
			bestScore = score
		}
	}

	// 4. Return the candidate to original hexadecimal shape and print the results
	t.Logf("The single-byte encripted hexadecimal string is: %s, from which we got the message: %s", hex.EncodeToString(encryptedBytes), decryptedText)
}
