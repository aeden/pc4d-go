package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ed25519"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v message\n\tmessage: message to sign\n", os.Args[0])
		os.Exit(1)
	}

	message := []byte(os.Args[1])

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	log.Printf("Public key: %v", hex.EncodeToString(publicKey))
	log.Printf("Private key: %v", hex.EncodeToString(privateKey))

	signature := ed25519.Sign(privateKey, message)
	log.Printf("Signature: %v", hex.EncodeToString(signature))
}
