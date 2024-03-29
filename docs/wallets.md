# Wallet

A wallet has multiple meaning. One refers to a user interface to enable users to create and sign transactions, as well as views of their digital assests or ether balance. The other, from a programmer's perspective, it is system used to store and manage a user’s keys. Please refer [Mastering Ethereum - Wallets](https://cypherpunks-core.github.io/ethereumbook/05wallets.html).

Wallets are further divided into:

* Nondeterministic Wallets.
* Deterministic (Seeded) Wallets.

## Non-deterministic wallet (keystore file)

A non-deterministic wallet store stores encrypted private keys.

* Keystore

### Keystore file

In this case, the wallet is in the form of JSON file. Generated using `Geth`. In this example the operation is to decrypt encrypted private key from the keystore.

```json
{
    "address": "001d3f1ef827552ae1114027bd3ecf1f086ba0f9",
    "crypto": {
        "cipher": "aes-128-ctr", <-- Cryptographic algorithm
        "ciphertext":
            "233a9f4d236ed0c13394b504b6da5df02587c8bf1ad8946f6f2b58f055507ece", <-- Encrypted private key
        "cipherparams": {
            "iv": "d10c6ec5bae81b6cb9144de81037fa15"
        },
        "kdf": "scrypt",
        "kdfparams": {
            "dklen": 32,
            "n": 262144,
            "p": 1,
            "r": 8,
            "salt":
                "99d37a47c7c9429c66976f643f386a61b78b97f3246adca89abe4245d2788407"
        },
        "mac": "594c8df1c8ee0ded8255a50caf07e8c12061fd859f4b7c76ab704b17c957e842"
    },
    "id": "4fcb2ba4-ccdb-424f-89d5-26cce304bf9c",
    "version": 3
}
```
The source code for this example is found [here](../examples/keystore/main.go)

## Deterministic Wallet

Please refer to the following specifications:

* [Hierarchical Deterministic Wallets - BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* [Hierarchical Deterministic Wallets - BIP 44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)
* [Seeds and Mnemonic Codes - BIP 39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)
    * [Go implementation of BIP 39](https://github.com/tyler-smith/go-bip39)