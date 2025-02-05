package server

import (
	"log"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/blockchain"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/wallet"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

func (ser *Server) GetChain() (*blockchain.Blockchain, error) {
	bc, ok := cache[constants.BLOCKCHAIN_NAME]
	if !ok {
		//if blockchain is not created
		//create miner's wallet
		minerWallet, err := wallet.NewWallet()
		if err != nil {
			log.Panicf("ERROR: In creating Miner wallet %v", err.Error())
			return nil, err
		}
		//create new blockchain
		bc = blockchain.NewBlockChain(minerWallet.BlockchainAddress, ser.Port)
		cache[constants.BLOCKCHAIN_NAME] = bc
	}
	return bc, nil

}
