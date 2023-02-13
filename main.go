package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"regexp"
)

type AddressType struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")


	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "missing address parameter", http.StatusBadRequest)
			return
		}

		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

		if (!re.MatchString(common.HexToAddress(address).String())) {
			http.Error(w, "invalid address parameter", http.StatusBadRequest)
			return
		}

		client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + apiKey)
		if err != nil {
			http.Error(w, "failed to connect to Ethereum client", http.StatusInternalServerError)
			return
		}

		code, err := client.CodeAt(r.Context(), common.HexToAddress(address), nil)
		if err != nil {
			http.Error(w, "failed to retrieve code from Ethereum client", http.StatusInternalServerError)
			return
		}

		var addressType AddressType
		addressType.Address = address
		if len(code) == 0 {
			addressType.Type = "EOA"
		} else {
			addressType.Type = "contract"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(addressType)
	})

	http.ListenAndServe(":8080", nil)
}
