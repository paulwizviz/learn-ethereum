# Accounts

This section discuss Ethereum accounts.

An Ethereum account is an entity with an ether (ETH) balance that can send transactions on Ethereum. Accounts can be user-controlled or deployed as smart contracts ([source](https://ethereum.org/en/developers/docs/accounts/)).



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
