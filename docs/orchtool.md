# Orchestration tools

This section discuss aspects of Ethereum tools used to orchestrate the network from a source code perspective.

## Building tools from source codes

I have provided a combination of bash shell script and docker image build script so you can experiement with building tools from source codes in a Linux like environment. The scripts are here:

* [./build/gethclient](../build/gethclient)
* [./scripts/gethclient.sh](../scripts/gethclient.sh)

For this project, we focus principally on Go based source codes. The source code is [here](https://github.com/ethereum/go-ethereum).

How to use this tool:

1. Review the Docker image build specification `./build/gethclient/gethclient.dockerfile` to learn the step involved in building the tools from source
1. Build the docker image, run the command `./scripts/gethclient.sh build`
1. Log into the docker container via shell, run the command `./scripts/gethclient.sh shell`

## Clef

Clef is a tool for signing transactions and data in a secure local environment[1](https://geth.ethereum.org/docs/tools/clef/introduction).

<u>Command line</u>

[Commands and flags](https://geth.ethereum.org/docs/tools/clef/introduction)

## Geth

Geth is an execution client. Historically, an execution client alone was enough to run a full Ethereum node. However, since Ethereum swapped from proof-of-work (PoW) to proof-of-stake (PoS) based consensus, Geth needs to be coupled to another piece of software called a "consensus client"[source](https://geth.ethereum.org/docs/getting-started/consensus-clients).

### Command line

Geth commands and flags are listed [here](https://geth.ethereum.org/docs/interface/command-line-options)

### Source code

* [Geth cli entrypoint](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L266). The entry point uses a cli development call [urfave/cli](https://cli.urfave.org/v2/).
* [List of commands](https://github.com/ethereum/go-ethereum/blob/master/cmd/geth/chaincmd.go)
* [List of flags](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/utils/flags.go#L81). These flags are shared by all apps under the `cmd` folders.
* [Run node command](https://github.com/ethereum/go-ethereum/blob/de1cecb22e2a18ad70d4cb92bee122f4549c5b79/cmd/geth/main.go#L343)

## Puppeth

Puppeth was a tool for quickly spinning up and managing private development networks. 

**This is no longer available**. 
