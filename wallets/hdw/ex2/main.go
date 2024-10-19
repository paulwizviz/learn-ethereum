package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var (
	ErrMnemonic = errors.New("Mnemonic error")
)

const (
	Mnemonic128 = 128
	Mnemonic256 = 256
)

// Mnemonic takes a secret phrase and bitsize argemnents to generate
// a mnemonic if no error ecountered.
func Mnemonic(bitsize int) (string, error) {
	entropy, err := bip39.NewEntropy(bitsize)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	m, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("%w-%s", ErrMnemonic, err.Error())
	}
	return m, nil
}

func MnemonicEntropy(entropy string) (string, error) {
	m, err := bip39.NewMnemonic([]byte(entropy))
	if err != nil {
		return "", nil
	}
	return m, nil
}

// MasterHDKey takes a mnemonic and passphrase arguments to return a bip32 key
// if no error encountered.
func MasterHDKey(mnemonic string, passphrase string) (*bip32.Key, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	mkey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	return mkey, nil
}

func main() {

	// Create Mnemonic
	m, err := Mnemonic(Mnemonic128)
	if err != nil {
		log.Fatal(err)
	}

	// Create master key with a combination of Mnemonic and passphrase
	passphrase := "hello"
	mkey, err := MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey.Key)

	// Recover key from Mnemonic and passphrase
	mkey1, err := MasterHDKey(m, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %[1]T Value: %[1]v\n", mkey1.Key)

	mkeyHex := hexutil.Encode(mkey.Key)
	mkey1Hex := hexutil.Encode(mkey1.Key)
	fmt.Println(mkeyHex, mkey1Hex)
}
