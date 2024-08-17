package node

import (
	"log"
	"op-node/eas"
	"time"
)

type RollupDriver struct{}

func NewRollupDriver() *RollupDriver {
	return &RollupDriver{}
}

func (driver *RollupDriver) Start(seq *Sequencer, ver *Verifier, eas *eas.EASManager) {
	for {
		// Simulate block creation
		time.Sleep(5 * time.Second)

		// Attest each transaction in the sequencer
		for _, tx := range seq.transactions {
			eas.AttestTransaction(tx.Transaction.ID)
		}

		block := seq.CreateBlock(eas)

		// Verify block
		if ver.VerifyBlock(block) {
			log.Printf("Block %s is valid and added to the chain", block.ID)
		} else {
			log.Printf("Block %s is invalid!", block.ID)
		}
	}
}
