package set1

import (
	"cryptopals/config"
	"cryptopals/pkg/analysis"
	"cryptopals/pkg/cipher"
	"encoding/base64"
	"os"
	"testing"
)

const (
	MAX_KEYSIZE int    = 40
	MIN_KEYSIZE int    = 2
	N_KEYS      int    = 3
	CORPUS_6    string = "corpus_6.txt"
	RESULTS_6_1 string = "llave_1.txt"
	RESULTS_6_2 string = "llave_2.txt"
	RESULTS_6_3 string = "llave_3.txt"
)

// divideAndTranspose takes a byte slice and a chunk size n, splitting the slice
// into blocks of size n and transposing them. The resulting slice of byte slices
// will contain n slices, where the i-th slice contains the i-th byte of every block.
func divideAndTranspose(b []byte, n int) [][]byte {
	if len(b) < n {
		return [][]byte{b}
	}
	r := len(b) % n
	l := len(b) / n

	chunks := make([][]byte, n)
	for i := range n {
		if i < r {
			chunks[i] = make([]byte, l+1)
		} else {
			chunks[i] = make([]byte, l)
		}
		for j := range l {
			chunks[i][j] = b[n*j+i]
		}
	}
	for i := range r {
		chunks[i][l] = b[n*l+i]
	}
	return chunks
}

// TestBreakRepeatingKeyText is the main test function that reads a base64 encoded
// file, deduces the key size, transposes the blocks, breaks the single-byte XOR
// cipher for each block, and writes the possible decrypted results to files.
func TestChallenge06_BreakRepeatingKey(t *testing.T) {
	str1 := "this is a test"
	str2 := "wokka wokka!!!"

	n, err := analysis.Hamming([]byte(str1), []byte(str2))
	if err != nil {
		t.Fatalf("error computing hamming distance: %v", err)
	}
	t.Logf("Hamming distance is %v\n", n)

	rawCorpus, err := os.ReadFile(config.DataPath(CORPUS_6))
	if err != nil {
		t.Fatalf("error reading cryptopals' solution: %v", err)
	}

	// The corpus has been base64'd
	corpus := make([]byte, len(rawCorpus)*3/4+1)
	nDecoded, err := base64.StdEncoding.Decode(corpus, rawCorpus)
	if err != nil {
		t.Fatalf("error decoding base64'd corpus: %v", err)
	}
	corpus = corpus[:nDecoded]

	keysizes, err := cipher.FindKeySize(corpus, MAX_KEYSIZE, MIN_KEYSIZE, N_KEYS)
	if err != nil {
		t.Fatalf("error finding key sizes: %v", err)
	}
	t.Logf("The three best sizes found are: %v, %v and %v\n", keysizes[0].Size, keysizes[1].Size, keysizes[2].Size)

	possibleKeys := make([][]byte, 3)
	freqMap := setupFrequencyMap(t)
	scorer := func(candidate []byte) float64 {
		return analysis.Score(candidate, freqMap)
	}

	for i, keydata := range keysizes {
		dividedCorpus := divideAndTranspose(corpus, keydata.Size)
		possibleKeys[i] = make([]byte, keydata.Size)

		for j, text := range dividedCorpus {
			_, x, _, err := cipher.BreakSingleByteXor(text, scorer)
			if err != nil {
				t.Fatalf("error obtaining a part of the key using single byte xor: %v", err)
			}
			possibleKeys[i][j] = x
		}
	}
	err = config.WriteResult(RESULTS_6_1, cipher.RepeatingKey(string(possibleKeys[0]), corpus))
	if err != nil {
		t.Fatalf("error writing file llave_1.txt: %v", err)
	}

	err = config.WriteResult(RESULTS_6_2, cipher.RepeatingKey(string(possibleKeys[1]), corpus))
	if err != nil {
		t.Fatalf("error writing file llave_2.txt: %v", err)
	}

	err = config.WriteResult(RESULTS_6_3, cipher.RepeatingKey(string(possibleKeys[2]), corpus))
	if err != nil {
		t.Fatalf("error writing file llave_3.txt: %v", err)
	}
}
