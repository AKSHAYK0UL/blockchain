package blockchain

import "github.com/AKSHAYK0UL/koulnetworkblockchain/constants"

func (bc *Blockchain) Mining() {
	NewTransaction(Transaction{From: constants.BLOCKCHAIN_NAME, To: bc.BlockchainAddress, Value: constants.MINING_REWARD})
	bc.CreateBlock()
}

//blockchainAddress address of the user
func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) uint64 {

	var totalAmount uint64
	totalAmount = 0
	for _, chain := range bc.Chain {
		for _, txn := range chain.Transactions {
			if blockchainAddress == txn.To {
				totalAmount += txn.Value

			}
			if blockchainAddress == txn.From {
				totalAmount -= txn.Value
			}
		}
	}
	return totalAmount
}
