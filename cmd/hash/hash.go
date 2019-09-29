package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v message\n\tmessage: message to hash\n", os.Args[0])
		os.Exit(1)
	}

	message := []byte(os.Args[1])
	hash := make([]byte, 64)

	sha3.ShakeSum256(hash, message)
	log.Printf("Hash: %v", hex.EncodeToString(hash))
}
