package wallet

import "crypto/ecdsa"

// Transaction struct represents a transaction data for sending coins from one wallet to another
type Transaction struct {
	SenderPrivateKey      *ecdsa.PrivateKey `json:"sender_private_key"`
	SenderPublicKey       *ecdsa.PublicKey  `json:"sender_public_key"`
	FromBlockChainAddress string            `json:"from_blockchain_address"`
	ToBlockChainAddress   string            `json:"to_blockchain_address"`
	Value                 int64             `json:"value"`
	Data                  []byte            `json:"data"` //for smart contracts
}

func NewTransaction(senderPrivateKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey, from, to string, value int64, data []byte) *Transaction {
	return &Transaction{
		SenderPrivateKey:      senderPrivateKey,
		SenderPublicKey:       senderPublicKey,
		FromBlockChainAddress: from,
		ToBlockChainAddress:   to,
		Value:                 value,
		Data:                  data,
	}

}
