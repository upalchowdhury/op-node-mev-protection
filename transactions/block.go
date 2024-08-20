package transactions

type Block struct {
	ID           string
	Transactions []*Transaction
}

// NewBlock creates a new block from a list of transactions.
func NewBlock(txs []*Transaction) *Block {
	return &Block{
		ID:           "block-id", // You can generate a unique block ID here
		Transactions: txs,
	}
}
