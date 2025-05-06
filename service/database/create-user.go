package database

import (
	"fmt"
)

// CreateUser crea un nuovo utente con il nome specificato e restituisce il sui ID
func (db *appdbimpl) CreateNewUser(name string) (int, error) {

	defaultPhoto := "/images/Screenshot_2025-05-06_alle_17.08.12-removebg-preview.png"
	res, err := db.c.Exec("INSERT INTO users_table (name, photo) VALUES (?,?)", name, defaultPhoto)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}

	return int(userId), nil
}
