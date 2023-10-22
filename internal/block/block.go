package block

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

const (
	ConsensusEthash = "Ethash - proof-of-work"
	ConsensusClique = "Clique - proof-of-authority"
)

var (
	ErrNotHex = errors.New("not hex number")
)

func initGenesis() *core.Genesis {
	return &core.Genesis{
		Alloc: make(core.GenesisAlloc),
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

type AddressBalance struct {
	Address string
	Balance string
}

func parseAddressBalance(addrAmt AddressBalance) (common.Address, core.GenesisAccount, error) {

	var cAddr common.Address
	if !common.IsHexAddress(addrAmt.Address) {
		return cAddr, core.GenesisAccount{}, fmt.Errorf("%w-%s", ErrNotHex, fmt.Sprintf("input value: %s", addrAmt.Address))
	}

	amt, err := strconv.ParseUint(addrAmt.Balance, 10, 64)
	if err != nil {
		return cAddr, core.GenesisAccount{}, err
	}

	cAddr = common.HexToAddress(addrAmt.Address)
	gacct := core.GenesisAccount{
		Balance: big.NewInt(int64(amt)),
	}
	return cAddr, gacct, nil
}

func parseAddressBalances(addrBals []AddressBalance) (core.GenesisAlloc, error) {
	gAlloc := make(core.GenesisAlloc, 1)
	for _, a := range addrBals {
		addr, acct, err := parseAddressBalance(a)
		if err != nil {
			return nil, err
		}
		gAlloc[addr] = acct
	}
	return gAlloc, nil
}

func cliqueExtraData(signerAddr string) ([]byte, error) {
	prefix := "0x00000000000000000000000000000000"
	postfix := "00000000000000000000000000000000000000000000000000000000000000000"
	if !common.IsHexAddress(signerAddr) {
		return nil, fmt.Errorf("%w-%s", ErrNotHex, fmt.Sprintf("input value: %s", signerAddr))
	}
	extra := prefix + signerAddr + postfix
	return []byte(extra), nil
}

func CreateCliqueGenesis(chainID uint64, period uint64, epoch uint64, difficulty uint64, gasLimit uint64, signerAddr string, addrBals []AddressBalance) ([]byte, error) {
	genesis := initGenesis()
	genesis.Config.ChainID = big.NewInt(int64(chainID))
	genesis.Config.Clique = &params.CliqueConfig{
		Period: period,
		Epoch:  epoch,
	}
	genesis.Difficulty = big.NewInt(int64(difficulty))
	genesis.GasLimit = gasLimit
	eData, err := cliqueExtraData(signerAddr)
	if err != nil {
		return nil, err
	}
	genesis.ExtraData = eData
	alloc, err := parseAddressBalances(addrBals)
	if err != nil {
		return nil, err
	}
	genesis.Alloc = alloc
	result, err := json.Marshal(genesis)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateEthashGenesis(chainID uint64, difficulty uint64, gasLimit uint64, addrBals []AddressBalance) ([]byte, error) {
	genesis := initGenesis()
	genesis.Config.ChainID = big.NewInt(int64(chainID))
	genesis.Config.Ethash = new(params.EthashConfig)
	genesis.Difficulty = big.NewInt(int64(difficulty))
	genesis.GasLimit = gasLimit
	alloc, err := parseAddressBalances(addrBals)
	if err != nil {
		return nil, err
	}
	genesis.Alloc = alloc
	result, err := json.Marshal(genesis)
	if err != nil {
		return nil, err
	}
	return result, nil
}
