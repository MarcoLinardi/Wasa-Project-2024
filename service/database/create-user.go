package database

import (
	"fmt"
)

// CreateUser crea un nuovo utente con il nome specificato e restituisce il sui ID
func (db *appdbimpl) CreateUser(name string) (int, error) {

	res, err := db.c.Exec("INSERT INTO users_table (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("error crating user: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}

	return int(lastId), nil
}
