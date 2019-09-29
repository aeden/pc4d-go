package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %v key box\n\tkey: hex-encoded 32 byte key\n\tbox: base64 encoded encrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var box []byte
	var nonce [24]byte
	var key [32]byte

	keyArg, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}
	copy(key[:], keyArg)

	box, err = base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		panic(err)
	}

	copy(nonce[:], box[:24])
	log.Printf("Nonce: %v", nonce)

	message, ok := secretbox.Open(nil, box[24:], &nonce, &key)
	if ok {
		log.Printf("Message: %s", message)
	} else {
		log.Printf("Decryption failed")
	}
}
