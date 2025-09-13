package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusjv11/go-banking/service"
)

// The "json" can map how the properties will look like in a json format
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "applicaton/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	w.Header().Add("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := ch.service.GetCustomer(vars["customer_id"])

	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
