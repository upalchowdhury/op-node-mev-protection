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
		log.Fatalf("Usage: go run main.go <MockEAS_Address> <Anvil_Account> <Anvil_Private_Key>")
	}

	mockEASAddress := os.Args[1]
	anvilAccount := os.Args[2]
	anvilPrivateKey := os.Args[3]

	fmt.Println("MockEAS Address:", mockEASAddress)
	fmt.Println("Anvil Account:", anvilAccount)

	// Connect to the local Anvil node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to local Ethereum client: %v", err)
	}

	// Initialize the EAS manager with the deployed contract address
	easManager := eas.NewEASManager(client, common.HexToAddress(mockEASAddress))

	// Initialize the Sequencer
	sequencer := node.NewSequencer()

	// Initialize the RollupDriver with the Sequencer and EASManager
	rollupDriver := node.NewRollupDriver(sequencer, easManager)

	// Create the RollupNode with the initialized EASManager and RollupDriver
	rollupNode := node.NewRollupNode(easManager, rollupDriver)
	rollupNode.Start()

	// Simulate adding transactions (without the unlockTime)
	go func() {
		for {
			time.Sleep(3 * time.Second)
			tx := transactions.NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data")

			// Simulate adding the transaction to the sequencer
			log.Printf("Adding transaction: %s", tx.ID)
			rollupNode.Sequencer.AddTransaction(tx)

			// Log the number of transactions in the sequencer after adding
			log.Printf("Current transaction pool size: %d", len(rollupNode.Sequencer.GetTransactions()))
		}
	}()

	// Monitor blocks and transactions
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Println("Monitoring block creation...")
		}
	}()

	select {}
}
