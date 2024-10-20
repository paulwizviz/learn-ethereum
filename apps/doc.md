# Application development

In this section, we discuss techniques to build Web3 application.

Here we identified three approaches:

* [Web3.js](#web3js---javascript-api-libraries)
* [JSON-RRPC](#json-rpc-api)
* [Go ABI from Solidity](#go-abi-from-solidity)

## Web3.js - Javascript API Libraries

* [API Reference](https://web3js.readthedocs.io/en/v1.7.0/)
* [Web3 Tutorial](https://www.youtube.com/watch?v=2TV0r94p8OY&list=PLbbtODcOYIoFs0PDlTdxpEsZiyDR2q9aA)

## JSON-RPC API

* [Ethereum JSON-RPC Specification](https://playground.open-rpc.org/?schemaUrl=https://raw.githubusercontent.com/ethereum/eth1.0-apis/assembled-spec/openrpc.json&uiSchema%5BappBar%5D%5Bui:splitView%5D=false&uiSchema%5BappBar%5D%5Bui:input%5D=false&uiSchema%5BappBar%5D%5Bui:examplesDropdown%5D=false)

## Go ABI from Solidity

This approach involves generating source Go ABI from solidity contract using a Geth tool call [abigen](./geth.md).

### Working example

Step 1: Build an image containing `abigen` by running the command: `./scripts/ops.sh image build:node`. 

Step 2: Using the `HelloWorld` solidity contract as an example, the proceess for generating a Go ABI is [here](../scripts/solc.sh). To see a working example of a generated binding, run the command `./scripts/ops.sh solidity abi` and this will generate a packge `internal/hello`.

Step 3: Build application capable of interacting with a deployed contract is to be able to connect to a node in the Ethereum network. You need to dial a connection. Here is the step:

```go
import "github.com/ethereum/go-ethereum/ethclient"

nodeurl := "<url to node>"
conn, err := ethclient.Dial(nodeurl)

                 // Generated factory
contract, err := hello.NewHelloWorld("<contract address>", conn)

```

## References

* [The Architecture of a Web 3.0 application](https://www.preethikasireddy.com/post/the-architecture-of-a-web-3-0-application)
* [What are dApps: A 2021 guide to decentralized applications](https://limechain.tech/blog/what-are-dapps-the-2021-guide/)
* [What Are Decentralized Apps?](https://www.gemini.com/cryptopedia/decentralized-applications-defi-dapps)
* [Introduction to DApp](https://ethereum.org/en/developers/docs/dapps/)