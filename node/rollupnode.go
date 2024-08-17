package main

/*The RollupNode will be the main struct that controls the flow of the node.
It will interact with the Ethereum L1 chain,
manage block derivation, and communicate with the sequencer and verifier.
*/
import (
	"log"
)

type RollupNode struct {
	Sequencer *Sequencer
	Verifier  *Verifier
	Driver    *Driver
}

func NewRollupNode(sequencer *Sequencer, verifier *Verifier, driver *Driver) *RollupNode {
	return &RollupNode{
		Sequencer: sequencer,
		Verifier:  verifier,
		Driver:    driver,
	}
}

func (node *RollupNode) Start() {
	log.Println("Starting rollup node")
	go node.Driver.Start(node.Sequencer, node.Verifier)

}
