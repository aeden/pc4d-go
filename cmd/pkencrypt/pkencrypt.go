package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/nacl/box"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v plaintext\n\tplaintext: unencrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var nonce [24]byte
	message := []byte(os.Args[1])

	senderPublicKey, senderPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	log.Printf("Sender public key: %v", base64.StdEncoding.EncodeToString(senderPublicKey[:]))
	log.Printf("Sender private key: %v", base64.StdEncoding.EncodeToString(senderPrivateKey[:]))

	recipientPublicKey, recipientPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	log.Printf("Recipient public key: %v", base64.StdEncoding.EncodeToString(recipientPublicKey[:]))
	log.Printf("Recipient private key: %v", base64.StdEncoding.EncodeToString(recipientPrivateKey[:]))

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	log.Printf("Nonce: %v", nonce)

	// The nonce is prefixed to the encrypted data. The correct nonce must be used to decrypt the encrypted data
	box := box.Seal(nonce[:], message, &nonce, recipientPublicKey, senderPrivateKey)
	log.Printf("Box: %v", base64.StdEncoding.EncodeToString(box))
}
