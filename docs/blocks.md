# Blocks

Blocks are batches of transactions with a hash of the previous block in the chain. This links blocks together (in a chain) because hashes are cryptographically derived from the block data. This prevents fraud, because one change in any block in history would invalidate all the following blocks as all subsequent hashes would change and everyone running the blockchain would notice. ([source](https://ethereum.org/en/developers/docs/blocks/))

## Genesis Block

* [Explaining the Genesis Block in Ethereum](https://arvanaghi.com/blog/explaining-the-genesis-block-in-ethereum/)

## Geth implementation

The aspects of blocks and blockchain implementations are found here:

```
go-ethereum
  + core
    + rawdb
      + database.go
      + ...
    + types
      + block.go
      + ...
    + blockchain.go
    + genesis.go
    + ...
```

* [Blockchain](https://github.com/ethereum/go-ethereum/blob/master/core/blockchain.go)
* [Block struct](https://github.com/ethereum/go-ethereum/blob/master/core/types/block.go)
* [Database](https://github.com/ethereum/go-ethereum/blob/master/core/rawdb/database.go)
* [Genesis struct](https://github.com/ethereum/go-ethereum/blob/master/core/genesis.go)

## References

* [How it works](https://www.youtube.com/watch?v=AjAinVsJUA0)