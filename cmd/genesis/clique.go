package main

import (
	"reflect"

	"github.com/spf13/cobra"
)

type CliqueParam struct {
	ChainID         int
	Period          int
	Epoch           int
	SignerAddr      string
	AddressBalances string
	output          string
}

var cliqueParam CliqueParam = CliqueParam{}

var cliqueCmd = &cobra.Command{
	Use:   "clique",
	Short: "clique is a subcommand to orchestrate PoA consensus.",
	Long: `Clique is a Proof-of-Authority (PoA) network.
		This subcommand enable you to provision a POA setup`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cliqueNewCmd = &cobra.Command{
	Use:   "new",
	Short: "new clique genesis.json file",
	Run: func(cmd *cobra.Command, args []string) {
		if reflect.DeepEqual(cliqueParam, CliqueParam{}) {
			cmd.Help()
			return
		}
	},
}

func initCliqueCmd() {
	cliqueNewCmd.Flags().IntVarP(&cliqueParam.ChainID, "chainid", "c", 12345, "network chain id")
	cliqueNewCmd.Flags().IntVarP(&cliqueParam.Period, "period", "p", 5, "clique period")
	cliqueNewCmd.Flags().IntVarP(&cliqueParam.Epoch, "e", "epoch", 30000, "clique epoch")
	cliqueNewCmd.Flags().StringVarP(&cliqueParam.SignerAddr, "signer", "s", "", "Address of minter")
	cliqueNewCmd.Flags().StringVarP(&cliqueParam.AddressBalances, "balance", "b", "", "Address balances \"<addr>:<bal>,<addr>:<bal>,...\"")
	cliqueNewCmd.Flags().StringVarP(&cliqueParam.output, "output", "o", "genesis.json", "location of genesis.json")
	cliqueCmd.AddCommand(cliqueNewCmd)
}
