package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:9000", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{Name: "Forest", City: "ChengDu", ZipCode: "610000"},
		{Name: "Aspish", City: "New Delhi", ZipCode: "11075"},
		{Name: "Rob", City: "New Delhi", ZipCode: "11075"},
	}

	w.Header().Add("Content-Type", "Application/json")

	json.NewEncoder(w).Encode(customer)
}
