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

	log.Printf("Adding time-locked transaction %s to sequencer (UnlockTime: %s)", tx.ID, unlockTime)
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
		log.Printf("Checking transaction %s for block inclusion. UnlockTime: %s, CurrentTime: %s",
			timeLockedTx.Transaction.ID, timeLockedTx.UnlockTime, now)

		if now.After(timeLockedTx.UnlockTime) && eas.VerifyAttestation(timeLockedTx.Transaction.ID) {
			log.Printf("Transaction %s is unlocked and valid. Adding to block.", timeLockedTx.Transaction.ID)
			validTxs = append(validTxs, timeLockedTx.Transaction)
		} else {
			log.Printf("Transaction %s is still locked or not valid.", timeLockedTx.Transaction.ID)
		}
	}

	if len(validTxs) == 0 {
		log.Println("No valid transactions for block creation.")
	} else {
		log.Printf("Creating new block with %d valid transactions", len(validTxs))
	}

	block := transactions.NewBlock(validTxs)
	seq.transactions = []*TimeLockedTransaction{} // Clear processed transactions
	return block
}
