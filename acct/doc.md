# Accounts

An Ethereum account is an entity with an ether (ETH) balance that can send transactions on Ethereum. Accounts can be user-controlled or deployed as smart contracts ([source](https://ethereum.org/en/developers/docs/accounts/)).

There are two well known types of Ethereum accounts:

* Externally Owned Account
* Contract Account

There is another "official" contract-based version known as [ERC 4337](https://eips.ethereum.org/EIPS/eip-4337) also called "Account abstraction".

## Externally Owned Accounts (EOAs):

* Controlled by private keys (users interact with Ethereum using EOAs).
* Contains:
	* **Nonce**: Prevents replay attacks by counting transactions.
	* **Balance**: The amount of Ether in the account.
	* **No code or storage**.

## Contract Accounts:

* Associated with deployed smart contracts.
* Contains:
    * **Nonce**: Tracks the number of contract creations by the account.
    * **Balance**: Holds Ether.
    * **Code**: The compiled bytecode of the smart contract.
    * **Storage**: Persistent key-value storage for the contract’s state.

## ERC-4337

* [ERC 4337 in 7 minutes (the 5 roles in Account Abstraction)](https://www.youtube.com/watch?v=FjK5rYznJjU)
* [The Transformation of Ethereum Wallet: The Potential and Challenges of Account Abstraction and ERC-4337](https://medium.com/@sevenxventures/the-transformation-of-ethereum-wallet-the-potential-and-challenges-of-account-abstraction-and-c0dbdb384c7e)

## Source code

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
