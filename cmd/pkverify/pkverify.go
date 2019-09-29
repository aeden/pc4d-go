package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ed25519"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %v publicKey signature message\n\tpublicKey: hex-encoded public key\n\tsignature: hex-encoded signature\n\tmessage: message string\n", os.Args[0])
		os.Exit(1)
	}

	publicKey, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}

	signature, err := hex.DecodeString(os.Args[2])
	if err != nil {
		panic(err)
	}

	message := []byte(os.Args[3])

	ok := ed25519.Verify(publicKey, message, signature)
	if ok {
		log.Printf("Verified")
	} else {
		log.Printf("NOT verified")
	}
}
