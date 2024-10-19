package rlp

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func Example_str() {

	var format []string
	str := ""
	enc, _ := rlp.EncodeToBytes(str)
	for _, b := range enc {
		format = append(format, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("Empty string. %s\n", format)

	var format1 []string
	str = "dog"
	enc, _ = rlp.EncodeToBytes(str)
	for _, b := range enc {
		format1 = append(format1, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("Single string. %s\n", format1)

	var format2 []string
	strs := []string{"cat", "dog"}
	enc, _ = rlp.EncodeToBytes(strs)

	for _, b := range enc {
		format2 = append(format2, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("String slices. %s\n", format2)

	// Output:
	// Empty string. [0x80()]
	// Single string. [0x83() 0x64(d) 0x6f(o) 0x67(g)]
	// String slices. [0xc8(È) 0x83() 0x63(c) 0x61(a) 0x74(t) 0x83() 0x64(d) 0x6f(o) 0x67(g)]

}

func Example_emptyslice() {

	var format []string
	strs := []string{}
	enc, _ := rlp.EncodeToBytes(strs)
	for _, b := range enc {
		format = append(format, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("Empty string sice. %s\n", format)

	var format1 []string
	intVals := []int{}
	enc, _ = rlp.EncodeToBytes(intVals)
	for _, b := range enc {
		format1 = append(format1, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("Empty int sice. %s\n", format1)

	// Output:
	// Empty string sice. [0xc0(À)]
	// Empty int sice. []

}

func Example_boolan() {

	var format []string
	enc, _ := rlp.EncodeToBytes(true)
	for _, b := range enc {
		format = append(format, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("True. %s\n", format)

	var format1 []string
	enc, _ = rlp.EncodeToBytes(false)
	for _, b := range enc {
		format1 = append(format1, fmt.Sprintf("0x%x(%v)", b, string(b)))
	}
	fmt.Printf("False. %s\n", format1)

	// Output:
	// True. [0x1()]
	// False. [0x80()]
}

// Refer to this (https://github.com/ethereum/go-ethereum/blob/master/rlp/encode_test.go) for full test coverage.

var (
	tmpDir = func() string {
		pwd, _ := os.Getwd()
		return path.Join(pwd, "..", "..", "tmp")
	}()
)

func Example_struct() {
	c := &struct {
		A uint
		B string
	}{}

	b := &bytes.Buffer{}

	rlp.Encode(b, c)

	for _, ch := range b.Bytes() {
		fmt.Printf("%x(%v) ", ch, string(ch))
	}

	// Output:
}

func Example_emptytxn() {

	emptyTx := types.NewTransaction(
		0,
		common.HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87"),
		big.NewInt(0), 0, big.NewInt(0),
		nil,
	)

	b := &bytes.Buffer{}

	emptyTx.EncodeRLP(b)
	for _, ch := range b.Bytes() {
		fmt.Printf("%x(%v) ", ch, string(ch))
	}

	// Output:
	//
}
