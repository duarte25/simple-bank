package transfer

import "github.com/google/uuid"

func (transfer_handler *HandlerApi) Balance(amount int64, originId uuid.UUID) bool {
	result, err := transfer_handler.DB.Query("SELECT balance FROM accounts WHERE id = ?", originId)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var valueBalance Get_Balance
	for result.Next() {
		err := result.Scan(&valueBalance.Balance)
		if err != nil {
			panic(err.Error())
		}
	}

	higherBalance := false

	if amount > valueBalance.Balance {
		higherBalance = true
	}

	return higherBalance
}
