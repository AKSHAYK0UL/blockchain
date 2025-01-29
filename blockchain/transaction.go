package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int64  `json:"value"` //amount
	Data  []byte `json:"data"`  // will be used for smart contracts
}

// create new transaction
func NewTransaction(txn Transaction) *Transaction {
	t := new(Transaction)
	t.From = txn.From
	t.To = txn.To
	t.Value = txn.Value
	t.Data = txn.Data
	return t
}

// toJson
func (txn *Transaction) ToJson() (string, error) {
	json, err := json.Marshal(txn)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// print
func (txn *Transaction) Print() {
	fmt.Printf("%s TXN Data %s\n", strings.Repeat("-", 30), strings.Repeat("-", 30))
	fmt.Printf("From    	 %s\n", txn.From)
	fmt.Printf("To		 %s\n", txn.To)
	fmt.Printf("Value		 %d\n", txn.Value)
	fmt.Printf("Data		 %v\n", txn.Data)
	fmt.Printf("%s\n", strings.Repeat("*", 70))

}
