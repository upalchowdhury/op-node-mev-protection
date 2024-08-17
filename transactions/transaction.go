package transactions

type Transaction struct {
	ID   string
	Data string
}

func NewTransaction(id, data string) *Transaction {
	return &Transaction{
		ID:   id,
		Data: data,
	}
}
