# Go-Ethereum (Geth)

This section discsses the operations and inner workings of Geth and its support tools:

* Clef
* Abigen
* Devp2p

## Geth as Executing Client

Geth commands and flags are listed [here](https://geth.ethereum.org/docs/interface/command-line-options)

### Source code

* [Geth cli entrypoint](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L266). The entry point uses a cli development call [urfave/cli](https://cli.urfave.org/v2/).
* [List of commands](https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/chaincmd.go)
* [List of flags](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/utils/flags.go#L81). These flags are shared by all apps under the `cmd` folders.
* [Run node command](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L343)

## Clef

Clef is a tool for signing transactions and data in a secure local environment. It is intended to become a more composable and secure replacement for Geth's built-in account management (see [Introduction to Clef](https://geth.ethereum.org/docs/tools/clef/introduction)).

## Abigen

Abigen is a binding-generator for easily interacting with Ethereum using Go. Abigen creates easy-to-use, type-safe Go packages from Ethereum smart contract definitions known as ABIs (see [https://geth.ethereum.org/docs/tools/abigen](https://geth.ethereum.org/docs/tools/abigen)).

## devp2p

DevP2P is a set of network protocols that form the Ethereum peer-to-peer network. The DevP2P specifications define precisely how nodes should find each other and communicate. Geth implements the DevP2P specifications in Go (see [devp2p](https://geth.ethereum.org/docs/tools/devp2p)).

## Working examples

In this project you will end working examples demonstrating:

* Building `geth` and other tools from source code
* Using tools

### Building tools from source codes

I have provided a combination of bash shell script and docker image build script so you can experiement with building tools from source codes in a Linux like environment. The source codes behind these tools are found here:

* [Geth tools](../build/ethnode.dockerfile)
* [Solidity compiler](../build/solc.dockerfile)

### Using tools

You can find working examples of `geth` and `clef` in a container. To use these tools, follow these steps:

* STEP 1: Build the docker image, run the command `./scripts/ops.sh image build`
* STEP 2: Log into the docker container via shell, run the command `./scripts/ops.sh geth`
