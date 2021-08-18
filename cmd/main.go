package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/duarte25/simple-bank/auth"
	"github.com/duarte25/simple-bank/pkg/account"
	"github.com/duarte25/simple-bank/pkg/transfer"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// função principal para executar a api
func main() {

	// Abrindo banco de dados
	db, err := sql.Open("sqlite3", "./banking.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	error := Migration(db)
	if error != nil {
		log.Fatal(error)
	}

	account_handler := &account.HandlerApi{DB: db}
	transfer_handler := &transfer.HandlerApi{DB: db}

	// Servidor -------------------------------------------
	router := mux.NewRouter()

	publicV1 := router.PathPrefix("").Subrouter()
	AuthV1 := router.PathPrefix("").Subrouter()

	publicV1.HandleFunc("/accounts", account_handler.CreateAccount).Methods("POST")
	publicV1.HandleFunc("/login", account_handler.Login).Methods("POST")

	AuthV1.HandleFunc("/accounts", account_handler.GetAccount).Methods("GET")
	AuthV1.HandleFunc("/accounts/{account_id}/balance", account_handler.GetBalance).Methods("GET")
	AuthV1.HandleFunc("/transfer", transfer_handler.Transfer).Methods("POST")
	AuthV1.HandleFunc("/transfer", transfer_handler.ListTransfer).Methods("GET")
	AuthV1.Use(auth.Middleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}
