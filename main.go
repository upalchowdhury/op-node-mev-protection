package main

import (
	"fmt"
	"log"
	"time"

	"op-node/eas"
	"op-node/node"
	"op-node/transactions" // Correctly importing the transactions package

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Instantiate the EAS manager with your EAS contract address
	easManager := eas.NewEASManager(client, common.HexToAddress("0xYourEASContractAddress"))

	// Create the rollup node
	rollupNode := node.NewRollupNode(easManager)
	rollupNode.Start()

	// Simulate adding time-locked transactions
	go func() {
		for {
			time.Sleep(3 * time.Second)
			tx := transactions.NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data") // Call from transactions package

			unlockTime := time.Now().Add(10 * time.Second)
			rollupNode.Sequencer.AddTransaction(tx, unlockTime)
		}
	}()

	select {}
}
