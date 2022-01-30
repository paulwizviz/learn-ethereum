package block

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

const (
	ConsensusEthash = "Ethash - proof-of-work"
	ConsensusClique = "Clique - proof-of-authority"
)

func initGenesis() *core.Genesis {
	return &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(524288),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock:      big.NewInt(0),
			EIP150Block:         big.NewInt(0),
			EIP155Block:         big.NewInt(0),
			EIP158Block:         big.NewInt(0),
			ByzantiumBlock:      big.NewInt(0),
			ConstantinopleBlock: big.NewInt(0),
			PetersburgBlock:     big.NewInt(0),
			IstanbulBlock:       big.NewInt(0),
		},
	}
}

func CreateEthashGenesis() []byte {

	genesis := initGenesis()
	genesis.Config.Ethash = new(params.EthashConfig)
	genesis.ExtraData = make([]byte, 32)
	result, _ := json.Marshal(genesis)
	return result
}

func CreateCliqueGenesis(blocksec int) []byte {
	genesis := initGenesis()
	genesis.Difficulty = big.NewInt(1)
	genesis.Config.Clique = &params.CliqueConfig{
		Period: 15,
		Epoch:  30000,
	}
	genesis.Config.Clique.Period = uint64(blocksec)
	result, _ := json.Marshal(genesis)
	return result
}
