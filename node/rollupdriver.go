package main

import (
	"log"
	"time"
)

type RollupDriver struct{}

func NewRollupDriver() *RollupDriver {
	return &RollupDriver{}
}

func (driver *RollupDriver) Start(seq *Sequencer, ver *Verifier) {
	for {
		// Simulate block creation
		time.Sleep(5 * time.Second)
		block := seq.CreateBlock()

		// Verify block
		if ver.VerifyBlock(block) {
			log.Printf("Block %s is valid and added to the chain", block.ID)
		} else {
			log.Printf("Block %s is invalid!", block.ID)
		}
	}
}
