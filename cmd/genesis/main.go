package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "genesis",
	Short: "Genesis is a cli private Ethereum network provisioning tool",
	Long: `A private Ethereum network provisioning tool.
		Features:
			- Create Genesis configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func initRootCmd() {
	initCliqueCmd()
	rootCmd.AddCommand(cliqueCmd)
}

func execute() {
	initRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}

func main() {
	execute()
}
