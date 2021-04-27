package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"local.crypto-go/client"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	})

	r.HandleFunc("/{value}", func(w http.ResponseWriter, r *http.Request) {
		// Get url value - value will be passed as a string to the FetchCrypto func
		v := mux.Vars(r)
		value := v["value"]
		crypto, err := client.FetchCrypto("USD", value)
		if err != nil {
			fmt.Fprintf(w, "Error - please enter new URL")
		}
		fmt.Fprintf(w, "%v", crypto)

	})

	log.Fatal(http.ListenAndServe(":9000", r))
}
