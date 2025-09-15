package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusjv11/go-banking/domain"
	"github.com/matheusjv11/go-banking/service"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	customerHandler := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// Defining routes
	router.HandleFunc("/customers", customerHandler.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getUserById).Methods(http.MethodGet)

	// starting server
	// With log.Fatal we can catch an error in the server start
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
