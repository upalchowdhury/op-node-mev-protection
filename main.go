package main

import (
	"fmt"
	"log"
	"time"

	"op-node/eas"
	"op-node/node"
	"op-node/transactions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Instantiate the EAS manager with your EAS contract address
	easManager := eas.NewEASManager(client, common.HexToAddress("0xYourEASContractAddress"))

	// Create the rollup node
	rollupNode := node.NewRollupNode(easManager)
	rollupNode.Start()

	// Simulate adding time-locked transactions with very short unlock time
	go func() {
		for {
			time.Sleep(3 * time.Second) // Simulating transaction creation every 3 seconds
			tx := transactions.NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data")

			unlockTime := time.Now().Add(2 * time.Second) // Fast unlock time for testing
			log.Printf("Simulating transaction with unlock time: %s", unlockTime)
			rollupNode.Sequencer.AddTransaction(tx, unlockTime)
		}
	}()

	select {}
}
