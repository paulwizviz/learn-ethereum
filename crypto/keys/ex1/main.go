// This example uses btcsuite to generate ECC private key

package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
)

func main() {
	// Create private key
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 2. Extract the public key
	publicKey := privateKey.PubKey()

	// 3. Print the private and public keys
	fmt.Printf("Private Key: %x\n", privateKey.Serialize())
	fmt.Printf("Public Key: %x\n", publicKey.SerializeUncompressed())
}
