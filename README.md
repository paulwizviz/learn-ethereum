# Overview

Welcome to my collection of original (created by me) and reference to materials (created by others) about Ethereum.

## Ethereum in a nutshell

Ethereum is a kind of transaction-based state machine where each machine is hosted on decentralised and independently executing computing nodes. The machine operations begins with a genesis state and incrementally mutating by executing transactions into some current state.

The state can include such information about a physical thing (e.g. house, cars) or an intangible thing (e.g. money, reputation) that can be represented by a computer.

A transaction represents a cryptographically signed action (e.g. transfer ownership of a property) initiated by a human or some computational data source via an agent. An agent, also known as ABI (Application Binary Interface) is an interface between the initiator of a transaction and the Ethereum state machine.

A transaction can be viewed as a connection between two states. It can be viewed as transitioning one state to another valid state or invalid state. A valid state is one where its relationship to prior states is consistent. An invalud state is one that is not consistent with prior states. Each state machine in the network of state machines is responsible for validing transaction. There is no single machine to validate transactions on behalf of every machines.

Transactions are collated into blocks through a process known as minimg (via a process known as Proof-of-Work) or minting (via a process known as Proof-of-Stake). The machine are rewarded with a sum of reward denominated in Ether when it collate the transactions into blocks.

Blocks are chained together using a cryptographic hash as a means of reference. Blocks function as a journal, recording a series of transactions together with the previous block and an identifier for the final state.

For further details, please refer to the following documents:

* [Beige paper](https://github.com/chronaeon/beigepaper/blob/master/beigepaper.pdf)
* [Yellow paper](https://ethereum.github.io/yellowpaper/paper.pdf)
* [White paper](https://ethereum.org/en/whitepaper/)
* [Introduction to Ethereum](https://ethereum.org/en/developers/docs/intro-to-ethereum/)
* [Mastering Ethereum](https://cypherpunks-core.github.io/ethereumbook/)

## Topics

<u>Governance</u>

* [Introduction to Ethereum governance](https://arvanaghi.com/blog/explaining-the-genesis-block-in-ethereum/)
* [Ethereum Improvement Process](https://eips.ethereum.org/all)

</u>Agent or application developments</u>

* [DApp](./docs/dapp.md)
* [Smart contracts](./docs/smart.md)
* [Wallet](./docs/wallets.md)

<u>Platform engineering</u>

* [Setting up a node and private network](./docs/network.md)
* [Orchestration tools](./docs/orchtool.md)

<u>System internals</u>

* [Accounts](./docs/acct.md)
* [Blocks](./docs/blocks.md)
* [Cryptographic tools](./docs/crypto.md)
* [Ethereum Virtual Machine](./docs/evm.md)
* [Transactions](./docs/txn.md)

## Disclaimer

* The contents in this project are intended for educational purpose only.
* This contents are constantly updated and items may be removed and modified without warning.

## Copyright

Unless otherwise specificed, the copyright in this project are assigned as follows.

Copyright 2022 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.