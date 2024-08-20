package node

import (
	"log"
	"op-node/transactions"
	"sync"
)

type Sequencer struct {
	mu           sync.Mutex
	transactions []*transactions.Transaction
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		transactions: make([]*transactions.Transaction, 0),
	}
}

func (seq *Sequencer) AddTransaction(tx *transactions.Transaction) {
	seq.mu.Lock()
	defer seq.mu.Unlock()
	seq.transactions = append(seq.transactions, tx)
	log.Printf("Transaction %s added to sequencer. Pool size: %d", tx.ID, len(seq.transactions)) // Log transaction addition
}

func (seq *Sequencer) NextTransaction() *transactions.Transaction {
	seq.mu.Lock()
	defer seq.mu.Unlock()
	if len(seq.transactions) == 0 {
		log.Println("No transactions available in sequencer")
		return nil
	}
	tx := seq.transactions[0]
	seq.transactions = seq.transactions[1:]
	log.Printf("Transaction %s retrieved from sequencer. Pool size: %d", tx.ID, len(seq.transactions)) // Log transaction retrieval
	return tx
}

func (seq *Sequencer) GetTransactions() []*transactions.Transaction {
	seq.mu.Lock()
	defer seq.mu.Unlock()
	return seq.transactions
}
