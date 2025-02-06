package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AKSHAYK0UL/koulnetworkblockchain/wallet"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Invalid request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Hello world!\nBlockChain Server")
}

func (ser *Server) GetBlockChain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		blkchn, err := ser.GetChain()
		if err != nil {
			log.Printf("Error: %v", err.Error())
			return
		}
		blkchn.Print()
		bs, _ := blkchn.Tojson()

		json.NewEncoder(w).Encode(bs)

		return

	default:
		fmt.Fprintf(w, "Invalid request")
	}
}

func (ser *Server) CreateWallet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		wallet, err := wallet.NewWallet()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bs, err := wallet.ToJson()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(bs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
