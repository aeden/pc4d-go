package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v plaintext\n\tplaintext: unencrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var nonce [24]byte
	var key [32]byte
	message := []byte(os.Args[1])

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	log.Printf("Nonce: %v", nonce)

	if _, err := io.ReadFull(rand.Reader, key[:]); err != nil {
		panic(err)
	}
	log.Printf("Key: %v", hex.EncodeToString(key[:]))

	// The nonce is prefixed to the encrypted data. The correct nonce must be used to decrypt the encrypted data
	box := secretbox.Seal(nonce[:], message, &nonce, &key)
	log.Printf("Box: %v", base64.StdEncoding.EncodeToString(box))
}
