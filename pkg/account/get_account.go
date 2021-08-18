package account

import (
	"encoding/json"
	"net/http"
)

// Mostra todas as contas -----------------------------------------------
func (account_handler *HandlerApi) GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []Account
	result, err := account_handler.DB.Query("SELECT id, name, cpf,balance from accounts")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Account
		err := result.Scan(&post.ID, &post.Name, &post.Cpf, &post.Balance)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)

	}
	json.NewEncoder(w).Encode(posts)

}
