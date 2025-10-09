package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/matheusjv11/go-banking/domain"
	"github.com/matheusjv11/go-banking/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER_ADDRESS or SERVER_PORT not defined")
	}

	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" || os.Getenv("DB_ADDR") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("One or more database environment variables are not defined")
	}
}

func initEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}
}

func Start() {
	initEnvVars()
	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDb(dbClient)
	//accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)

	customerHandler := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	accountHandler := AccountHandlers{service.NewAccountService(accountRepositoryDB)}

	// Defining routes
	router.HandleFunc("/customers", customerHandler.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getUserById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accountHandler.MakeTransaction).Methods(http.MethodPost)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// With log.Fatal we can catch an error in the server start
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
