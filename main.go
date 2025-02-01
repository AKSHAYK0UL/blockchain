package main

import (
	"log"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/blockchain"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

//	func printLog(data interface{}) {
//		log.Println(data)
//	}
func main() {

	blockchain := blockchain.NewBlockChain("MINER") //create a new BlockChain with the genesis block and user address
	// blockchain.Print()
	//Add transaction
	blockchain.AddTransaction("Alex", "Jason", 11, []byte{055, 41, 144})
	//create a new block and add it to the chain
	// blockchain.CreateBlock()
	// blockchain.Print()

	blockchain.Mining()
	// prev_hash = blockchain.FindlastBlock().GenerateHash()
	blockchain.AddTransaction("Ben", "Sam", 15, []byte{78, 41, 044})
	blockchain.AddTransaction("Smith", "Ali", 77, []byte{78, 41, 044})

	// blockchain.CreateBlock()
	blockchain.Mining()

	blockchain.Print()

}
