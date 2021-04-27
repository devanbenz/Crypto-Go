package client

import (
	"encoding/json"
	"log"
	"net/http"

	"local.crypto-go/model"
)

func FetchCrypto(fiat string, crypto string) (string, error) {

	// URL for crypto - api key obtained at https://nomics.com
	URL := "https://api.nomics.com/v1/currencies/ticker?key=7be6ceace1819967e6628eaa14b16dc5&interval=1d&ids=" + crypto + "&convert=" + fiat

	// HTTP GET response via URL variable provided above
	resp, getErr := http.Get(URL)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// Defer closing of response body
	defer resp.Body.Close()

	var cryptRes model.CryptoResponse

	if err := json.NewDecoder(resp.Body).Decode(&cryptRes); err != nil {
		log.Fatal(err)
	}

	return cryptRes.TextOutput(), nil
}
