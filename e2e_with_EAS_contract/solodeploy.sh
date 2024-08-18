ANVIL_ACCOUNT=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# $(grep -m 1 "Available Accounts" -A 2 anvil_output.log | tail -1 | awk '{print $2}')
ANVIL_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
# $(grep -m 1 "Private Keys" -A 2 anvil_output.log | tail -1 | awk '{print $2}')

echo "Anvil Account: $ANVIL_ACCOUNT"
echo "Anvil Private Key: $ANVIL_PRIVATE_KEY"

# Step 3: Export the private key for Foundry to use
export FOUNDRY_ETH_KEY=$ANVIL_PRIVATE_KEY

forge script ./DeployMockEAS.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --sender $ANVIL_ACCOUNT --private-key $ANVIL_PRIVATE_KEY