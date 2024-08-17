package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Sequencer struct {
	mu           sync.Mutex
	transactions []*Transaction
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		transactions: make([]*Transaction, 0),
	}
}

func (s *Sequencer) AddTransaction(tx *Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Printf("Adding transaction %s to sequencer", tx.ID)

	s.transactions = append(s.transactions, tx)
}

//	func (s *Sequencer) CreateBlock() *Block {
//		s.mu.Lock()
//		defer s.mu.Unlock()
//		log.Printf("Creating block with %d transactions", len(s.transactions))
//		block := NewBlock(s.transactions)
//		s.transactions = make([]*Transaction, 0)
//		return block
//	}

// ShuffleTransactions shuffles the transactions to add randomness
func (seq *Sequencer) ShuffleTransactions() {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(seq.transactions), func(i, j int) {
		seq.transactions[i], seq.transactions[j] = seq.transactions[j], seq.transactions[i]
	})
}

func (seq *Sequencer) CreateBlock() *Block {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	// Shuffle transactions for MEV protection
	seq.ShuffleTransactions()

	log.Printf("Creating new block with %d transactions", len(seq.transactions))
	block := NewBlock(seq.transactions)
	seq.transactions = []*Transaction{}
	return block
}
