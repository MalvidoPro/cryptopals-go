package cipher

// RepeatingKey applies a repeating-key XOR cipher to a single byte slice.
// It returns a newly allocated byte slice containing the encoded data.
func RepeatingKey(key string, text []byte) []byte {
	byteKey := []byte(key)
	lenKey := len(byteKey)
	n := len(text)
	encoded := make([]byte, n)

	// Basic safeguard in case an empty key is provided
	if lenKey == 0 {
		copy(encoded, text)
		return encoded
	}

	rest := n % lenKey
	for i := 0; i < n-rest; i += lenKey {
		for j := range lenKey {
			encoded[i+j] = text[i+j] ^ byteKey[j]
		}
	}

	for j := range rest {
		encoded[n-rest+j] = text[n-rest+j] ^ byteKey[j]
	}

	return encoded
}
