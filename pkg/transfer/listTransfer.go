package transfer

import (
	"encoding/json"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func (transfer_handler *HandlerApi) ListTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var header = r.Header.Get("x-access-token") //Grab the token from the header
	header = strings.TrimSpace(header)

	tk := &Token{}
	_, _ = jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	result, err := transfer_handler.DB.Query("SELECT account_destination_id, amount, created_at FROM transfers WHERE account_origin_id = ?", tk.OriginId)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var posts []ListTransfer
	for result.Next() {
		var post ListTransfer
		err := result.Scan(&post.Account_destination_id, &post.Amount, &post.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)

	}
	json.NewEncoder(w).Encode(posts)
}
