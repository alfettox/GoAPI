package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	Number      string `json:"AccountNumber"`
	Balance     string `json:"Balance"`
	Description string `json:"AccountDescription"`
}

var Accounts []Account

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Giovanni's bank")
	fmt.Println("Endpoint: /")
}

func returnAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Accounts)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/accounts", returnAllAccounts)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	Accounts = []Account{
		{
			Number:      "C0001",
			Balance:     "24545.5",
			Description: "Checking Account",
		},
		{
			Number:      "C0002",
			Balance:     "4440.4",
			Description: "Saving Account",
		},
	}
	handleRequests()
}
