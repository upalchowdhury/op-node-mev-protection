package node

import (
	"log"
	"op-node/eas"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// RollupDriver handles the rollup process, including sequencing and attesting transactions.
type RollupDriver struct {
	Sequencer  *Sequencer
	EASManager *eas.EASManager
	stopCh     chan struct{}
}

// NewRollupDriver initializes a new RollupDriver with a Sequencer and EASManager.
func NewRollupDriver(seq *Sequencer, easManager *eas.EASManager) *RollupDriver {
	return &RollupDriver{
		Sequencer:  seq,
		EASManager: easManager,
		stopCh:     make(chan struct{}),
	}
}

// Run starts the RollupDriver's process loop, which continuously sequences and attests transactions.
func (driver *RollupDriver) Run() {
	log.Println("Running Rollup Driver...")

	for {
		select {
		case <-driver.stopCh:
			log.Println("Stopping Rollup Driver...")
			return
		default:
			driver.ProcessTransactions()
		}
	}
}

// ProcessTransactions handles transaction sequencing and attestation.
// func (driver *RollupDriver) ProcessTransactions() {
// 	tx := driver.Sequencer.NextTransaction()

// 	if tx != nil {
// 		log.Printf("Processing transaction %s with data: %s\n", tx.ID, tx.Data)

//			// Attest the transaction using the EASManager
//			auth := &bind.TransactOpts{} // Add the correct transaction options here
//			_, err := driver.EASManager.AttestTransaction(tx.ID, tx.Data, auth)
//			if err != nil {
//				log.Printf("Failed to attest transaction %s: %v", tx.ID, err)
//			}
//		} else {
//			log.Println("No transactions to process")
//		}
//	}
func (driver *RollupDriver) ProcessTransactions() {
	time.Sleep(2 * time.Second) // Add delay before processing to ensure transactions are available
	tx := driver.Sequencer.NextTransaction()

	if tx != nil {
		log.Printf("Processing transaction %s with data: %s\n", tx.ID, tx.Data)

		// Attest the transaction using the EASManager
		auth := &bind.TransactOpts{} // Ensure this is correctly initialized
		_, err := driver.EASManager.AttestTransaction(tx.ID, tx.Data, auth)
		if err != nil {
			log.Printf("Failed to attest transaction %s: %v", tx.ID, err)
		} else {
			log.Printf("Successfully attested transaction %s", tx.ID)
		}
	} else {
		log.Println("No transactions to process")
	}
}

// Stop signals the RollupDriver to stop processing.
func (driver *RollupDriver) Stop() {
	close(driver.stopCh)
}
