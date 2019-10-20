package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/nacl/sign"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v message\n\tmessage: message to sign\n", os.Args[0])
		os.Exit(1)
	}

	message := []byte(os.Args[1])

	publicKey, privateKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	log.Printf("Public key: %v", hex.EncodeToString(publicKey[:]))
	log.Printf("Private key: %v", hex.EncodeToString(privateKey[:]))

	signature := sign.Sign(nil, message, privateKey)
	log.Printf("Signature: %v", base64.StdEncoding.EncodeToString(signature[:]))

	fmt.Printf("\nVerify with: bin/pkverify %v %v\n\n", hex.EncodeToString(publicKey[:]), base64.StdEncoding.EncodeToString(signature[:]))
}
