package blockchain

import "github.com/AKSHAYK0UL/koulnetworkblockchain/constants"

func (bc *Blockchain) Mining() {
	NewTransaction(Transaction{From: constants.BLOCKCHAIN_NAME, To: bc.BlockchainAddress, Value: constants.MINING_REWARD})
	bc.CreateBlock()
}
