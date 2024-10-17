package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Create returns and an Ethereum account from arguments:
// - path: to keystore and passphrase
// - pp:   passphrase
func Create(path string, pp string) (accounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.NewAccount(pp)
	if err != nil {
		return accounts.Account{}, err
	}
	return acct, nil
}

// KeysFromStore returns ECSD private and public keys from arguments:
// - fn:  fully qualified path to keystore file
// - pp:  passphrase
func KeysFromStore(fn string, pp string) (privkey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey, err error) {
	b, err := os.ReadFile(fn)
	if err != nil {
		return nil, nil, err
	}
	// decrypt private key from keystore
	key, err := keystore.DecryptKey(b, pp)
	if err != nil {
		return nil, nil, err
	}
	privkey = key.PrivateKey
	pubKey = &key.PrivateKey.PublicKey
	return privkey, pubKey, nil
}

func main() {
	acct, err := Create("./tmp", "test")
	if err != nil {
		log.Fatal(err)
	}
	privKey, pubKey, err := KeysFromStore(acct.URL.Path, "test")
	if err != nil {
		fmt.Println(err)
	}
	// byte slice of the private key
	privB := crypto.FromECDSA(privKey)
	// byte slice of the public key
	pubB := crypto.FromECDSAPub(pubKey)
	fmt.Println("Private key: ", hexutil.Encode(privB))
	fmt.Println("Public key: ", hexutil.Encode(pubB))
}
