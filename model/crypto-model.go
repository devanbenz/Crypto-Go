package model

import "fmt"

// CryptoResponse is a struct that will model the data recieved from the crypto API
type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

// Returns a function TextOutput() that will return a string with cryptoresponse info
func (c CryptoResponse) TextOutput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice : %s\nRank: %s\nHigh: %s\nCirculatingSupply: %s\n",
		c[0].Name, c[0].Price, c[0].Rank, c[0].High, c[0].CirculatingSupply)
	return p
}
