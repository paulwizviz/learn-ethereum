# Accounts

An Ethereum account is an entity with an ether (ETH) balance that can send transactions on Ethereum. Accounts can be user-controlled or deployed as smart contracts ([source](https://ethereum.org/en/developers/docs/accounts/)).

There are two well known types of Ethereum accounts:

* Externally Owned Account
* Contract Account

There is another "official" contract-based version known as [ERC 4337](https://eips.ethereum.org/EIPS/eip-4337) also called "Account abstraction".

## Creating an EOA using Clef

Please refer to [Clef official documentation](https://geth.ethereum.org/docs/tools/clef/introduction).

To create an EOA follow this steps:

1. `cd` to the root of this project.
1. Run the command `./scripts/gethclient.sh build` to build the image.
1. Run the command `./scripts/gethclient.sh shell` to opem a shell in the container.
1. In the shell run the following command:
```sh
clef newaccount --keystore ./data
```
1. You will find an account created in the directory `./tmp`

At the end of the step you will find a file created in the following format `UTC-<timestamp>-<account id in hash>`. For example: 
```
UTC--2024-03-24T17-54-26.876362756Z--0caf5082c30ffafdbd2941fa5ee64d2156d2a0ab
```

You can view the file in `./tmp`.

The file contains in a JSON format (representing the account state) looking like this:
```json
{
    "address":"3e056424aaf3e6cb819571eacfcdaad23e967fe1",
    "crypto": {
        "cipher":"aes-128-ctr",        "ciphertext":"04e097119722c127a931712878daedd8a41a2f9f311ecd8a7d5f43e808be07cc",
        "cipherparams":{
            "iv":"c85f8bfd8476b275a690c7042771a447"
        },
        "kdf":"scrypt",
        "kdfparams":{
            "dklen":32,
            "n":262144,
            "p":1,
            "r":8,
            "salt":"e4616683a81c751b0d2b577fb8beb7cd9960f7e6368888ad6624907a22b16278"
        },
        "mac":"3fecde41d0c4d04c1a23a7e29bdb2b7e8e433df8070f446677691f4a3570cdcd"
    },
    "id":"eb822fd0-0857-4043-ad08-fefd972fdb0a",
    "version":3
}
```

## Account management with Geth

From [Go Account managament](https://geth.ethereum.org/docs/developers/dapp-developer/native-accounts)):

> Once an encrypted keystore for Ethereum accounts exists it, it can be used to manage accounts for the entire account lifecycle requirements of a Go native application. This includes the basic functionality of creating new accounts and deleting existing ones as well as updating access credentials, exporting existing accounts, and importing them on other devices (see.

### Source code

The source code implementing the account related operations are here.

* [Package](https://github.com/ethereum/go-ethereum/tree/master/accounts)
* [Address type](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/common/types.go#L198)
* [New Key](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/accounts/keystore/key.go#L167)

* Steps creating account
    * [accountCreate](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/cmd/geth/accountcmd.go#L262)
    * [keystore.StoreKey](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/accounts/keystore/passphrase.go#L101)
    * [keystorePassPhrase](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/accounts/keystore/passphrase.go#L106)
    * [EncryptKey](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/accounts/keystore/passphrase.go#L186)
    * [EncryptDataV3](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/accounts/keystore/passphrase.go#L141)
    * [Keccak256](https://github.com/ethereum/go-ethereum/blob/fb3a6528cfa49f623570575c4fe9e8a716cfcdf7/crypto/crypto.go#L77)

### Working examples

* Keystore
    * [main package](../examples/keystore/main.go)
    * [keystore package](../internal/kstore/kstore.go)

## ERC-4337

* [ERC 4337 in 7 minutes (the 5 roles in Account Abstraction)](https://www.youtube.com/watch?v=FjK5rYznJjU)
* [The Transformation of Ethereum Wallet: The Potential and Challenges of Account Abstraction and ERC-4337](https://medium.com/@sevenxventures/the-transformation-of-ethereum-wallet-the-potential-and-challenges-of-account-abstraction-and-c0dbdb384c7e)