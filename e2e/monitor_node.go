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
		log.Fatalf("Failed to connect to local Ethereum client: %v", err)
	}

	// Mock EAS contract address deployed locally
	easManager := eas.NewEASManager(client, common.HexToAddress("0xYourMockEASContractAddress"))

	// Create the rollup node
	rollupNode := node.NewRollupNode(easManager)
	rollupNode.Start()

	// Simulate adding time-locked transactions with short unlock times
	go func() {
		for {
			time.Sleep(3 * time.Second)
			tx := transactions.NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data")

			unlockTime := time.Now().Add(5 * time.Second) // Fast unlock time for testing
			log.Printf("Simulating transaction with unlock time: %s", unlockTime)
			rollupNode.Sequencer.AddTransaction(tx, unlockTime)
		}
	}()

	// Monitor blocks and transactions
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Println("Monitoring block creation...")
			// Optionally print block details here
		}
	}()

	select {}
}
