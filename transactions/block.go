package transactions

import (
	"fmt"
	"time"
)

type Block struct {
	ID           string
	Transactions []*Transaction
}

func NewBlock(txs []*Transaction) *Block {
	return &Block{
		ID:           fmt.Sprintf("block-%d", time.Now().Unix()),
		Transactions: txs,
	}
}
