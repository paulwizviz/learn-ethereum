package genesis

import (
	"encoding/json"
	"fmt"
)

func Example_parseAddressBalances() {
	a := []AddressBalance{
		{
			Address: "7df9a875a174b3bc565e6424a0050ebc1b2d1d82",
			Balance: "300000",
		},
		{
			Address: "f41c74c9ae680c1aa78f42e5647a62f353b7bdde",
			Balance: "400000",
		},
	}

	alloc, err := parseAddressBalances(a)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(alloc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	b := []AddressBalance{
		{
			Address: "7df9a875a174b3bc565e6424a0050ebc1b2d1d82",
			Balance: "300000",
		},
		{
			Address: "f4",
			Balance: "400000",
		},
	}

	_, err = parseAddressBalances(b)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// {"0x7df9a875a174b3bc565e6424a0050ebc1b2d1d82":{"balance":"0x493e0"},"0xf41c74c9ae680c1aa78f42e5647a62f353b7bdde":{"balance":"0x61a80"}}
	// not hex number-input value: f4
}

func Example_cliqueExtraData() {
	address := "7df9a875a174b3bc565e6424a0050ebc1b2d1d82"
	extra, err := cliqueExtraData(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(extra))
	address = "7j"
	_, err = cliqueExtraData(address)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 0x000000000000000000000000000000007df9a875a174b3bc565e6424a0050ebc1b2d1d8200000000000000000000000000000000000000000000000000000000000000000
	// not hex number-input value: 7j
}

func Example_createClique() {
	a := []AddressBalance{
		{
			Address: "7df9a875a174b3bc565e6424a0050ebc1b2d1d82",
			Balance: "300000",
		},
		{
			Address: "f41c74c9ae680c1aa78f42e5647a62f353b7bdde",
			Balance: "400000",
		},
	}
	signerAddr := "7df9a875a174b3bc565e6424a0050ebc1b2d1d82"
	chainID := uint64(1234)
	period := uint64(5)
	epoch := uint64(3000)
	difficulty := uint64(1)
	gasLimit := uint64(8000000)
	genesis, err := CreateClique(chainID, period, epoch, difficulty, gasLimit, signerAddr, a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(genesis))
	// Output:
	// {"config":{"chainId":1234,"homesteadBlock":0,"eip150Block":0,"eip155Block":0,"eip158Block":0,"byzantiumBlock":0,"constantinopleBlock":0,"petersburgBlock":0,"istanbulBlock":0,"clique":{"period":5,"epoch":3000}},"nonce":"0x0","timestamp":"0x0","extraData":"0x30783030303030303030303030303030303030303030303030303030303030303030376466396138373561313734623362633536356536343234613030353065626331623264316438323030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030","gasLimit":"0x7a1200","difficulty":"0x1","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x0000000000000000000000000000000000000000","alloc":{"7df9a875a174b3bc565e6424a0050ebc1b2d1d82":{"balance":"0x493e0"},"f41c74c9ae680c1aa78f42e5647a62f353b7bdde":{"balance":"0x61a80"}},"number":"0x0","gasUsed":"0x0","parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","baseFeePerGas":null,"excessBlobGas":null,"blobGasUsed":null}
}

func Example_createEthash() {
	a := []AddressBalance{
		{
			Address: "7df9a875a174b3bc565e6424a0050ebc1b2d1d82",
			Balance: "300000",
		},
		{
			Address: "f41c74c9ae680c1aa78f42e5647a62f353b7bdde",
			Balance: "400000",
		},
	}
	chainID := uint64(1234)
	difficulty := uint64(1)
	gasLimit := uint64(8000000)
	genesis, err := CreateEthash(chainID, difficulty, gasLimit, a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(genesis))
	// Output:
	// {"config":{"chainId":1234,"homesteadBlock":0,"eip150Block":0,"eip155Block":0,"eip158Block":0,"byzantiumBlock":0,"constantinopleBlock":0,"petersburgBlock":0,"istanbulBlock":0,"ethash":{}},"nonce":"0x0","timestamp":"0x0","extraData":"0x","gasLimit":"0x7a1200","difficulty":"0x1","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","coinbase":"0x0000000000000000000000000000000000000000","alloc":{"7df9a875a174b3bc565e6424a0050ebc1b2d1d82":{"balance":"0x493e0"},"f41c74c9ae680c1aa78f42e5647a62f353b7bdde":{"balance":"0x61a80"}},"number":"0x0","gasUsed":"0x0","parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","baseFeePerGas":null,"excessBlobGas":null,"blobGasUsed":null}
}
