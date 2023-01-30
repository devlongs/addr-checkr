package main

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AddressType struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

func main() {
	http.HandleFunc("/address", func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "missing address parameter", http.StatusBadRequest)
			return
		}

		client, err := ethclient.Dial("https://mainnet.infura.io/v3/sd")
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
