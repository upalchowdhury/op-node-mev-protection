package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID   string
	Data string
}

func NewTransaction(id, data string) *Transaction {
	return &Transaction{
		ID:   id,
		Data: data,
	}
}

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
