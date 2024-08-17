package main

import (
	"fmt"
	"time"
)

func main() {
	node := NewRollupNode()
	node.Start()

	// Simulate adding transactions
	go func() {
		for {
			time.Sleep(3 * time.Second)
			tx := NewTransaction(fmt.Sprintf("tx-%d", time.Now().Unix()), "data")
			node.Sequencer.AddTransaction(tx)
		}
	}()

	select {}
}
