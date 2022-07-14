package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
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

	contentType := r.Header.Get("Content-Type")

	if contentType == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}
