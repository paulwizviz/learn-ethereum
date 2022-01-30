package keccak

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func Example_one() {

	data := []byte("Hello")
	state := crypto.NewKeccakState()

	h := crypto.HashData(state, data)
	fmt.Println(h)

	b := &bytes.Buffer{}

	rlp.Encode(b, h)
	for _, ch := range b.Bytes() {
		fmt.Printf("%x(%v) ", ch, string(ch))
	}

	// Output:
	// 0x06b3dfaec148fb1bb2b066f10ec285e7c9bf402ab32aa78a5d38e34566810cd2
	// a0( ) 6() b3(³) df(ß) ae(®) c1(Á) 48(H) fb(û) 1b() b2(²) b0(°) 66(f) f1(ñ) e() c2(Â) 85() e7(ç) c9(É) bf(¿) 40(@) 2a(*) b3(³) 2a(*) a7(§) 8a() 5d(]) 38(8) e3(ã) 45(E) 66(f) 81() c() d2(Ò)
}
