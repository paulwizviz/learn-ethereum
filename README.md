# Overview

[Introduction](https://ethereum.org/en/developers/docs/intro-to-ethereum/)

* Ethereum, taken as a whole, can be viewed as a
transaction-based state machine: we begin with a genesis state and incrementally execute transactions to morph it into some current state.

* The state can include such information as account balances, reputations, trust arrangements, data pertaining to information of the physical world; in short, anything that can currently be represented by a computer is admissible.

* Transactions thus represent a valid arc between two states; the ‘valid’ part is important—there exist far more invalid state changes than valid state changes. Invalid state changes might, e.g., be things such as reducing an account balance without an equal and opposite increase elsewhere.

* Transactions are collated into blocks; blocks are chained together using a cryptographic hash as a means of reference.

* Blocks function as a journal, recording a series of transactions together with the previous block and an identifier for the final state (though do not store the final state itself—that would be far too big). They also punctuate the transaction series with incentives for nodes to mine. This incentivisation takes place as a state-transition function, adding value to a nominated account.

* Mining is the process of dedicating effort (working) to bolster one series of transactions (a block) over any other potential competitor block.

## Components

* [Accounts](./docs/acct.md)
* [Blocks](./docs/blocks.md)
* [Crypto](./docs/crypto.md)
* [DApp](./docs/dapp.md)
* [Ethereum Virtual Machine](./docs/evm.md)
* [Ether](./docs/ether.md)
* [Nodes](./docs/nodes.md)
* [Smart contracts](./docs/smart.md)
* [Transactions](./docs/txn.md)
* [Wallet](./docs/wallet) - topics discussing the inner workings of wallets and interaction with Ethereum.

## Official Documentation

* [Beige paper](https://github.com/chronaeon/beigepaper/blob/master/beigepaper.pdf)
* [Yellow paper](https://ethereum.github.io/yellowpaper/paper.pdf)
* [White paper](https://ethereum.org/en/whitepaper/)
* [Ethereum Improvement Process](https://eips.ethereum.org/all)

## Other information

* [Mastering Ethereum](https://cypherpunks-core.github.io/ethereumbook/)