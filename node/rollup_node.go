package node

import (
	"log"
	"op-node/eas"
)

// RollupNode coordinates the Sequencer, Verifier, and Driver in the rollup node.
type RollupNode struct {
	Sequencer *Sequencer
	Verifier  *Verifier
	Driver    *RollupDriver
	EAS       *eas.EASManager
}

// NewRollupNode initializes a new RollupNode with the provided EASManager and RollupDriver.
func NewRollupNode(easManager *eas.EASManager, driver *RollupDriver) *RollupNode {
	seq := NewSequencer()
	verifier := NewVerifier()

	return &RollupNode{
		Sequencer: seq,
		Verifier:  verifier,
		Driver:    driver, // Already initialized RollupDriver
		EAS:       easManager,
	}
}

// Start begins running the RollupNode by starting the RollupDriver.
func (node *RollupNode) Start() {
	log.Println("Starting Rollup Node...")
	go node.Driver.Run()
}

// Stop gracefully shuts down the RollupNode.
func (node *RollupNode) Stop() {
	log.Println("Stopping Rollup Node...")
	node.Driver.Stop()
}
