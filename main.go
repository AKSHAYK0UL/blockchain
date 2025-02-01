package main

import (
	"fmt"
	"log"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/wallet"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

//	func printLog(data interface{}) {
//		log.Println(data)
//	}
func main() {

	wallet, err := wallet.NewWallet()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("PriavteKey %s\n", wallet.WalletPrivateKeyToString())
	fmt.Printf("PublicKey %s\n", wallet.WalletPublicKeyToString())
	fmt.Printf("Wallet address %s\n", wallet.WalletAddress())

	// blockchain := blockchain.NewBlockChain("0x0y87dbdj9dnd4jb8") //create a new BlockChain with the genesis block and user address
	// //Add transaction
	// blockchain.AddTransaction("Alex", "Jason", 11, []byte{055, 41, 144})
	// //create a new block and add it to the chain

	// blockchain.Mining()
	// blockchain.AddTransaction("Ben", "Sam", 15, []byte{78, 41, 044})
	// blockchain.AddTransaction("Smith", "Ali", 77, []byte{78, 41, 044})

	// blockchain.Mining()

	// blockchain.Print()

	// fmt.Printf("Total amount of User Jason %d\n", blockchain.CalculateTotalAmount("Jason"))

	// fmt.Printf("Total amount of User Ali %d\n", blockchain.CalculateTotalAmount("Ali"))

}
