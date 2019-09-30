package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/nacl/sign"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %v publicKey signedMessage\n\tpublicKey: hex-encoded public key\n\tsignedMessage: base64-encoded signed message\n", os.Args[0])
		os.Exit(1)
	}

	var publicKey [32]byte

	publicKeyArg, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}
	copy(publicKey[:], publicKeyArg)

	message, err := base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		panic(err)
	}

	_, ok := sign.Open(nil, message, &publicKey)
	if ok {
		log.Printf("Verified")
	} else {
		log.Printf("NOT verified")
	}
}
