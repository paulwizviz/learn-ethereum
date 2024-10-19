package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Input data to be hashed (in Ethereum, this would be the public key)
	data := []byte("Hello World")

	// Create a Keccak-256 hash using Geth's crypto package
	hash := crypto.Keccak256(data)

	// Print the result in hexadecimal format
	fmt.Printf("Keccak-256 Hash: %x\n", hash)
}
