package genesis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

type CliqueConfig struct {
	ChainID     int
	Period      int
	Epoch       int
	Signer      string
	Allocations []AddressBalance
}

var fns = template.FuncMap{
	"isLast": func(index, size int) bool {
		return index == size-1
	},
	"notLast": func(index, size int) bool {
		return index < size-1
	},
}

var cliqueConfigTmpl = `
{
	"config": {
	  "chainId": {{ .ChainID }},
	  "homesteadBlock": 0,
	  "eip150Block": 0,
	  "eip155Block": 0,
	  "eip158Block": 0,
	  "byzantiumBlock": 0,
	  "constantinopleBlock": 0,
	  "petersburgBlock": 0,
	  "istanbulBlock": 0,
	  "muirGlacierBlock": 0,
	  "berlinBlock": 0,
	  "londonBlock": 0,
	  "arrowGlacierBlock": 0,
	  "grayGlacierBlock": 0,
	  "clique": {
		"period": {{ .Period }},
		"epoch": {{ .Epoch }}
	  }
	},
	"difficulty": "1",
	"gasLimit": "800000000",
	"extradata": "0x0000000000000000000000000000000000000000000000000000000000000000{{ .Signer }}0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	{{ $size := len .Allocations -}}
	"alloc":{
{{- range $i, $a := .Allocations -}}
	{{- if isLast $i $size }}
		"{{ $a.Address }}":{ "balance": "{{ $a.Balance }}"}
	{{- else }}
		"{{ $a.Address }}":{ "balance": "{{ $a.Balance }}"},
	{{- end }}
{{- end }} 
	}
}`

func cliqueConfigJson(chainID int, period int, epoch int, signer string, allocs []AddressBalance) []byte {
	tmpl := template.Must(template.New("").Funcs(fns).Parse(cliqueConfigTmpl))
	c := CliqueConfig{
		ChainID:     chainID,
		Period:      period,
		Epoch:       epoch,
		Signer:      signer,
		Allocations: allocs,
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, c)
	return buf.Bytes()
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

func CreateClique(chainID uint64, period uint64, epoch uint64, difficulty uint64, gasLimit uint64, signerAddr string, addrBals []AddressBalance) ([]byte, error) {
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
