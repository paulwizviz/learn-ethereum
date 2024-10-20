# Smart contracts

A "smart contract" is simply a program that runs on the Ethereum blockchain. It's a collection of code (its functions) and data (its state) that resides at a specific address on the Ethereum blockchain.
([source](https://ethereum.org/en/developers/docs/smart-contracts/))

This section discuss techniques for creating and deploying smart contracts.

## Solidity

* [Building Solidity compiler from source](https://docs.soliditylang.org/en/latest/installing-solidity.html#building-from-source)
* [Deploy Your First Smart Contract](https://www.web3.university/tracks/create-a-smart-contract/deploy-your-first-smart-contract)
* [Official Reference](https://soliditylang.org/)
* [Remix - solidity playground](https://remix.ethereum.org/)
* [Solidity source code](https://github.com/ethereum/solidity)
* [Smart Contract with Golang](https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9)

### Working examples

* Example 1 - This [Dockerfile](../build/solc/solc.dockerfile) demonstrates the steps to build solidity compiler from source code. To see the docker images in action run the command `./scripts/ops.sh image build:solc`.

* Example 2 - This [script](../scripts/solc.sh) demonstrate the compilation of a solidity file (see [./solidity/HelloWorld.sol](../solidity/HelloWorld.sol)). Running the script will produce `HelloWorld.abi` and `HelloWorld.bin`.