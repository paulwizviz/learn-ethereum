// This example demonstrates the creation of
// a wallet and a private key

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	passphrase := "passphrase"
	keydir := "./tmp"

	// STEP 1: Create wallet with a default private key
	ks := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.NewAccount(passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account address:", acct.Address)
	fmt.Println("Keystore location: ", acct.URL.Path)

	// STEP 2: Extract keys from store
	ksContent, err := os.ReadFile(acct.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(ksContent, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	privKey := key.PrivateKey
	pubKey := &privKey.PublicKey
	fmt.Println("Raw Private Key", privKey)
	fmt.Println("Raw Public Key", pubKey)

	// Byte slice of keys
	privKeyB := crypto.FromECDSA(privKey)
	pubKeyB := crypto.FromECDSAPub(pubKey)
	fmt.Println("Byte slice Private Key", privKeyB)
	fmt.Println("Byte slice Public Key", pubKeyB)

	// Hex version of keys
	fmt.Println("Private key: ", hexutil.Encode(privKeyB))
	fmt.Println("Public key: ", hexutil.Encode(pubKeyB))
}
