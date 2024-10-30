# Cryptographic Tools

In this section, you will find references to crytographic tools used in Ethereum.

## Keys

Ethereum and Bitcoin both use the same elliptic curve, `secp256k1`, for generating and verifying cryptographic keys.

### Working examples

* [Example 1](./keys/ex1/main.go) use `standard lib crypto and geth` to generate keys.
* [Example 2](./keys/ex2/main.go) use `ethereum/go-ethereum/crypto` to generate keys.

## Keccak (SHA-3)

This is a hashing function.

* [Keccak by binance](https://academy.binance.com/en/glossary/keccak)
* [Geth source code](https://github.com/ethereum/go-ethereum/blob/master/crypto/crypto.go)

This is used as part of a process to generate Ethereum Address from the public key. The steps to generate address are:

1. Public Key: Like Bitcoin, the public key is derived from the private key using secp256k1.
1. Keccak-256 Hashing: Instead of using SHA-256 followed by RIPEMD-160, Ethereum applies a Keccak-256 hash (not SHA-3) to the public key.
1. Last 20 Bytes: The last 20 bytes of the resulting Keccak-256 hash are taken to create the Ethereum address (a 40-character hexadecimal string prefixed by 0x).

### Working examples

* Example 1 - This [example](./keccak/ex1/main.go) demonstrate the use of Geth crypto tool to hash a message.

## Recursive Length Prefix (RLP)

This is a serialization tool.

* [Ethereum Wiki](https://eth.wiki/fundamentals/rlp)
* [Data structure in Ethereum | Episode 1: Recursive Length Prefix (RLP) Encoding/Decoding.](https://medium.com/coinmonks/data-structure-in-ethereum-episode-1-recursive-length-prefix-rlp-encoding-decoding-d1016832f919)
* [What is RLP encoding in blockchain?](https://www.codetd.com/en/article/12305023)
* [Go coding format](https://github.com/ethereum/go-ethereum/blob/59f0e8ae60c777bef384f045edf2a816c4a3ca9d/rlp/encode_test.go#L91)