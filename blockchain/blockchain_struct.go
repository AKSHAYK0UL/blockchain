package blockchain

import (
	"fmt"
	"strings"
	"time"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
)

type Blockchain struct {
	TransactionPool []*Transaction `json:"transaction_pool"`
	Chain           []*Block       `json:"chain"` //chain of block == blockchain
}

// add transaction
func (bc *Blockchain) AddTransaction(from, to string, value int64, data []byte) {
	txn := NewTransaction(Transaction{from, to, value, data})
	bc.TransactionPool = append(bc.TransactionPool, txn)

}

// create block in blockchain
func (bc *Blockchain) CreateBlock(nonce int) *Block {
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
func CreateNewBlockChain() *Blockchain {
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

// Valid Proof
func (bc *Blockchain) ValidProof(nonce int, prev_hash [32]byte, txns []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", constants.MINING_DIFFICULTY)
	Generateblock := Block{TimeStamp: 0, PreviousHash: prev_hash, Nonce: nonce, Transactions: txns}
	checkNewBlockHash := fmt.Sprintf("%x", Generateblock.GenerateHash())
	fmt.Println(checkNewBlockHash)
	return checkNewBlockHash[:constants.MINING_DIFFICULTY] == zeros
}

// PoW [return the  valid nonce value]
func (bc *Blockchain) ProofOfWork() int {
	txns := bc.CopyTransactionPool()
	nonce := 0

	lastBlock := bc.FindlastBlock()
	if lastBlock == nil {
		// If there's no last block, return a default nonce
		return nonce
	}

	for !bc.ValidProof(nonce, lastBlock.Hash, txns, constants.MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}
