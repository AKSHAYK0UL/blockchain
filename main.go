package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/constants"
	"github.com/AKSHAYK0UL/koulnetworkblockchain/server"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

//	func printLog(data interface{}) {
//		log.Println(data)
//	}
func main() {

	// w, err := wallet.NewWallet()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Printf("PriavteKey %s\n", w.WalletPrivateKeyToString())
	// fmt.Printf("PublicKey %s\n", w.WalletPublicKeyToString())
	// fmt.Printf("Wallet address %s\n", w.WalletAddress())

	// wtxn := wallet.NewTransaction(w.PrivateKey, w.PublicKey, w.BlockchainAddress, "testuser", 17, []byte{})
	// ws, err := wtxn.GenerateSignature()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Printf("Signature %s ", ws)

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
	//-------------------05-02-2025----------------------------------

	//create wallets

	// MinerWallet, _ := wallet.NewWallet()
	// AlexWallet, _ := wallet.NewWallet()
	// BenWallet, _ := wallet.NewWallet()

	// wT := wallet.NewTransaction(AlexWallet.PrivateKey, AlexWallet.PublicKey, AlexWallet.BlockchainAddress, BenWallet.BlockchainAddress, 14, []byte{})

	// //create new blockchain

	// blkchn := blockchain.NewBlockChain(MinerWallet.BlockchainAddress)
	// //create a transaction
	// gS, err := wT.GenerateSignature()
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())
	// }
	// addTxn := blkchn.AddTransaction(AlexWallet.PublicKey, gS, AlexWallet.BlockchainAddress, BenWallet.BlockchainAddress, 14, []byte{})
	// blkchn.Mining()
	// fmt.Println("Transaction Added?", addTxn)
	// blkchn.Print()

	// fmt.Printf("Alex amount %d\n", blkchn.CalculateTotalAmount(AlexWallet.BlockchainAddress))
	// fmt.Printf("Ben amount %d\n", blkchn.CalculateTotalAmount(BenWallet.BlockchainAddress))
	// fmt.Printf("Miner amount %d\n", blkchn.CalculateTotalAmount(MinerWallet.BlockchainAddress))

	//-------------------06-02-2025----------------------------------
	// use this -> go run main.go -[port number] to change the port number
	port := flag.Uint("port", 8000, "TCP PORT NUMBER")
	flag.Parse()
	fmt.Println(*port)
	server := server.NewServer(uint16(*port))
	server.Run()

}
