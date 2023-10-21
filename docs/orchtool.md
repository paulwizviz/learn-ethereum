# Orchestration tools

Please refer to this [doc](https://geth.ethereum.org/docs/) for description of official tools

## Clef

Clef is a tool for signing transactions and data in a secure local environment[1](https://geth.ethereum.org/docs/tools/clef/introduction).


## Geth

Geth is a command line tool to help you initialise, start and stop Ethereum nodes.

Geth commands and flags are listed [here](https://geth.ethereum.org/docs/interface/command-line-options)

<u>Source code:</u>

* [Geth cli entrypoint](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L266). The entry point uses a cli development call [urfave/cli](https://cli.urfave.org/v2/).
* [List of commands](https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/chaincmd.go)
* [List of flags](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/utils/flags.go#L81). These flags are shared by all apps under the `cmd` folders.
* [Run node command](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L343)

## Puppeth

Puppeth is a tool to help you manage genesis block.

Please refer to [using Puppeth, the Ethereum Private Network Manager](https://www.sitepoint.com/puppeth-introduction/) for more information.

<u>Source code</u>

The entrypoint is [here](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/puppeth/puppeth.go#L31)
The cli kicks off a [wizard](https://github.com/ethereum/go-ethereum/blob/afe344bcf31bfb477a6e1ad5b862a70fc5c1a22b/cmd/puppeth/wizard_genesis.go#L39) to enable use specific ethereum network configuration.
