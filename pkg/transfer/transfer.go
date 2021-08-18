package transfer

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func (transfer_handler *HandlerApi) Transfer(w http.ResponseWriter, r *http.Request) {
	var header = r.Header.Get("x-access-token") //Grab the token from the header
	header = strings.TrimSpace(header)

	tk := &Token{}
	_, _ = jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	Creds := &Transfer{}
	err := json.NewDecoder(r.Body).Decode(Creds)

	higherBalance := transfer_handler.Balance(Creds.Amount, tk.OriginId)
	if higherBalance != false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Exception{Message: "Account balance less than transaction"})
		return
	}

	ID := uuid.New().String()
	Data := time.Now()

	_, err = transfer_handler.DB.Exec("insert into transfers values ($1, $2, $3, $4, $5)", ID, tk.OriginId, Creds.Account_destination_id, Creds.Amount, Data)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transfer_handler.UpdateAccount_Origin(Creds.Amount, tk.OriginId)
	transfer_handler.UpdateAccount_Destin(Creds.Amount, Creds.Account_destination_id)
}
