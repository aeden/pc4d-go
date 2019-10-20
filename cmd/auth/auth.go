package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/nacl/auth"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v plaintext\n\tplaintext: unencrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var key [auth.KeySize]byte
	if _, err := io.ReadFull(rand.Reader, key[:]); err != nil {
		panic(err)
	}
	log.Printf("Key: %v", hex.EncodeToString(key[:]))

	message := []byte(os.Args[1])

	digest := auth.Sum(message, &key)
	log.Printf("Digest: %v", hex.EncodeToString(digest[:]))

	fmt.Printf("\n\nVerify with: bin/verify %v %v \"%v\"\n", hex.EncodeToString(key[:]), hex.EncodeToString(digest[:]), os.Args[1])
}
