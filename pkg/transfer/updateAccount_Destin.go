package transfer

import (
	"github.com/google/uuid"
)

func (transfer_handler *HandlerApi) UpdateAccount_Destin(amount int64, destinId uuid.UUID) {

	result, err := transfer_handler.DB.Query("SELECT balance from accounts where id = ?", destinId)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var account Account
	for result.Next() {
		err := result.Scan(&account.Balance)
		if err != nil {
			panic(err.Error())
		}
	}

	balanceDestin := account.Balance + amount

	_, err = transfer_handler.DB.Exec("UPDATE accounts SET balance = ($1) WHERE id = ($2)", balanceDestin, destinId)
	if err != nil {
		panic(err.Error())
	}
}
