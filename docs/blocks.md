# Blocks

Blocks are batches of transactions with a hash of the previous block in the chain. This links blocks together (in a chain) because hashes are cryptographically derived from the block data. This prevents fraud, because one change in any block in history would invalidate all the following blocks as all subsequent hashes would change and everyone running the blockchain would notice. ([source](https://ethereum.org/en/developers/docs/blocks/))

## Genesis Block

* [Explaining the Genesis Block in Ethereum](https://arvanaghi.com/blog/explaining-the-genesis-block-in-ethereum/)

## Geth implementation:

* [Genesis struct](https://github.com/ethereum/go-ethereum/blob/afe344bcf31bfb477a6e1ad5b862a70fc5c1a22b/core/genesis.go#L49)
* [Block struct](https://github.com/ethereum/go-ethereum/blob/master/core/types/block.go)

## References

* [How it works](https://www.youtube.com/watch?v=AjAinVsJUA0)