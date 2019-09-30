package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/nacl/auth"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %v key digest message\n\tkey: the key in hex format\n\tdigest: the digest in hex format\n\tmessage: unencrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var key [auth.KeySize]byte

	keyArg, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}
	copy(key[:], keyArg)

	digest, err := hex.DecodeString(os.Args[2])

	message := []byte(os.Args[3])

	ok := auth.Verify(digest, message, &key)
	if ok {
		log.Printf("Verified")
	} else {
		log.Printf("NOT verified")
	}
}
