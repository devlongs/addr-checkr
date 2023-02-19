// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/joho/godotenv"
// )

// type response struct {
// 	Type string `json:"type"`
// }

// var apiKey string

// func checkAddress(w http.ResponseWriter, r *http.Request) {
// 	// Connect to Ethereum client
// 	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + apiKey)
// 	if err != nil {
// 		fmt.Println("Error connecting to Ethereum client:", err)
// 		return
// 	}

// 	// Get the address from the request parameters
// 	address := r.URL.Query().Get("address")
// 	if address == "" {
// 		http.Error(w, "Address not provided", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate the address format using a regular expression
// 	validAddress := common.IsHexAddress(address)
// 	if !validAddress {
// 		response := response{Type: "invalid"}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Check if the address is a smart contract or an EOA
// 	code, err := client.CodeAt(context.TODO(), common.HexToAddress(address), nil)
// 	if err != nil {
// 		fmt.Println("Error checking code at address:", err)
// 		return
// 	}

// 	var res response
// 	if len(code) == 0 {
// 		res = response{Type: "EOA"}
// 	} else {
// 		res = response{Type: "Smart Contract"}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(res)
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Error loading .env file")
// 	}

// 	apiKey = os.Getenv("API_KEY")

// 	http.HandleFunc("/check", checkAddress)
// 	fmt.Println(http.ListenAndServe(":8080", nil))

// }
