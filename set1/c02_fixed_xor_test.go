package set1

import (
	"bytes"
	"cryptopals/pkg/cipher" // Asegúrate de cambiar "tu_modulo" por el nombre real de tu go.mod
	"encoding/hex"
	"testing"
)

func TestChallenge02_FixedXOR(t *testing.T) {
	input1Hex := "1c0111001f010100061a024b53535009181c"
	input2Hex := "686974207468652062756c6c277320657965"
	expectedHex := "746865206b696420646f6e277420706c6179"

	// 1. Decode inputs
	input1, err := hex.DecodeString(input1Hex)
	if err != nil {
		t.Fatalf("failed to decode input 1: %v", err)
	}

	input2, err := hex.DecodeString(input2Hex)
	if err != nil {
		t.Fatalf("failed to decode input 2: %v", err)
	}

	expected, err := hex.DecodeString(expectedHex)
	if err != nil {
		t.Fatalf("failed to decode expected output: %v", err)
	}

	// 2. Perform raw byte-level cryptographic operation
	result, err := cipher.XorBytes(input1, input2)
	if err != nil {
		t.Fatalf("unexpected error in XorBytes: %v", err)
	}

	// 3. Raw memory assertion
	if !bytes.Equal(result, expected) {
		t.Errorf("XOR mismatch:\ngot:  %x\nwant: %x", result, expected)
	}
}
