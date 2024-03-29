package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/paulwizviz/learn-ethereum/internal/kstore"
)

func main() {
	acct, err := kstore.Create("./tmp", "test")
	if err != nil {
		log.Fatal(err)
	}

	privKey, pubKey, err := kstore.KeysFromStore(acct.URL.Path, "test")
	if err != nil {
		fmt.Println(err)
	}

	privB := crypto.FromECDSA(privKey)
	pubB := crypto.FromECDSAPub(pubKey)
	fmt.Println("Private key: ", hexutil.Encode(privB))
	fmt.Println("Public key: ", hexutil.Encode(pubB))
}
