# Wallet

This section discusses wallet related topics.

There are two types of Wallets:

* Nondeterministic Wallets
* Deterministic (Seeded) Wallets

## Nondeterministic Wallets

In the first Ethereum wallet (produced for the Ethereum pre-sale), each wallet file stored a single randomly generated private key. Such wallets are being replaced with deterministic wallets because these "old-style" wallets are in many ways inferior. For example, it is considered good practice to avoid Ethereum address reuse as part of maximizing your privacy while using Ethereum—i.e., to use a new address (which needs a new private key) every time you receive funds. You can go further and use a new address for each transaction, although this can get expensive if you deal a lot with tokens. To follow this practice, a nondeterministic wallet will need to regularly increase its list of keys, which means you will need to make regular backups. If you ever lose your data (disk failure, drink accident, phone stolen) before you’ve managed to back up your wallet, you will lose access to your funds and smart contracts. The "type 0" nondeterministic wallets are the hardest to deal with, because they create a new wallet file for every new address in a "just in time" manner. (Source: [Matering Ethereum](https://cypherpunks-core.github.io/ethereumbook/05wallets.html))

## Deterministic (Seeded) Wallets

Deterministic or "seeded" wallets are wallets that contain private keys that are all derived from a single master key, or seed. The seed is a randomly generated number that is combined with other data, such as an index number or "chain code". (Source: [Matering Ethereum](https://cypherpunks-core.github.io/ethereumbook/05wallets.html))

Specifications:

* [Hierarchical Deterministic Wallets - BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* [Hierarchical Deterministic Wallets - BIP 44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)
* [Seeds and Mnemonic Codes - BIP 39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)

## Implementations

* [Go implementation of BIP 39](https://github.com/tyler-smith/go-bip39)