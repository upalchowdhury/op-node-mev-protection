package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"op-node/eas"
	"op-node/node"
	"op-node/transactions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: go run monitor_node.go <MockEAS_Address> <Anvil_Account> <Anvil_Private_Key>")
	}

	mockEASAddress := os.Args[1]
	anvilAccount := os.Args[2]
	// anvilPrivateKey := os.Args[3]

	fmt.Println("MockEAS Address:", mockEASAddress)
	fmt.Println("Anvil Account:", anvilAccount)

	// Connect to the local Anvil node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to local Ethereum client: %v", err)
	}

	// Initialize the EAS manager with the deployed contract address
	easManager := eas.NewEASManager(client, common.HexToAddress(mockEASAddress))

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
