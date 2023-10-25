package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Account struct {
	Number      string `json:"AccountNumber"`
	Balance     string `json:"Balance"`
	Description string `json:"AccountDescription"`
	Name        string `json:"Name"`
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

func returnAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["number"]
	for _, account := range Accounts {
		if account.Number == key {
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) // 404 Not Found
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := io.ReadAll(r.Body)
    vars := mux.Vars(r)
    key := vars["number"]

    fmt.Println("Received PUT request for account:", key)
    fmt.Println("Request Body:", string(reqBody))

    for index, account := range Accounts {
        if account.Number == key {
            fmt.Println("Updating account:", key)
            err := json.Unmarshal(reqBody, &Accounts[index])
            if err != nil {
                fmt.Println("Error updating account:", err.Error())
                w.WriteHeader(http.StatusBadRequest) // 400 Bad Request
                return
            }
            fmt.Println("Account updated successfully!")
            w.WriteHeader(http.StatusOK) // 200 OK
            return
        }
    }
    fmt.Println("Account not found:", key)
    w.WriteHeader(http.StatusNotFound) // 404 Not Found
}


func createAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var account Account
	json.Unmarshal(reqBody, &account)
	Accounts = append(Accounts, account)
	json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["number"]
	found := false

	var updatedAccounts []Account

	for _, account := range Accounts {
		if account.Number == id {
			found = true
		} else {
			updatedAccounts = append(updatedAccounts, account)
		}
	}

	if found {
		Accounts = updatedAccounts
		w.WriteHeader(http.StatusNoContent) // 204 No Content
	} else {
		w.WriteHeader(http.StatusNotFound) // 404 Not Found
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/accounts", returnAllAccounts)

	router.HandleFunc("/account/{number}", returnAccount)
	router.HandleFunc("/account", createAccount).Methods("POST")
	router.HandleFunc("/account/{number}", updateAccount).Methods("PUT")
	router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":13001", router))
}

func main() {
	Accounts = []Account{
		{
			Number:      "C0001",
			Balance:     "24000.5",
			Description: "Checking Account",
			Name:        "Giovanni",
		},
		{
			Number:      "C0002",
			Balance:     "3000",
			Description: "Saving Account",
			Name:        "Andrea",
		},
		{
			Number:      "C0003",
			Balance:     "1200",
			Description: "Saving Account",
			Name:        "Rea",
		},
	}
	handleRequests()
}
