package main

import (
	"fmt"
	"log"
	"time"

	"minimal-op-node/eas"
	"minimal-op-node/node"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/XwJgApR00bxnw_4STcEvXZxSkixPAcDI")
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
			tx := node.NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data")

			unlockTime := time.Now().Add(10 * time.Second)
			rollupNode.Sequencer.AddTransaction(tx, unlockTime)
		}
	}()

	select {}
}
