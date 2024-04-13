# Application development

In this section, we discuss techniques to build Web3 application.

## Architecture

* [The Architecture of a Web 3.0 application](https://www.preethikasireddy.com/post/the-architecture-of-a-web-3-0-application)
* [What are dApps: A 2021 guide to decentralized applications](https://limechain.tech/blog/what-are-dapps-the-2021-guide/)
* [What Are Decentralized Apps?](https://www.gemini.com/cryptopedia/decentralized-applications-defi-dapps)
* [Introduction to DApp](https://ethereum.org/en/developers/docs/dapps/)

## Web3.js - Javascript API Libraries

* [API Reference](https://web3js.readthedocs.io/en/v1.7.0/)
* [Web3 Tutorial](https://www.youtube.com/watch?v=2TV0r94p8OY&list=PLbbtODcOYIoFs0PDlTdxpEsZiyDR2q9aA)

## JSON-RPC API

* [Ethereum JSON-RPC Specification](https://playground.open-rpc.org/?schemaUrl=https://raw.githubusercontent.com/ethereum/eth1.0-apis/assembled-spec/openrpc.json&uiSchema%5BappBar%5D%5Bui:splitView%5D=false&uiSchema%5BappBar%5D%5Bui:input%5D=false&uiSchema%5BappBar%5D%5Bui:examplesDropdown%5D=false)

## Solidity

* [Official Reference](https://soliditylang.org/)
* [Deploy Your First Smart Contract](https://www.web3.university/tracks/create-a-smart-contract/deploy-your-first-smart-contract)
* [Smart Contract with Golang](https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9)

### Compiling solidity code

To compile solidity, you will need a solidity compiler. You can find them in many tools such as [Remix](https://remix.ethereum.org/) or [Hardhat](https://hardhat.org/). 

You can also download compiler or build one from source. Here is a [working example](../build/solc/solc.dockerfile) demonstrating the steps to build from source. To build the compiler image run the command `./scripts/ops.sh image build:solc`.

To help you appreciate the process of compiling solidity file (see [./solidity/HelloWorld.sol](../solidity/HelloWorld.sol)), please refer to this [script](../scripts/solc.sh). Run the command `./script/ops.sh solidity compile` and this will produce an `HelloWorld.abi` and `HelloWorld.bin`. 


## Developing Go-based Application

The first step in building a Go-based application is to generate a Go binding of the solidity contract from the contract abi file. Geth provides a tool call [abigen](./geth.md). You can find an instance of the tool by building an image running this command: `./scripts/ops.sh image build:node`. 

Using the `HelloWorld` solidity contract as an example, the proceess for generating a Go-binding is [here](../scripts/solc.sh). To see a working example of a generated binding, run the command `./scripts/ops.sh solidity abi` and this will generate a packge `internal/hello`.

A key aspect of building application capable of interacting with a deployed contract is to be able to connect to a node in the Ethereum network. You need to dial a connection. Here is the step:

```go
import "github.com/ethereum/go-ethereum/ethclient"

nodeurl := "<url to node>"
conn, err := ethclient.Dial(nodeurl)

                 // Generated factory
contract, err := hello.NewHelloWorld("<contract address>", conn)

```