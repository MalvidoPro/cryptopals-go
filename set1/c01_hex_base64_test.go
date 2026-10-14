package set1

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestChallenge01_HexToBase64(t *testing.T) {
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// 1. Decode hexadecimal input to raw bytes ([]byte)
	rawBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		t.Fatalf("error inesperado decodificando hex: %v", err)
	}

	// 2. Encode those raw bytes to base64
	resultBase64 := base64.StdEncoding.EncodeToString(rawBytes)

	// 3. Test
	if resultBase64 != expectedBase64 {
		t.Errorf("fallo en la conversión:\nObtenido: %s\nEsperado: %s", resultBase64, expectedBase64)
	}
}
