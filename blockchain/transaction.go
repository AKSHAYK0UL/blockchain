package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/wallet"
)

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value uint64 `json:"value"` //amount
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
func (txn *Transaction) ToByteSlice() ([]byte, error) {
	byteSlice, err := json.Marshal(txn)
	if err != nil {
		return []byte{}, err
	}
	return byteSlice, nil
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

// verify transaction
func (txn Transaction) VerifyTranaction(senderPublicKey *ecdsa.PublicKey, s *wallet.Signature) bool {
	byteSlice, err := txn.ToByteSlice()
	if err != nil {

		return false
	}
	hash := sha256.Sum256(byteSlice)

	if txn.Value == 0 && !ecdsa.Verify(senderPublicKey, hash[:], s.R, s.S) { //check the signature
		return false

	}
	return true

}
