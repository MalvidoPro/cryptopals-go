package set1

import (
	"bytes"
	"cryptopals/config"
	"cryptopals/pkg/analysis"
	"cryptopals/pkg/cipher"
	"encoding/csv"
	"encoding/hex"
	"os"
	"strconv"
	"testing"
)

const (
	LETTER_FREQ string = "letter_frequency.csv"
)

// setupFrequencyMap is a test helper that reads the frequency CSV from testdata.
// It returns a map binding each character to its expected frequency.
func setupFrequencyMap(t *testing.T) map[byte]float64 {
	t.Helper()

	file, err := os.Open(config.DataPath(LETTER_FREQ))
	if err != nil {
		t.Fatalf("failed to open frequency data: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	freqDict, err := csvReader.ReadAll()
	if err != nil {
		t.Fatalf("failed to read the frequency file: %v", err)
	}

	freqEng := make(map[byte]float64, 26)
	for i := 0; i < 26; i++ {
		num, err := strconv.ParseFloat(freqDict[i+1][1], 64)
		if err != nil {
			t.Fatalf("error parsing float %d: %v", i+1, err)
		}
		freqEng[byte('a'+i)] = num
	}
	return freqEng
}

func TestChallenge03_SingleByteXOR(t *testing.T) {
	inputHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedPlaintext := []byte("Cooking MC's like a pound of bacon")

	// 1. Setup environment
	ciphertext, err := hex.DecodeString(inputHex)
	if err != nil {
		t.Fatalf("failed to decode hex input: %v", err)
	}

	freqMap := setupFrequencyMap(t)

	// 2. Define the scoring closure to inject into the decrypter
	scorer := func(candidate []byte) float64 {
		return analysis.Score(candidate, freqMap)
	}

	// 3. Execute the attack
	plaintext, key, _, err := cipher.BreakSingleByteXor(ciphertext, scorer)
	if err != nil {
		t.Fatalf("unexpected error breaking single-byte XOR: %v", err)
	}

	// 4. Verification
	if !bytes.Equal(plaintext, expectedPlaintext) {
		t.Errorf("incorrect plaintext.\nGot: %s\nExpected: %s", plaintext, expectedPlaintext)
	}

	// Optional: log the discovered key for visual verification when running with -v
	t.Logf("Discovered key: %c (0x%02x)", key, key)
}
