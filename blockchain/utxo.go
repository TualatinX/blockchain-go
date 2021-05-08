package blockchain

var (
	utxoPrefix   = []byte("utxo-")
	prefixLength = len(utxoPrefix)
)

// Unspent transaction outputs
type UTXOSet struct {
	BlockChain *BlockChain
}
