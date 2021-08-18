package transfer

import (
	"database/sql"

	"github.com/duarte25/simple-bank/pkg/account"
	"github.com/google/uuid"
)

type HandlerApi struct {
	DB *sql.DB
}

type Account account.Account
type Token account.Token
type Transfer account.Transfer
type Exception account.Exception
type Get_Balance account.Get_Balance

type ListTransfer struct {
	Account_destination_id uuid.UUID `json:"destin" db:"account_destination_id"`
	Amount                 int64     `json:"amount" db:"amount"`
	CreatedAt              string    `json:"created_at" db:"created_at"`
}
