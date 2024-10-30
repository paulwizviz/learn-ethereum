// This example uses the Geth library.

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// Generate key using crypto/ecdsa
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	log.Printf("Private key - Type: %T\n", privKey)
	log.Printf("Curve: %[1]v Type: %[1]T\n", privKey.Curve)
	log.Printf("D: %[1]v Type: %[1]T\n", privKey.D)
	log.Printf("X: %[1]v Type: %[1]T\n", privKey.X)
	log.Printf("Y: %[1]v Type: %[1]T\n", privKey.Y)

	log.Println(bytes.Equal(privKey.D.Bytes(), crypto.FromECDSA(privKey)))

}
