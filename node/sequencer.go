package node

import (
	"log"
	"math/rand"
	"op-node/eas"
	"op-node/transactions"
	"sync"
	"time"
)

type TimeLockedTransaction struct {
	Transaction *transactions.Transaction
	UnlockTime  time.Time
}

type Sequencer struct {
	mu           sync.Mutex
	transactions []*TimeLockedTransaction
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		transactions: make([]*TimeLockedTransaction, 0),
	}
}

// Add transaction with a time lock
func (seq *Sequencer) AddTransaction(tx *transactions.Transaction, unlockTime time.Time) {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	log.Printf("Adding time-locked transaction %s to sequencer", tx.ID)
	seq.transactions = append(seq.transactions, &TimeLockedTransaction{
		Transaction: tx,
		UnlockTime:  unlockTime,
	})
}

// ShuffleTransactions shuffles the transactions to add randomness
func (seq *Sequencer) ShuffleTransactions() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(seq.transactions), func(i, j int) {
		seq.transactions[i], seq.transactions[j] = seq.transactions[j], seq.transactions[i]
	})
}

// Create block only with transactions that are unlocked
func (seq *Sequencer) CreateBlock(eas *eas.EASManager) *transactions.Block {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	validTxs := []*transactions.Transaction{}
	now := time.Now()

	// Filter out transactions that are unlocked
	for _, timeLockedTx := range seq.transactions {
		if now.After(timeLockedTx.UnlockTime) && eas.VerifyAttestation(timeLockedTx.Transaction.ID) {
			validTxs = append(validTxs, timeLockedTx.Transaction)
		}
	}

	log.Printf("Creating new block with %d valid transactions", len(validTxs))
	block := transactions.NewBlock(validTxs)
	seq.transactions = []*TimeLockedTransaction{} // Clear processed transactions
	return block
}
