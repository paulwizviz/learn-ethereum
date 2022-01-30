# Accounts

[Overview](https://ethereum.org/en/developers/docs/accounts/)

## Ethereum Address

## Externally Owned Address

An Ethereum address is a 42-character hexadecimal address derived from the last 20 bytes of the public key controlling the account with 0x appended in front. e.g., 0x71C7656EC7ab88b098defB751B7401B5f6d8976F.

The Ethereum address is the "public" address that you would need to receive funds from another party.

## Contract Address

Contract address refers to the address hosting a collection of code on the Ethereum blockchain that executes functions.

## Geth implementations

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

## References

[What is an Ethereum Address?](https://info.etherscan.com/what-is-an-ethereum-address/#:~:text=An%20Ethereum%20address%20is%20a,receive%20funds%20from%20another%20party.)