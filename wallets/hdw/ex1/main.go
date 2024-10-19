package main

import (
	"fmt"
	"log"

	"github.com/tyler-smith/go-bip39"
)

// //go:embed index.html
// var web embed.FS

// func router() {
// 	port := "3030"
// 	mux := http.NewServeMux()
// 	// Webserver serving htmx
// 	mux.Handle("GET /", http.FileServerFS(web))
// 	mux.HandleFunc("GET /mnemonic", func(w http.ResponseWriter, r *http.Request) {
// 		m, _ := wallet.Mnemonic(wallet.Mnemonic128)
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintf(w, m)
// 	})
// 	log.Printf("Starting web %s", port)
// 	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), mux))
// }

const (
	Mnemonic128 = 128
	Mnemonic256 = 256
)

func main() {
	e128, err := bip39.NewEntropy(Mnemonic128)
	if err != nil {
		log.Fatal(err)
	}
	m128, err := bip39.NewMnemonic(e128)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mnemonic128: ", m128)

	e256, err := bip39.NewEntropy(Mnemonic128)
	if err != nil {
		log.Fatal(err)
	}
	m256, err := bip39.NewMnemonic(e256)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mnemonic256: ", m256)

	fmt.Println("compare length", len(m128), len(m256))
}
