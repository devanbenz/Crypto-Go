package main

import (
	"flag"
	"fmt"
	"log"

	"local.crypto-go/client"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency you would like to know the crypto in.",
	)

	cryptoName := flag.String(
		"crypto", "BTC", "The crypto currency you would like to analyze.",
	)
	flag.Parse()

	crypto, err := client.FetchCrypto(*fiatCurrency, *cryptoName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(crypto)
}
