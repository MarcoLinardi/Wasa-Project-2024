package database

import (
	"database/sql"
	"fmt"
)

// UserIdExist verifica se l'userId specificato esiste nella tabella degli utenti.
func (db *appdbimpl) UserIdExist(userId int) (bool, error) {

	// Query SQL per verificare l'esistenza
	query := "SELECT 1 FROM users_table WHERE id = ?"

	var exist int
	err := db.c.QueryRow(query, userId).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			// Caso in cui non ci sono righe, consideriamo che l'userId non esista
			return false, nil
		}
		return false, fmt.Errorf("error checking user existence: %w", err)
	}

	// Se il conteggio è maggiore di 0, l'userId esiste
	return exist > 0, nil
}
