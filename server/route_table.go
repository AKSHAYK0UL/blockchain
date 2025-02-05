package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (ser *Server) Route() *http.ServeMux {
	serMux := http.NewServeMux()
	serMux.HandleFunc("GET /", HandleHome)
	serMux.HandleFunc("GET /chain", ser.GetBlockChain)
	return serMux
}

func (ser *Server) Run() {
	if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ser.Port)), ser.Route()); err != nil {
		fmt.Print(err.Error())
	}
}
