package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (account_handler *HandlerApi) Login(w http.ResponseWriter, r *http.Request) {

	creds := &Account{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := account_handler.DB.Query("SELECT id, secret FROM accounts WHERE cpf", creds.Cpf)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var account Account
	for result.Next() {
		err = result.Scan(&account.ID, &account.Secret)
		if err != nil {
			panic(err.Error())
		}
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(creds.Secret))
	if err != nil {
		return
	}

	expiresAt := time.Now().Add(time.Minute * 2).Unix()

	tk := &Token{
		Cpf:      creds.Cpf,
		OriginId: account.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["creds"] = creds

	json.NewEncoder(w).Encode(resp)

}
