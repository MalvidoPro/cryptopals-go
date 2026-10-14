package cipher

import "fmt"

// XorBytes is a function that takes two byte buffers and returns the result of doing a XOR
// operation between them.
func XorBytes(buf1, buf2 []byte) ([]byte, error) {
	// 1. Verify size
	size := len(buf1)
	if size != len(buf2) {
		return nil, fmt.Errorf("error: byte slices do not have the same size")
	}

	// 2. Do the XOR operation and return
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = buf1[i] ^ buf2[i]
	}
	return result, nil
}
