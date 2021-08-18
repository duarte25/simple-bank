package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPerson mostra apenas uma conta
func (account_handler *HandlerApi) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := account_handler.DB.Query("SELECT name, cpf, balance FROM accounts WHERE id = ?", params["account_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Get_Balance
	for result.Next() {
		err := result.Scan(&post.Name, &post.Cpf, &post.Balance)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}
