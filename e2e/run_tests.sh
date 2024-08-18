#!/bin/bash

# Step 1: Start Ganache
ganache-cli -p 8545 -m "test test test test test test test test test test test junk" --accounts 10 &
GANACHE_PID=$!
echo "Ganache running with PID $GANACHE_PID"

# Step 2: Deploy Mock EAS Contract using Hardhat
echo "Deploying Mock EAS contract..."
npx hardhat run deploy.js --network ganache

# Capture the deployed contract address from Hardhat logs
MOCK_EAS_ADDRESS="0xYourMockEASContractAddress" # Update this dynamically if necessary

# Step 3: Run the Go Rollup Node with Mock EAS
echo "Running Rollup Node with EAS integration..."
go run monitor_node.go

# Step 4: Cleanup
kill $GANACHE_PID
echo "Ganache stopped."
