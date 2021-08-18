package account

import (
	"database/sql"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type HandlerApi struct {
	DB *sql.DB
}

type Account struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Balance int64     `json:"balance" db:"balance"`
	Name    string    `json:"name" db:"name"`
	Cpf     string    `json:"cpf" db:"cpf"`
	Secret  string    `json:"secret" db:"secret"`
}
type Token struct {
	Cpf      string    `json:"cpf" db:"cpf"`
	OriginId uuid.UUID `json:"originid" db:"account_origin_id"`
	*jwt.StandardClaims
}
type Transfer struct {
	ID                     uuid.UUID `json:"id" db:"id"`
	OriginId               uuid.UUID `json:"originid" db:"account_origin_id"`
	Amount                 int64     `json:"amount" db:"amount"`
	Account_destination_id uuid.UUID `json:"destin" db:"account_destination_id"`
}
type Exception struct {
	Message string `json:"message"`
}

type Get_Balance struct {
	Balance int64  `json:"balance" db:"balance"`
	Name    string `json:"name" db:"name"`
	Cpf     string `json:"cpf" db:"cpf"`
}

type ListTransfer struct {
	Account_destination_id uuid.UUID `json:"destin" db:"account_destination_id"`
	Amount                 int64     `json:"amount" db:"amount"`
}
