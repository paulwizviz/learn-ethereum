# Clients

A client is an implementation of Ethereum that verifies data against the protocol rules and keeps the network secure. A client is an implementation of Ethereum that verifies data against the protocol rules and keeps the network secure. 

A client is implemented in several languages: Go, Rust, C++ and more.

There are two types of clients:

* Executing client.
* Consensus client.

## Executing client

The executing client listens to new transactions broadcasted in the network, executes them in EVM, and holds the latest state and database of all current Ethereum data.

An executing client is developed in multiple languages. 

A version that is implemented in Go is known as [Geth](./geth.md).

## Consensus client

The consensus client implements the proof-of-stake consensus algorithm, which enables the network to achieve agreement based on validated data from the execution client. There is also a third piece of software, known as a 'validator' that can be added to the consensus client, allowing a node to participate in securing the network.

A version implemented in Go is known as [Prysm](https://docs.prylabs.network/docs/getting-started)

## References

* [Clients and Nodes](https://ethereum.org/en/developers/docs/nodes-and-clients/)
* [Consensus client](https://geth.ethereum.org/docs/getting-started/consensus-clients).
* [Execution client](https://ethereum.org/en/developers/docs/nodes-and-clients/#execution-clients)

