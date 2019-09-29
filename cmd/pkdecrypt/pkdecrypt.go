package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/nacl/box"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %v senderPublicKey recipientPrivateKey box\n\tsenderPublicKey: base-64 encoded public key\n\trecipientPrivateKey: base64-encoded private key\n\tbox: base64 encoded encrypted data\n", os.Args[0])
		os.Exit(1)
	}

	var encrypted []byte
	var nonce [24]byte
	var senderPublicKey [32]byte
	var recipientPrivateKey [32]byte

	senderPublicKeyArg, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}
	copy(senderPublicKey[:], senderPublicKeyArg)

	recipientPrivateKeyArg, err := base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		panic(err)
	}
	copy(recipientPrivateKey[:], recipientPrivateKeyArg)

	encrypted, err = base64.StdEncoding.DecodeString(os.Args[3])
	if err != nil {
		panic(err)
	}

	copy(nonce[:], encrypted[:24])
	log.Printf("Nonce: %v", nonce)

	message, ok := box.Open(nil, encrypted[24:], &nonce, &senderPublicKey, &recipientPrivateKey)
	if ok {
		log.Printf("Message: %s", message)
	} else {
		log.Printf("Decryption failed")
	}
}
