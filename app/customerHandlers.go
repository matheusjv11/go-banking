package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusjv11/go-banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, _ := ch.service.GetAllCustomer(status)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "applicaton/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := ch.service.GetCustomer(vars["customer_id"])

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
