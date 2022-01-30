package block

import "fmt"

func Example_creategenesis() {
	genesis := CreateEthashGenesis()
	fmt.Println(string(genesis))

	// Output:
	//
}
