# Transactions

Transactions are cryptographically signed instructions from accounts. An account will initiate a transaction to update the state of the Ethereum network. The simplest transaction is transferring ETH from one account to another. ([source](https://ethereum.org/en/developers/docs/transactions/))

There are several transactions data types.

| Type | Identifier | Introduced In | Description |
| --- | --- | --- | --- |
| Legacy Transaction | 0x0 | Pre-EIP-1559 | The original Ethereum transaction type. Includes gasPrice for gas fees. |
| Access List Transaction | 0x1 | EIP-2930 (Berlin Upgrade) | Adds access lists for state optimization, includes gasPrice. |
| Dynamic Fee Transaction |	0x2	| EIP-1559 (London Upgrade) | Introduces maxFeePerGas and maxPriorityFeePerGas for a dynamic fee model (replacing gasPrice). |

Recommended Transaction Type Based on Use Case.

| Use Case | Recommended Type | Reason |
| --- | --- | --- |
| General transactions on modern Ethereum | Dynamic Fee (0x2) | Predictable fees and widely supported. |
| Interacting with EIP-1559-compatible networks | Dynamic Fee (0x2) | Optimized for current fee structures. |
| Complex smart contract interactions | Access List (0x1) | Reduced gas costs via access lists. |
| Private or older networks | Legacy (0x0) | Compatibility with networks that don’t support EIP-1559 or EIP-2930. |
| Layer 2 solutions | Dynamic Fee (0x2) or 0x1 | Most L2s support EIP-1559 and/or access lists for state optimization. |

## Pre-EIP-1559 (Legacy Transaction)

When to use:

* On older networks or private chains that do not support EIP-1559 or EIP-2930.
* If you’re interacting with tooling or services that only support legacy transactions.

Disadvantages:

* Overpayment Risk: Fixed gasPrice can lead to overpaying during high network demand.
* Limited Features: No dynamic fees or access list optimizations.

Fields:

* `gasPrice`: Fixed price per gas unit.

## EIP-2930  (Access List Transaction)

When to Use:

* When interacting with complex smart contracts that could benefit from access lists.
* If you’re optimizing for gas costs on EIP-2930 compatible networks.

Advantages:

* Access Lists: Reduce gas costs by predefining storage and contract access.
* Backward Compatibility: Includes gasPrice, similar to legacy transactions.

Fields:

* `accessList`: Specifies addresses and storage keys to access.
* `gasPrice`: Static gas price (as in legacy transactions).

## EIP-1559 (Dynamic Fee Transaction)

A transaction pricing mechanism that includes fixed-per-block network fee that is burned and dynamically expands/contracts block sizes to deal with transient congestion.

When to Use:

* For most use cases on modern Ethereum networks.
* If the network supports EIP-1559 (introduced in the London upgrade).

Advantages:

* Efficient Fees: The maxFeePerGas ensures you don’t overpay for transactions.
* Better UX: The dynamic fee model provides more predictable fees.
* Network Compatibility: Widely supported on Ethereum Mainnet and most Layer 2 solutions.

Fields:

* `maxFeePerGas`: The maximum total fee per gas unit you’re willing to pay.
* `maxPriorityFeePerGas` (or GasTipCap): The tip (priority fee) for miners.
* `gasLimit`: The maximum gas allowed for the transaction.

### GasTipCap

GasTipCap (also referred to as `maxPriorityFeePerGas` in EIP-1559) is a parameter in Ethereum’s EIP-1559 transaction model that specifies the maximum tip (in Wei) a user is willing to pay per unit of gas to incentivize miners to prioritize their transaction.

Purpose of GasTipCap.

1. **Miner Incentive**: Miners receive the priority fee (tip) as direct compensation for including a transaction in a block.
2. **Transaction Priority**: Higher GasTipCap values increase the likelihood of a transaction being included in the next block, especially in congested conditions.
3. **User Control**: Lets users specify how much extra they are willing to pay beyond the mandatory base fee.

### GasFeeCap*

`GasFeeCap` (also referred to as maxFeePerGas in EIP-1559) is a parameter introduced in Ethereum’s London upgrade to define the maximum total fee per gas unit that a user is willing to pay for a transaction. This ensures that users don’t pay more than a specified amount per gas unit, even if the network fee (base fee) increases significantly during transaction execution. Can be thought of as a safety cap to control the maximum transaction cost. It is measured in `wei`.

When is GasFeeCap Used?

* In EIP-1559 dynamic fee transactions (Type 0x2 transactions).
* Replaces the old gasPrice field used in legacy transactions.

It is used in the dynamic fee model introduced by EIP-1559 to separate transaction fees into two components:

1. Base Fee: A mandatory fee burned by the network, dynamically adjusted based on network demand.
1. `GasTipCap` or Priority Fee (Tip): An optional fee paid to miners to prioritize the transaction.

**The Upper Limit** or the actual gas price paid per unit of gas is:

> `Effective Gas Price= min(GasFeeCap,Base Fee + GasTipCap)`

### Example of Gas Fee Calculation.

Assumptions:

* Base Fee = 50 Gwei
* GasTipCap (Priority Fee) = 2 Gwei
* GasFeeCap = 70 Gwei

Calculation:

1.	Base Fee + Tip: `Base Fee + GasTipCap = 50 + 2 = 52Gwei`
2.	Compare with GasFeeCap: `Effective Gas Price = min(GasFeeCap, Base Fee + GasTipCap) = min(70, 52) = 52Gwei`

The transaction will execute with an effective gas price of 52 Gwei.

## Geth implementation:

Package in `core/types`.

* [Transaction interface](https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction.go)
* [Legacy transactions](https://github.com/ethereum/go-ethereum/blob/master/core/types/tx_legacy.go)
* [Access List transactions](https://github.com/ethereum/go-ethereum/blob/master/core/types/tx_access_list.go)
* [Dynamic Fee transactions](https://github.com/ethereum/go-ethereum/blob/master/core/types/tx_dynamic_fee.go)

## References

* [The Most Misunderstood Concept in Ethereum | Ethereum transactions explained](https://www.youtube.com/watch?v=2EhKeQHFeTs)
* [Ethereum Transactions](https://www.youtube.com/watch?v=2QHGPH88WAI)
