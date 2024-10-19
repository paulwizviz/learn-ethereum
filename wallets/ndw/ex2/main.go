// This example generates keys and store in wallet

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	passphrase := "passphrase"
	keydir := "./tmp"

	// STEP 1: Create private key
	privKeyOrig, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(privKeyOrig.PublicKey).Hex()

	// STEP 2: Create keystore and import private key
	ks := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.ImportECDSA(privKeyOrig, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	// Keystore address is the same that from public key
	fmt.Println(acct.Address.Hex() == address)

	// STEP 3: Private key from keystore
	ksContent, err := os.ReadFile(acct.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(ksContent, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	pk1 := crypto.FromECDSA(privKeyOrig)
	pk2 := crypto.FromECDSA(key.PrivateKey)

	// Original private key is the same as the one in keystore
	fmt.Println(bytes.Equal(pk1, pk2))

}
