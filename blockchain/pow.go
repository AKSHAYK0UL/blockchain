package blockchain

import (
	"fmt"
	"strings"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
)

// Valid Proof
func (bc *Blockchain) ValidProof(nonce int, timeStamp int64, prev_hash [32]byte, txns []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", constants.MINING_DIFFICULTY)
	Generateblock := Block{TimeStamp: timeStamp, PreviousHash: prev_hash, Nonce: nonce, Transactions: txns}
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
	for !bc.ValidProof(nonce, lastBlock.TimeStamp, lastBlock.Hash, txns, constants.MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}
