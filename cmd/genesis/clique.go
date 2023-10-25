package main

import (
	"github.com/spf13/cobra"
)

type CliqueParam struct {
	SignerAddr      string
	AddressBalances string
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
		cmd.Help()
	},
}

func initCliqueCmd() {
	cliqueNewCmd.Flags().StringVarP(&cliqueParam.SignerAddr, "signer", "s", "", "Address of minter")
	cliqueNewCmd.Flags().StringVarP(&cliqueParam.AddressBalances, "balance", "b", "", "Address balances")
	cliqueCmd.AddCommand(cliqueNewCmd)
}
