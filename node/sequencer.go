package main

import (
	"context"
	"log"
	"sync"
	"time"
)


type Sequencer struct{
	mu sync.Mutex
	transactions []*Transaction
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		transactions: make([]*Transaction, 0)
	}
}

func (s *Sequencer) AddTransaction(tx *Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Printf("Adding transaction %s to sequencer", tx.ID)

	s.transactions = append(s.transactions, tx)
}