package eas

import (
	"log"
	"op-node/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EASManager struct {
	client           *ethclient.Client
	ContractInstance *contract.Contract
}

// NewEASManager initializes a new EASManager and binds to the deployed contract.
func NewEASManager(client *ethclient.Client, contractAddress common.Address) *EASManager {
	contractInstance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind to deployed contract: %v", err)
	}

	return &EASManager{
		client:           client,
		ContractInstance: contractInstance,
	}
}

func (eas *EASManager) AttestTransaction(txHash string, encryptedData string, auth *bind.TransactOpts) (*types.Transaction, error) {
	log.Printf("Attesting transaction %s with encrypted data", txHash)

	tx, err := eas.ContractInstance.AttestTransactionWithPrivateData(auth, txHash, encryptedData)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
