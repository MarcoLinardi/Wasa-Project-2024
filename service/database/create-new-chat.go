package database

import (
	"fmt"
)

// CreateChat inserisce una nuova chat nel database e restituisce il suo ID
func (db *appdbimpl) CreateChat(name string, users []int) (int, error) {
	result, err := db.c.Exec("INSERT INTO chats_table (name, userId) VALUES (?, ?)", name, users)
	if err != nil {
		return 0, fmt.Errorf("errore nell'inserimento della chat: %w", err)
	}

	chatID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("errore nel recupero dell'ID della chat: %w", err)
	}

	return int(chatID), nil
}
