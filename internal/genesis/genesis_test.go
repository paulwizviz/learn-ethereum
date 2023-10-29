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

func Example_stringToAddrBal() {

	input := "f41c74c9ae680c1aa78f42e5647a62f353b7bdde:3000"
	ab, err := stringToAddrBal(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ab)

	input = "f41c74c9ae680c1aa78f42e5647a62f353b7bdde:a00"
	_, err = stringToAddrBal(input)
	if err != nil {
		fmt.Println(err)
	}

	input = "f41c74c9ae680c1aa78f42e5647a62f353b7bdd:4000"
	_, err = stringToAddrBal(input)
	if err != nil {
		fmt.Println(err)
	}

	input = "f41c74G9ae680c1aa78f42e5647a62f353b7bdd:4000"
	_, err = stringToAddrBal(input)
	if err != nil {
		fmt.Println(err)
	}

	input = "f41c74c9ae680c1aa78f42e5647a62f353b7bdd-4000"
	_, err = stringToAddrBal(input)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// {f41c74c9ae680c1aa78f42e5647a62f353b7bdde 3000}
	// unable to parse string to address balance-input: f41c74c9ae680c1aa78f42e5647a62f353b7bdde:a00
	// unable to parse string to address balance-input: f41c74c9ae680c1aa78f42e5647a62f353b7bdd:4000
	// unable to parse string to address balance-input: f41c74G9ae680c1aa78f42e5647a62f353b7bdd:4000
	// unable to parse string to address balance-input: f41c74c9ae680c1aa78f42e5647a62f353b7bdd-4000
}
