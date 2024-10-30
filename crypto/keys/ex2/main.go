// This example uses btcsuite to generate ECC private key

package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// Create private key using Geth crypto
	privKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	log.Printf("Private key - Type: %T\n", privKey)
	log.Printf("Curve: %[1]v Type: %[1]T\n", privKey.Curve)
	log.Printf("D: %[1]v Type: %[1]T\n", privKey.D)
	log.Printf("X: %[1]v Type: %[1]T\n", privKey.X)
	log.Printf("Y: %[1]v Type: %[1]T\n", privKey.Y)

	log.Printf("Public key - Type: %T\n", privKey.PublicKey)
	log.Printf("Curve: %[1]v Type: %[1]T\n", privKey.PublicKey.Curve)
	log.Printf("X: %[1]v Type: %[1]T\n", privKey.PublicKey.X)
	log.Printf("Y: %[1]v Type: %[1]T\n", privKey.PublicKey.Y)

	log.Println(bytes.Equal(privKey.D.Bytes(), crypto.FromECDSA(privKey)))

}
