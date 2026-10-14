package analysis

import "fmt"

func numberOfOnes(b byte) int {
	var sum int
	for range 7 {
		if b%2 == 1 {
			sum++
		}
		b = b >> 1
	}
	return sum
}

func Hamming(b1 []byte, b2 []byte) (int, error) {
	l := len(b1)
	if l != len(b2) {
		return 0, fmt.Errorf("cannot compute hamming: strings do not have the same length\n")
	}
	var hamDistance int
	for i := range l {
		hamDistance += numberOfOnes(b1[i] ^ b2[i])
	}
	return hamDistance, nil
}
