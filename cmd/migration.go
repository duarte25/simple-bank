package main

import (
	"database/sql"
	"log"
)

func Migration(db *sql.DB) error {
	sqlStmt := `
	BEGIN;

	CREATE TABLE IF NOT EXISTS accounts
	(
		id uuid PRIMARY KEY,
		name varchar(255) NOT NULL,
		cpf varchar(14) NOT NULL UNIQUE,
		secret varchar(100) NOT NULL,
		balance int,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS transfers
	(
		id uuid PRIMARY KEY,
		account_origin_id uuid REFERENCES accounts,
		account_destination_id uuid REFERENCES accounts,
		amount int NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	
	COMMIT;`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}
