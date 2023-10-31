package genesis

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
)

type SignerAddress string

func (s SignerAddress) String() string {
	return string(s)
}

func isValidAddress(sa SignerAddress) bool {
	return common.IsHexAddress(string(sa))
}

func cliqueExtraPrefix() string {
	prefix := "0x"
	for i := 0; i < 64; i++ {
		prefix += "0"
	}
	return prefix
}

func cliqueExtraPostfix() string {
	var postfix string
	for i := 0; i < 128; i++ {
		postfix += "0"
	}
	return postfix
}

type CliqueExtraData string

func genCliqueExtraData(signerAddrs []SignerAddress) (CliqueExtraData, error) {
	prefix := cliqueExtraPrefix()
	postfix := cliqueExtraPostfix()
	var s string
	for _, sa := range signerAddrs {
		if !isValidAddress(sa) {
			return "", fmt.Errorf("%w-%s", ErrNotHex, fmt.Sprintf("input value: %s in %s", sa, signerAddrs))
		}
		s += sa.String()
	}
	return CliqueExtraData(prefix + s + postfix), nil
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
	"extradata": "{{ .ExtraData }}",
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

func GenCliqueJSON(chainID int, period int, epoch int, signers []SignerAddress, allocs []AddressBalance) ([]byte, error) {
	tmpl := template.Must(template.New("").Funcs(fns).Parse(cliqueConfigTmpl))

	extra, err := genCliqueExtraData(signers)
	if err != nil {
		return nil, err
	}

	params := struct {
		ChainID     int
		Period      int
		Epoch       int
		ExtraData   CliqueExtraData
		Allocations []AddressBalance
	}{
		ChainID:     chainID,
		Period:      period,
		Epoch:       epoch,
		ExtraData:   extra,
		Allocations: allocs,
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, params)
	return buf.Bytes(), nil
}
