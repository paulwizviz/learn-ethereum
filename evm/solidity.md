# Solidity

Solidity is a high level language use to create Smart Contracts.

## Language Structure

### General keywords
    * `pragma` – compiler directives.
    * `import` – Includes external files.

### Data types

* Value types:
    * `uint`, `int` – Unsigned and signed integers.
    * `address` – Stores an Ethereum address.
    * `bool` – Boolean values (true or false).
    * `bytes` – Fixed-size or dynamically sized byte arrays.
    * `string` – UTF-8 encoded strings.

* Complex types:
    * `mapping` – Key-value pair storage.
    * `struct` – Custom-defined data structures.
    * `enum` – Enumerations.

### Visibility specifiers

* `public` – Accessible by everyone.
* `private` – Accessible only within the contract.
* `internal` – Accessible within the contract and derived contracts.
* `external` – Accessible only externally.

### State mutability specifiers

* `view` – Function does not modify the state.
* `pure` – Function neither reads nor modifies the state.
* `payable` – Function can receive Ether.

### Function modifiers

* `modifier` – Defines reusable code that modifies function behavior.

### Control structures

* `if`, `else` – Conditional statements.
* `while`, `do`, `for` – Loops.
* `break` – Exits the nearest enclosing loop.
* `continue` – Skips the current loop iteration.
* `return` – Returns a value from a function.

### Global variables and keywords

* `msg` – Contains information about the current call (e.g., msg.sender, msg.value).
* `tx` – Contains transaction-related data (e.g., tx.gasprice).
* `block` – Contains block-related data (e.g., block.timestamp, block.number).
* `this` – Refers to the current contract instance.
* `super` – Calls a parent contract’s function.

### Error handling

* `require` – Checks a condition and reverts on failure.
* `assert` – Checks conditions and reverts with a gas refund if false.
* `revert` – Aborts the transaction with an error message.
* `try`, `catch` – Handles external call exceptions.

### Contract lifecycle

* `contract` – Defines a new contract.
* `constructor` – Defines the initialization logic for a contract.
* `selfdestruct` – Destroys the contract and sends remaining Ether.

### Special keywords

* `event` – Defines events for logging.
* `emit` – Emits an event.
* `fallback` – Handles calls to non-existent functions.
* `receive` – Handles plain Ether transfers to the contract.

### Inheritance and interfaces

* `abstract` – Declares an abstract contract.
* `interface` – Declares an interface.
* `override` – Overrides a function from a parent contract.
* `virtual` – Allows a function to be overridden.

### Low-level constructs

* `delegatecall` – Executes code in the context of the calling contract.
* `call` – Low-level function call.
* `staticcall` – Executes a static call (read-only).

### Returns and return

* `returns`
    * Purpose: Specifies the type(s) of value(s) that a function will return.
    * Where It’s Used: In the function signature, after specifying the function parameters.
    * Example:
```solidity
function getSum(uint a, uint b) public pure returns (uint) {
    return a + b;
}
```
The `returns` (uint) part declares that this function will return a value of type uint.

* `return`
    * Purpose: Specifies the actual value(s) to return from a function during execution.
    * Where It’s Used: Inside the function body.
    * Example:
```solidity
function getSum(uint a, uint b) public pure returns (uint) {
    uint sum = a + b;
    return sum; // Specifies the value to return
}
```

## Compiler

`solc --abi --bin /opt/solidity/$sol -o /opt/abi --evm-version paris`

--evm-versions:

* byzantium
* constantinople
* istanbul
* berlin
* london
* paris (Pre-Shapella)
* shanghai or capella (Post-Shapella)

see [EVM](../evm/doc.md)

## Solidity and Go Types Equivalent

### Value types

| Solidity Type | Go Equivalent | Description |
| --- | --- | --- |
| bool  | bool | Boolean values (true or false). |
| uint, uint8-uint256 | uint8-uint64, big.Int | Unsigned integers. For large sizes like uint256, use big.Int in Go. |
| int, int8-int256 | int8-int64, big.Int | Signed integers. For large sizes like int256, use big.Int. |
| address | [20]byte or common.Address | Ethereum addresses (20-byte fixed). |
| bytes1-bytes32 | [N]byte | Fixed-size byte arrays in Go. Use [32]byte for bytes32, etc. |
| byte | byte (alias for uint8) | Single byte. |
| bytes | []byte | Dynamically-sized byte array. |
| string | string | UTF-8 encoded strings. |

### Reference types

| Solidity Type | Go Equivalent | Description |
| --- | --- | --- |
| uint[], int[], etc. | []uint, []int, etc.  | Dynamic arrays. |
| uint[10], int[10] | [10]uint, [10]int, etc. | Fixed-size arrays. |
| mapping(keyType => valueType) | map[keyType]valueType  | Mappings can be implemented using Go’s map.  |
| struct | struct | Custom user-defined types. |

### Special types

| Solidity Type | Go Equivalent | Description |
| --- | --- | --- |
| address payable | [20]byte or common.Address | Ethereum addresses that can receive Ether. |
| enum | const + iota or int | Represented by constants with iota or simple integers in Go. |
| tuple | struct or []interface{} | Tuples can be represented as structs or arrays of interfaces. |

### Struct equivalent

Solidity: 

```solidity
struct User {
    string name;
    uint age;
    address wallet;
}
``` 

Go:

```go
type User struct {
    Name   string
    Age    uint
    Wallet common.Address
}
```

## References

* [Building Solidity compiler from source](https://docs.soliditylang.org/en/latest/installing-solidity.html#building-from-source)
* [Deploy Your First Smart Contract](https://www.web3.university/tracks/create-a-smart-contract/deploy-your-first-smart-contract)
* [Introduction to smart contracts](https://ethereum.org/en/developers/docs/smart-contracts/)
* [Official Reference](https://soliditylang.org/)
* [Remix - solidity playground](https://remix.ethereum.org/)
* [Solidity source code](https://github.com/ethereum/solidity)
* [Smart Contract with Golang](https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9)

