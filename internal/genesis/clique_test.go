package genesis

import "fmt"

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

func Example_cliqueConfigJson() {
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
	result := cliqueConfigJson(1234, 5, 3000, "7df9a875a174b3bc565e6424a0050ebc1b2d1d82", a)
	fmt.Println(result)
	// Output:
	// [10 123 10 9 34 99 111 110 102 105 103 34 58 32 123 10 9 32 32 34 99 104 97 105 110 73 100 34 58 32 49 50 51 52 44 10 9 32 32 34 104 111 109 101 115 116 101 97 100 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 101 105 112 49 53 48 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 101 105 112 49 53 53 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 101 105 112 49 53 56 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 98 121 122 97 110 116 105 117 109 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 99 111 110 115 116 97 110 116 105 110 111 112 108 101 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 112 101 116 101 114 115 98 117 114 103 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 105 115 116 97 110 98 117 108 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 109 117 105 114 71 108 97 99 105 101 114 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 98 101 114 108 105 110 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 108 111 110 100 111 110 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 97 114 114 111 119 71 108 97 99 105 101 114 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 103 114 97 121 71 108 97 99 105 101 114 66 108 111 99 107 34 58 32 48 44 10 9 32 32 34 99 108 105 113 117 101 34 58 32 123 10 9 9 34 112 101 114 105 111 100 34 58 32 53 44 10 9 9 34 101 112 111 99 104 34 58 32 51 48 48 48 10 9 32 32 125 10 9 125 44 10 9 34 100 105 102 102 105 99 117 108 116 121 34 58 32 34 49 34 44 10 9 34 103 97 115 76 105 109 105 116 34 58 32 34 56 48 48 48 48 48 48 48 48 34 44 10 9 34 101 120 116 114 97 100 97 116 97 34 58 32 34 48 120 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 55 100 102 57 97 56 55 53 97 49 55 52 98 51 98 99 53 54 53 101 54 52 50 52 97 48 48 53 48 101 98 99 49 98 50 100 49 100 56 50 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 34 44 10 9 34 97 108 108 111 99 34 58 123 10 9 9 34 55 100 102 57 97 56 55 53 97 49 55 52 98 51 98 99 53 54 53 101 54 52 50 52 97 48 48 53 48 101 98 99 49 98 50 100 49 100 56 50 34 58 123 32 34 98 97 108 97 110 99 101 34 58 32 34 51 48 48 48 48 48 34 125 44 10 9 9 34 102 52 49 99 55 52 99 57 97 101 54 56 48 99 49 97 97 55 56 102 52 50 101 53 54 52 55 97 54 50 102 51 53 51 98 55 98 100 100 101 34 58 123 32 34 98 97 108 97 110 99 101 34 58 32 34 52 48 48 48 48 48 34 125 32 10 9 125 10 125]
}
