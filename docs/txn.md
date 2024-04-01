# Transactions

Transactions are cryptographically signed instructions from accounts. An account will initiate a transaction to update the state of the Ethereum network. The simplest transaction is transferring ETH from one account to another. ([source](https://ethereum.org/en/developers/docs/transactions/))

There are 

* Simple transactions (payment between External accounts)
* Advance transactions (calling a contract)
* Combining several transactions (contract calling other contracts)

## Geth implementation:

Package in `core/types`.

* [Unit tests](https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction_test.go)

## References

* [The Most Misunderstood Concept in Ethereum | Ethereum transactions explained](https://www.youtube.com/watch?v=2EhKeQHFeTs)
* [Ethereum Transactions](https://www.youtube.com/watch?v=2QHGPH88WAI)
