package database

import (
	"fmt"
)

// GetUsers recupera tutti gli utenti dal database
func (db *appdbimpl) GetUsers() ([]User, error) {

	// Esegui la query per ottenere tutti gli utenti
	rows, err := db.c.Query("SELECT userId, name, photo FROM users_table")
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	// Array per memorizzare gli utenti
	var users []User

	// Itera sui risultati della query
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Photo); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}

	// Controlla errori dopo aver iterato
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return users, nil
}
