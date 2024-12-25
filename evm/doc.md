# Ethereum Virtual Machine (EVM)

EVM is a decentralised execution environment that execute compiled version of Smart Contract code written in Solidity, Vyper, etc. A "smart contract" is simply a program that runs on the Ethereum blockchain. It's a collection of code (its functions) and data (its state) that resides at a specific address on the Ethereum blockchain.

## Core Functions

* Smart Contract Execution:
    * The EVM executes smart contracts written in high-level programming languages like Solidity or Vyper, which are compiled into bytecode.
    * This bytecode runs deterministically across all nodes, ensuring consistency.
* State Management:
    * Ethereum’s state includes account balances, smart contract code, and storage values.
    * The state is updated through transactions, which are processed within the EVM.
* Gas and Resource Management:
    * To prevent abuse and ensure efficient use of resources, the EVM requires “gas” for execution.
    * Gas is consumed based on the computational complexity and storage needs of operations.

## Architecture

* Storage and Memory:
    * Storage: Persistent storage tied to a smart contract; data here is retained between transactions.
    * Memory: Temporary, volatile memory used during the execution of a transaction.
* Stack Machine:
    * The EVM operates as a stack-based architecture, with a last-in, first-out (LIFO) execution model.
    * Instructions manipulate the stack for computation.
* Opcode Instruction Set:
    * The EVM processes a set of low-level opcodes (e.g., ADD, MUL, PUSH, POP) that represent its computational tasks.

## Versions

The Ethereum Virtual Machine (EVM) evolves with each major Ethereum network upgrade (often called “hard forks”), and these upgrades introduce changes or enhancements to the EVM. Although we refer to the EVM as one “machine,” its behavior has been modified across different Ethereum versions. Here’s an overview of the evolution:

1.	Frontier (July 2015):
    * The first version of the EVM introduced with the launch of Ethereum.
	* Basic functionality for deploying and interacting with smart contracts.
2.	Homestead (March 2016):
	* Introduced new opcodes and improved stability.
	* Optimized gas costs for certain operations.
3.	Metropolis (Two Parts):
	* Byzantium (October 2017):
	    * Added privacy features (e.g., zk-SNARKs).
	    * Introduced the REVERT opcode for better error handling.
	    * Reduced gas costs for some operations.
	* Constantinople (February 2019):
	    * Optimized gas costs for storage and execution.
	    * Added new opcodes like SHR, SHL, and CREATE2.
4.	Istanbul (December 2019):
	* Improved compatibility with Zcash and other cryptographic techniques.
	* Added new opcodes like CHAINID and SELFBALANCE.
	* Adjusted gas costs for certain opcodes.
5.	Berlin (April 2021):
	* Optimized gas costs further.
	* Introduced new opcodes like BASEFEE to support Ethereum’s gas fee market reforms.
6.	London (August 2021):
	* Implemented EIP-1559, changing the gas fee mechanism.
	* Added the BASEFEE opcode to reflect the base fee of transactions.
7.	Shanghai (2023):
	* Focused on changes outside the core EVM but indirectly affected it by finalizing the move to proof-of-stake.
8.	Cancun (Upcoming):
    * Planned to include changes like EIP-4844 (“Proto-Danksharding”), which could impact the way data blobs are handled, though primarily focused on Layer 2 scaling.

## How the EVM Interacts with Accounts

1.	Transactions as Inputs:
	* Transactions originate from EOAs and can either transfer Ether or invoke code execution in a contract account.
	* The EVM processes these transactions and updates the state.
2.	Code Execution:
	* For contract accounts, the EVM reads the contract’s bytecode and executes it.
	* The execution may:
	    * Modify the contract’s storage.
	    * Transfer Ether between accounts.
	    * Interact with other contracts.
3.	Gas Management:
	* Transactions consume gas, paid in Ether from the sender’s account balance.
	* Gas limits ensure efficient use of network resources and prevent infinite loops in contract execution.
4.	State Updates:
	* The EVM updates the global state, which includes the balances, nonces, and storage of all accounts, after processing transactions.
5.	Account Independence:
	* EOAs are passive; they rely on transactions to interact with the EVM.
	* Contract accounts are active, as their behavior is dictated by their associated code executed within the EVM.

## EVM Related Source Code

