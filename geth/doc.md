# Go-Ethereum (Geth)

Geth refers to two things:

* An open-source project for building `geth` and other supporting tools.
* A Go implementation of Ethereum client known as `geth`.

## Geth Project Architecture

The source code is organised into several major packages. 

The components are as follows:

<u>Main component</u>

* Node:
    * Responsible for managing the lifecycle of a Geth instance.
    * Handles configuration, data directories, and service components.
    * Entry point for initializing and running the Ethereum node.
* Eth:
    * Implements the Ethereum protocol.
    * Manages blockchain data, transaction pools, and consensus mechanisms.
    * Facilitates interaction with other Ethereum nodes through networking.
* Core:
    * Contains the blockchain and state database logic.
    * Handles Ethereum’s Merkle Patricia Trie, block validation, and state transitions.
* Miner:
    * Implements mining-related functionalities.
    * Responsible for creating new blocks by solving Proof-of-Work (PoW) puzzles (not relevant for Ethereum post-merge).
* Accounts:
    * Manages wallets and cryptographic key storage.
    * Supports different key storage backends (e.g., keystore files, hardware wallets).
* RPC:
    * Implements the Remote Procedure Call (RPC) layer for interacting with the Ethereum client.
    * Exposes APIs for JSON-RPC, WebSocket, and IPC communication.
* Consensus:
    * Abstracts consensus mechanisms.
    * Includes implementations for Proof of Work (PoW) and Proof of Stake (PoS).

<u>Networking components</u>

* P2P:
    * Handles peer-to-peer networking using the DevP2P protocol.
    * Manages peer discovery, connection establishment, and communication.
* LES (Light Ethereum Subprotocol):
    * Implements the protocol for light clients.
    * Allows nodes to fetch only necessary data rather than downloading the entire blockchain.
* Whisper (Deprecated): 
    * Used for decentralized messaging.

<u>Data storage</u>

* Trie:
    * Implements Ethereum’s Merkle Patricia Trie for state storage.
    * Ensures efficient and verifiable storage of account and contract data.
* LevelDB:
    * Utilized for persistent storage of blockchain data and state.
    * Supports key-value storage for blocks, state, and transaction indexes.

<u>Development utilities</u>

* Console:
    * Interactive JavaScript console for interacting with the Ethereum node.
    * Provides utilities to debug and interact with contracts and transactions.
* EVM:
    * Implements the Ethereum Virtual Machine (EVM).
    * Executes smart contracts and processes transactions.
* Test:
    * Contains utilities for unit and integration testing of Ethereum logic.

<u>Folder structure</u>

* `cmd/`: Contains the entry points for the geth and other command-line tools.
* `internal/`: Private packages used only within the Geth codebase.
* `core/`: Handles blockchain data and state transitions.
* `common/`: Contains utilities, constants and basic data structures.
* `eth/`: Implements the Ethereum node logic.
* `ethdb/`: Abstractions of levelDB and other key-value store
* `p2p/`: Implements the networking stack.
* `rpc/`: Provides the RPC layer for external communication.
* `crypto/`: Includes cryptographic utilities used across the codebase.
* `trie/`: Implements Ethereum’s state storage.
* `triedb/`: Abstraction of trie implementation.

Building components from source please refer to [this working example](../build/ethnode.dockerfile).

## Geth client

* [Getting started](https://geth.ethereum.org/docs/getting-started)
* [Fundamentals](https://geth.ethereum.org/docs/fundamentals)
* [Interacting with Geth](https://geth.ethereum.org/docs/interacting-with-geth/rpc)
* [Developers](https://geth.ethereum.org/docs/developers)
* [Monitoring](https://geth.ethereum.org/docs/monitoring/dashboards)

## Tools

* [Abigen](https://geth.ethereum.org/docs/tools/abigen) - A binding-generator for easily interacting with Ethereum using Go
* [Clef](https://geth.ethereum.org/docs/tools/clef/introduction) - A tool for signing transactions and data in a secure local environment
* [DevP2P](https://geth.ethereum.org/docs/tools/devp2p) - The DevP2P specifications define precisely how nodes should find each other and communicate.
