package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Blockchain struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Chain           []*Block       `json:"chain"` //chain of block == blockchain
}

// add transaction will add Transaction to the Transaction pool

func (bc *Blockchain) AddTransaction(from, to string, value uint64, data []byte) error {
	txn := NewTransaction(Transaction{from, to, value, data})

	if !txn.VerifyTransaction() {
		return errors.New("invalid transaction")
	}

	bc.TransactionPool = append(bc.TransactionPool, txn)
	return nil
}

// create block in blockchain and empty the Transaction pool
func (bc *Blockchain) CreateBlock() *Block {
	prev_hash := bc.FindlastBlock().Hash //hash of the last block
	block := NewBlock(prev_hash, bc.ProofOfWork(), bc.TransactionPool)

	bc.Chain = append(bc.Chain, block)
	//empty the transactionPool
	bc.TransactionPool = []*Transaction{}
	return block
}

// create Genesis Block
func (bc *Blockchain) CreateGenesisBlock(txn []*Transaction) *Block {
	genesisBlock := &Block{
		PreviousHash: [32]byte{}, // No previous hash for Genesis Block
		TimeStamp:    time.Now().UnixNano(),
		Transactions: txn,
	}

	// Append the Genesis Block first
	bc.Chain = append(bc.Chain, genesisBlock)

	genesisBlock.Nonce = bc.ProofOfWork()
	genesisBlock.Hash = genesisBlock.GenerateHash()

	bc.Chain[0] = genesisBlock

	return genesisBlock
}

// create new block chain
func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateGenesisBlock([]*Transaction{})
	return bc
}

// Print
func (bc *Blockchain) Print() {

	for idx, block := range bc.Chain {

		fmt.Printf("%s Block %d %s\n", strings.Repeat(".", 30), idx, strings.Repeat(".", 30))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("#", 69))
}

// find the last block in the blockChain
func (bc *Blockchain) FindlastBlock() *Block {
	if len(bc.Chain) == 0 {
		return nil
	}
	return bc.Chain[len(bc.Chain)-1]
}

// copy transaction Pool
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	txns := make([]*Transaction, 0)
	for _, t := range bc.TransactionPool {
		txns = append(txns, NewTransaction(Transaction{From: t.From, To: t.To, Value: t.Value, Data: t.Data}))
	}
	return txns
}
