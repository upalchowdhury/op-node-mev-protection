package node

import (
	"log"
	"op-node/eas"
)

type RollupNode struct {
	Sequencer *Sequencer
	Verifier  *Verifier
	Driver    *RollupDriver
	EAS       *eas.EASManager
}

func NewRollupNode(easManager *eas.EASManager) *RollupNode {
	return &RollupNode{
		Sequencer: NewSequencer(),
		Verifier:  NewVerifier(),
		Driver:    NewRollupDriver(),
		EAS:       easManager,
	}
}

func (node *RollupNode) Start() {
	log.Println("Starting Rollup Node...")
	go node.Driver.Start(node.Sequencer, node.Verifier, node.EAS)
}
