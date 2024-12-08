# Transactions

Transactions are cryptographically signed instructions from accounts. An account will initiate a transaction to update the state of the Ethereum network. The simplest transaction is transferring ETH from one account to another. ([source](https://ethereum.org/en/developers/docs/transactions/))

There are 

* Simple transactions (payment between External accounts)
* Advance transactions (calling a contract)
* Combining several transactions (contract calling other contracts)

## Transaction Fees

### EIP-1559

A transaction pricing mechanism that includes fixed-per-block network fee that is burned and dynamically expands/contracts block sizes to deal with transient congestion.

*GasFeeCap*

GasFeeCap (also referred to as maxFeePerGas in EIP-1559) is a parameter introduced in Ethereum’s London upgrade to define the maximum total fee per gas unit that a user is willing to pay for a transaction.

It is used in the dynamic fee model introduced by EIP-1559 to separate transaction fees into two components:

1. Base Fee: A mandatory fee burned by the network, dynamically adjusted based on network demand.
1. Priority Fee (Tip): An optional fee paid to miners to prioritize the transaction.

* Caps Total Cost: GasFeeCap ensures that users don’t pay more than a specified amount per gas unit, even if the network fee (base fee) increases significantly during transaction execution.
* Defines the Upper Limit: The actual gas price paid per unit of gas is:

`Effective Gas Price= min(GasFeeCap,Base Fee + GasTipCap)`

* GasTipCap is the tip (priority fee) offered to miners.
* Base Fee is adjusted by the network.

Key Characteristics:

* Measured in Wei (the smallest Ether unit).
* Protects users from excessive fees in case of network congestion.
* Can be thought of as a safety cap to control the maximum transaction cost.

When is GasFeeCap Used?

* In EIP-1559 dynamic fee transactions (Type 0x2 transactions).
* Replaces the old gasPrice field used in legacy transactions.

Example: Gas Fee Calculation

Assumptions:
	•	Base Fee = 50 Gwei
	•	GasTipCap (Priority Fee) = 2 Gwei
	•	GasFeeCap = 70 Gwei

Calculation:

1.	Base Fee + Tip:

`Base Fee + GasTipCap = 50 + 2 = 52Gwei`

2.	Compare with GasFeeCap:

`Effective Gas Price = min(GasFeeCap, Base Fee + GasTipCap) = min(70, 52) = 52Gwei`

The transaction will execute with an effective gas price of 52 Gwei.

**GasTipCap**

GasTipCap (also referred to as maxPriorityFeePerGas in EIP-1559) is a parameter in Ethereum’s EIP-1559 transaction model that specifies the maximum tip (in Wei) a user is willing to pay per unit of gas to incentivize miners to prioritize their transaction.

Purpose of GasTipCap.

1. Miner Incentive: Miners receive the priority fee (tip) as direct compensation for including a transaction in a block.
2. Transaction Priority: Higher GasTipCap values increase the likelihood of a transaction being included in the next block, especially in congested conditions.
3. User Control: Lets users specify how much extra they are willing to pay beyond the mandatory base fee.

## Geth implementation:

Package in `core/types`.

* [Unit tests](https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction_test.go)

## References

* [The Most Misunderstood Concept in Ethereum | Ethereum transactions explained](https://www.youtube.com/watch?v=2EhKeQHFeTs)
* [Ethereum Transactions](https://www.youtube.com/watch?v=2QHGPH88WAI)