1.	EVM Core Logic
    * Path: `core/vm/`
    * Files:
        * `evm.go`:
            * Defines the EVM’s core structure and provides the context for executing bytecode.
            * Includes methods for running transactions and managing gas.
        * `interpreter.go`:
            * Contains the bytecode interpreter that executes EVM opcodes.
            * Implements the main loop for processing instructions.
        * `instructions.go`:
            * Defines the EVM opcodes and their behaviors.
            * Maps opcodes to their execution logic.
        * `stack.go`:
            * Implements the stack data structure for the EVM’s stack-based execution model.
        * `memory.go`:
            * Implements the memory management for temporary data storage during execution.
        * `storage.go`:
            * Manages contract storage, which persists across transactions.
        * `gas_table.go`:
            * Defines the gas costs for EVM operations.

2.	Transaction Processing
    * Path: `core/`
    * Files:
        * `state_transition.go`:
            * Handles state transitions caused by transactions.
            * Updates balances, nonces, and storage based on transaction execution.
        * `evm_processor.go`:
            * Manages transaction execution in the context of the EVM.

3.	State Management
    * Path: `core/state/`
    * Files:
        * `state_object.go`:
            * Represents accounts and manages their associated data (e.g., storage, balance).
        * `statedb.go`:
            * Provides an interface for reading and writing state changes to the database.

4.	ABI Interaction
    * Path: `accounts/abi/`
    * Files:
        * `abi.go`:
            * Parses and encodes/decodes data for smart contract interaction.
        * `method.go`:
            * Manages the invocation of smart contract methods.

5.	Utilities
	* Path: `common/`
	* Files:
	    * `big.go`:
            * Utility functions for handling big integers used in EVM operations.
        * `math.go`:
            * Provides auxiliary mathematical functions for the EVM.

How these files work together:

* Transaction Flow:
    * A transaction is received and processed by the `state_transition.go` file.
    * The transaction invokes the EVM via `evm.go`.
    * The bytecode interpreter in interpreter.go executes the smart contract logic, leveraging `instructions.go`, `stack.go`, and `memory.go`.
* State Updates:
    * Changes to the global state (balances, storage, etc.) are managed by the `state_object.go` and `state_db.go` files.

## How to Extract EVM code and get it to run in Isolation

1.	Isolate the Core EVM Code:
	* Focus on the files in the core/vm/ directory:
	    * `evm.go`
	    * `interpreter.go`
	    * `instructions.go`
	    * `stack.go`
	    * `memory.go`
	    * `gas_table.go`
	* Extract these files into a separate Go module.

2.	Implement a Minimal Runtime Environment:
    * Create mock implementations or include simplified versions of:
        * State Management:
            * Mock `state_db.go` to manage account balances and contract storage.
        * Transaction Processing:
	        * Create dummy transactions to feed into the EVM.
        * Gas Management:
            * Simulate gas costs and limits.

3.	Write a Wrapper for Execution:
	* Build a minimal interface for running EVM bytecode:
	    * Load compiled bytecode (e.g., from a Solidity smart contract).
	    * Pass necessary inputs (e.g., calldata, sender address, gas limit).
	    * Execute the bytecode and capture the output or errors.

4.	Replace or Adapt Utility Functions:
	* Replace Geth-specific utility packages (e.g., common/) with equivalent libraries or your own implementations, especially for big integer math and cryptographic functions.

5.	Test the EVM:
	* Use known bytecode samples or Solidity contracts to test the isolated EVM.
	* Compare outputs with the Geth implementation to ensure correctness.

Here’s an example of how a minimal EVM execution might look:

1. Initialise the EVM:

```go
evm := vm.NewEVM(context, stateDB, config, chainConfig)
```

2. Load Bytecode:

```go
bytecode := []byte{0x60, 0x0a, 0x60, 0x0b, 0x01} // Example: PUSH10, PUSH11, ADD
```

3. Execute Bytecode:

```go
result, err := evm.Interpreter.Run(bytecode, nil, gasLimit)
if err != nil {
    log.Fatalf("EVM execution error: %v", err)
}
```

4.	Capture Results:

```go
fmt.Printf("Execution Result: %x\n", result)
```

## References

* [Ethereum Virtual Machine (EVM)](https://ethereum.org/en/developers/docs/evm/)
* [Opcodes](https://ethereum.org/en/developers/docs/evm/opcodes)
* [Ethereum EVM Illustrated](https://takenobu-hs.github.io/downloads/ethereum_evm_illustrated.pdf)