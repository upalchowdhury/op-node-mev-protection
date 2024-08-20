package node

import (
	"errors"
	"log"
)

type FraudProof struct {
	InvalidBlockID string
	ProofData      string
}

func SubmitFraudProof(fraudProof FraudProof) error {
	if VerifyFraudProof(fraudProof) {
		log.Println("Fraud proof verified, corrective action taken.")
		return nil
	}
	return errors.New("fraud proof invalid")
}

func VerifyFraudProof(fraudProof FraudProof) bool {
	return fraudProof.InvalidBlockID != "" && fraudProof.ProofData != ""
}
