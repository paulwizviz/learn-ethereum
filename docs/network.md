# Setting up a node and network

Ethereum is a distributed network of computers (known as nodes) running software that can verify blocks and transaction data. The software application, known as a client, must be run on your computer to turn it into an Ethereum node ([source](https://ethereum.org/en/developers/docs/nodes-and-clients/)).

You can create a node that connects to:

* mainnet or testnet;
* a private network of nodes. 

## Docker base single node playground

You will find a Docker based environment containing `geth`, `puppeth` and related tools for you to experiment with. Run the following commands:

* `./scripts/gethclient.sh buid` to build the Docker image
* `./scripts/getclient.sh shell` to access an Unbuntu shell 

## Private networks

Private network are completely separate from Ethereum networks and something you and/or your team used. 

The steps involved in private networks are:

1) Create a genesis block (you can use a tool call Puppeth) and export file into a location where it can be accessed by the next stop.
2) Initialise the node with the genesis block `geth --datadir <location of where you want your node to hold data> <name and location of genesis file>`
3) Create a new account (to receive mining reward) `geth --datadir <location of your node data store> `
4) Start the node.

To learn more, please refer to these eductional materials:

* [Create a Private Ethereum Blockchain](https://gr0uch0dev.github.io/2020-01-15-privchain/)
* [Setting up Geth on macOS](https://www.youtube.com/watch?v=m_Ry0LZjoRQ&list=PLHQi9fG7tS9nq5McHxjIE2tCaQcCpnRHw)
* [Setting up Geth on Linux](https://www.youtube.com/watch?v=A5V2jdLi5mI&list=PLHQi9fG7tS9nq5McHxjIE2tCaQcCpnRHw)







