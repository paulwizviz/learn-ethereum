// This example uses the Geth library.

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate a private key using secp256k1
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Get the public key from the private key
	publicKey := privateKey.PublicKey

	// Print the private and public keys
	fmt.Printf("Private Key: %x\n", crypto.FromECDSA(privateKey))
	fmt.Printf("Public Key: %x\n", crypto.FromECDSAPub(&publicKey))
}
