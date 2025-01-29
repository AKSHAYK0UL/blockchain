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

	blockchain := blockchain.CreateNewBlockChain() //create a new BlockChain with the genesis block
	// blockchain.Print()
	//Add transaction
	blockchain.AddTransaction("Alex", "Jason", 11, []byte{055, 41, 144})
	//create a new block and add it to the chain
	blockchain.CreateBlock(34)
	// blockchain.Print()

	// prev_hash = blockchain.FindlastBlock().GenerateHash()
	blockchain.AddTransaction("Ben", "Sam", 47, []byte{78, 41, 044})
	blockchain.AddTransaction("Smith", "Ali", 77, []byte{78, 41, 044})

	blockchain.CreateBlock(5)
	blockchain.Print()

}
