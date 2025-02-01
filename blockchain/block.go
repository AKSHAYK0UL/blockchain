package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Hash         [32]byte       `json:"hash"`
	PreviousHash [32]byte       `json:"previous_hash"`
	Nonce        int            `json:"nonce"`
	TimeStamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	// Transactions []string `json:"transactions"`
}

// create new block
func NewBlock(prev_hash [32]byte, nonce int, txn []*Transaction) *Block {
	block := &Block{
		PreviousHash: prev_hash,
		Nonce:        nonce,
		TimeStamp:    time.Now().UnixNano(),
		Transactions: txn,
	}

	block.Hash = block.GenerateHash() // Generate hash after setting all fields

	return block
}

// print block
func (b *Block) Print() {
	fmt.Printf("Timestamp                  %d\n", b.TimeStamp)
	fmt.Printf("Nonce                      %d\n", b.Nonce)
	fmt.Printf("Hash   	                   %x\n", b.Hash)
	fmt.Printf("Previous Hash              %x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.Print()
	}

}

// toJson
func (b *Block) toJson() (
	[]byte, error) {
	json, err := json.Marshal(b)
	if err != nil {
		return []byte{}, err
	}
	return json, nil
}

// Generate hash
func (b *Block) GenerateHash() [32]byte {
	byt, _ := b.toJson()
	return sha256.Sum256(byt)
}
