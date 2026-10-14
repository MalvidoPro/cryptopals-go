package set1

import (
	"bytes"
	"cryptopals/config"
	"cryptopals/pkg/cipher"
	"encoding/hex"
	"os"
	"testing"
)

const (
	TEXT_5     string = "text_5.txt"
	SOL_5      string = "personal_sol_5.txt"
	PALS_SOL_5 string = "sol_5.txt"
	KEY        string = "ICE"
)

func TestChallenge05_repeatingKeyXOR(t *testing.T) {
	text, err := os.ReadFile(config.DataPath(TEXT_5))
	if err != nil {
		t.Fatalf("error reading source text: %v", err)
	}
	text = bytes.ReplaceAll(text, []byte("\r\n"), []byte("\n"))

	// 1. Encrypt the text and put it in the solution file (in its hexadecimal form).
	ciphertextHexByte := []byte(hex.EncodeToString(cipher.RepeatingKey(KEY, text)))
	err = config.WriteResult(SOL_5, ciphertextHexByte)
	if err != nil {
		t.Fatalf("error writing solution: %v", err)
	}

	// 2. Verify that the text is what we expected, comparing the file
	// containing cryptopals' solution and our solution.

	solution, err := os.ReadFile(config.DataPath(PALS_SOL_5))
	if err != nil {
		t.Fatalf("error reading cryptopals' solution: %v", err)
	}
	if !bytes.Equal(solution, ciphertextHexByte) {
		t.Errorf("incorrect ciphertext.\nGot:      %s\nExpected: %s\n", ciphertextHexByte, solution)
	}
}
