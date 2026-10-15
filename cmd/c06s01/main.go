package main

import (
	"cryptopals/pkg/analysis"
	"cryptopals/pkg/cipher"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	MAX_KEYSIZE int = 40
	MIN_KEYSIZE int = 2
	N_KEYS      int = 3
)

// This function is on package set1, so when converting this to a test,
// it won't be necessary to put it here. Also, execute run command from
// root while using this function to avoid problems with this relative route.
func setupFrequencyMap() map[byte]float64 {
	file, err := os.Open("testdata/letter_frequency.csv")
	if err != nil {
		log.Fatalf("failed to open frequency data: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	freqDict, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read the frequency file: %v", err)
	}

	freqEng := make(map[byte]float64, 26)
	for i := 0; i < 26; i++ {
		num, err := strconv.ParseFloat(freqDict[i+1][1], 64)
		if err != nil {
			log.Fatalf("error parsing float %d: %v", i+1, err)
		}
		freqEng[byte('a'+i)] = num
	}
	return freqEng
}

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

func main() {
	str1 := "this is a test"

	str2 := "wokka wokka!!!"

	n, err := analysis.Hamming([]byte(str1), []byte(str2))
	if err != nil {
		log.Fatalf("error computing hamming distance: %v", err)
	}
	fmt.Printf("Hamming distance is %v\n", n)

	raw_corpus, err := os.ReadFile("testdata/corpus_6.txt")
	if err != nil {
		log.Fatalf("error reading cryptopals' solution: %v", err)
	}

	// I has been base64'd
	corpus := make([]byte, len(raw_corpus)*3/4+1)
	_, err = base64.StdEncoding.Decode(corpus, raw_corpus)
	if err != nil {
		log.Fatalf("error decoding base64'd corpus: %v", err)
	}

	keysizes, err := cipher.FindKeySize(corpus, MAX_KEYSIZE, MIN_KEYSIZE, N_KEYS)
	fmt.Printf("The three best sizes found are: %v, %v and %v\n", keysizes[0].Size, keysizes[1].Size, keysizes[2].Size)
	possibleKeys := make([][]byte, 3)

	freqMap := setupFrequencyMap()
	scorer := func(candidate []byte) float64 {
		return analysis.Score(candidate, freqMap)
	}
	for i, keydata := range keysizes {
		dividedCorpus := divideAndTranspose(corpus, keydata.Size)
		possibleKeys[i] = make([]byte, keydata.Size)
		fmt.Printf("First letter group is: %v\n", string(dividedCorpus[0]))

		for j, text := range dividedCorpus {
			_, x, _, err := cipher.BreakSingleByteXor(text, scorer)
			if err != nil {
				log.Fatalf("error obtaining a part of the key using single byte xor: %v", err)
			}
			possibleKeys[i][j] = x
		}
	}
	fmt.Printf("First key is: %v\n", possibleKeys[0])

	os.WriteFile("results/llave_1.txt", cipher.RepeatingKey(string(possibleKeys[0]), corpus), 0644)
	os.WriteFile("results/llave_2.txt", cipher.RepeatingKey(string(possibleKeys[1]), corpus), 0644)
	os.WriteFile("results/llave_3.txt", cipher.RepeatingKey(string(possibleKeys[2]), corpus), 0644)

}
