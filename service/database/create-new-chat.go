package database

import (
	"fmt"
)

// CreateChat inserisce una nuova chat e associa gli utenti
func (db *appdbimpl) CreateChat(name string, users []int) (int, error) {

	// Inserisco la chat nella tabella `chats_table`
	result, err := db.c.Exec("INSERT INTO chats_table (name) VALUES (?)", name)
	if err != nil {
		return 0, fmt.Errorf("errore nell'inserimento della chat: %w", err)
	}

	// Ottengo l'ID della chat appena creata
	chatID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("errore nel recupero dell'ID della chat: %w", err)
	}

	// Inserisco gli utenti nella tabella `chat_users`
	for _, userID := range users {
		_, err := db.c.Exec("INSERT INTO chat_members (chat_id, user_id) VALUES (?, ?)", chatID, userID)
		if err != nil {
			return 0, fmt.Errorf("errore nell'associare utenti alla chat: %w", err)
		}
	}

	return int(chatID), nil
}
