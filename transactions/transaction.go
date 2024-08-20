package transactions

type Transaction struct {
	ID              string
	Data            string
	PrivateDataHash string
	// Other fields as needed
}

// Example constructor
func NewTransaction(id, data string) *Transaction {
	return &Transaction{
		ID:   id,
		Data: data,
	}
}

func (tx *Transaction) GetID() string {
	return tx.ID
}
