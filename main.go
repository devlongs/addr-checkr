package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type response struct {
	Type string `json:"type"`
}

func checkAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

	// Connect to Ethereum client
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ac7fe75399a146fe821d83ce4c7c512e")
	if err != nil {
		fmt.Println("Error connecting to Ethereum client:", err)
		return
	}

	// Get the address from the request parameters
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address not provided", http.StatusBadRequest)
		return
	}

	// Validate the address format using a regular expression
	validAddress := common.IsHexAddress(address)
	if !validAddress {
		response := response{Type: "invalid"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if the address is a smart contract or an EOA
	code, err := client.CodeAt(context.TODO(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println("Error checking code at address:", err)
		return
	}

	var res response
	if len(code) == 0 {
		res = response{Type: "EOA"}
	} else {
		res = response{Type: "Smart Contract"}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/check", checkAddress)
	fmt.Println(http.ListenAndServe(":8080", nil))

}
