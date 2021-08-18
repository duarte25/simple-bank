package account

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (account_handler *HandlerApi) CreateAccount(w http.ResponseWriter, r *http.Request) {

	creds := &Account{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Secret), 8)
	valueBalance := 10000
	ID := uuid.New().String()

	Data := time.Now()
	_, err = account_handler.DB.Exec("insert into accounts values ($1, $2, $3, $4, $5, $6)", ID, creds.Name, creds.Cpf, string(hashedPassword), valueBalance, Data)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
