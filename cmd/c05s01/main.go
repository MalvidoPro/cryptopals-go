package main

import (
	"bytes"
	"cryptopals/pkg/encoding"
	"encoding/hex"
	"log"
	"os"
)

const (
	TEXT_FILE string = "../testdata/umu.txt"
	SOL_FILE  string = "../results/umuEncrypted.txt"
	KEY       string = "ICE"
)

func main() {
	text, err := os.ReadFile(TEXT_FILE)
	if err != nil {
		log.Fatalf("error reading source text: %v", err)
	}
	text = bytes.ReplaceAll(text, []byte("\r\n"), []byte("\n"))

	// 2. Encrypt the text and put it in the solution file (in its hexadecimal form).
	ciphertextHexByte := []byte(hex.EncodeToString(encoding.RepeatingKey(KEY, text)))
	os.WriteFile(SOL_FILE, ciphertextHexByte, 0644)

}
