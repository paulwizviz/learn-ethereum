# Nodes, Clients and Networks

This section discusses client, node and network.

## Client and node

A client is an implementation of Ethereum that verifies data against the protocol rules and keeps the network secure. A client is an implementation of Ethereum that verifies data against the protocol rules and keeps the network secure. A node has to run two clients: a consensus client and an execution client (see [Nodes and Clients](https://ethereum.org/en/developers/docs/nodes-and-clients/)).

### Working example

Ethereum is a distributed network of computers (known as nodes) running software that can verify blocks and transaction data.

There are several types of networks. Refer to this [list of networks](https://ethereum.org/en/developers/docs/networks/).

## Local network

Included in this project are docker-compose based networks. You will find a network centred on Geth nodes.

### Geth based network

Here is the project's local [Geth base network](../deployments/gethnet.yaml).

The network is still under development.

## References

* [Official Geth documentation -- private network](https://geth.ethereum.org/docs/fundamentals/private-network)
* [Build your Own Private Geth PoA Ethereum Network (Blockholic)](https://www.youtube.com/watch?v=pz7-JGG6T2Y&list=PLkM0MH7Grb25poKEiId5pEQg-OzLQRNM4)