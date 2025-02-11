package blockchain

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/wallet"
)

type Blockchain struct {
	TransactionPool   []*Transaction `json:"transaction_pool"`
	Chain             []*Block       `json:"chain"`             //chain of block == blockchain
	BlockchainAddress string         `json:"blockchainaddress"` //address of the user in the blockchain network
	Port              uint16         `json:"port"`
}

// add transaction will add Transaction to the Transaction pool

func (bc *Blockchain) AddTransaction(SenderPublicKey *ecdsa.PublicKey, S *wallet.Signature, From, To string, Value uint64, Data []byte) bool {
	txn := NewTransaction(Transaction{From, To, Value, Data})

	//if the reward is being added to the user in that case the signature of the transaction will be ignored
	if txn.From == constants.BLOCKCHAIN_NAME {
		bc.TransactionPool = append(bc.TransactionPool, txn)
		return true
	}

	if !txn.VerifyTranaction(SenderPublicKey, S) {

		return false
	}
	// if bc.CalculateTotalAmount(From) < Value {
	// 	fmt.Println("ERROR : Not enough balance in the wallet")
	// 	return false

	// }

	bc.TransactionPool = append(bc.TransactionPool, txn)
	return true
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
func NewBlockChain(address string, port uint16) *Blockchain {
	bc := new(Blockchain)
	bc.BlockchainAddress = address
	bc.CreateGenesisBlock([]*Transaction{})
	return bc
}

// Print
func (bc *Blockchain) Print() {

	for idx, block := range bc.Chain {

		fmt.Printf("%s Block %d %s\n", strings.Repeat(".", 30), idx, strings.Repeat(".", 30))
		fmt.Printf("%s MINER DETAIL'S %s\n", strings.Repeat("-", 26), strings.Repeat("-", 27))

		fmt.Printf("BLOCKCHAIN ADDRESS         %s\n", constants.BLOCKCHAIN_NAME)
		fmt.Printf("MINER ADDRESS              %s\n", bc.BlockchainAddress)
		fmt.Printf("MINING REWARD              %d\n", constants.MINING_REWARD)
		fmt.Printf("%s\n", strings.Repeat("*", 69))
		block.Print()
		fmt.Printf("%s", strings.Repeat("\n", 2))
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

// tojson
func (bc *Blockchain) Tojson() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"blocks"`
	}{
		Blocks: bc.Chain,
	})
}
