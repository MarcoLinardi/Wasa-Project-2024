package database

import (
	"fmt"
)

// CreateChat inserisce una nuova chat e associa gli utenti
func (db *appdbimpl) CreateChat(name string, users []int, isGroup bool, authenticatedUserID int) (int, error) {

	var photo string

	// Determino l'immagine da associare alla chat
	if isGroup {
		photo = "/images/default-group-avatar.png"
	} else {
		// Verifica che la chat privata abbia esattamente due utenti
		if len(users) != 2 {
			return 0, fmt.Errorf("chat privata con numero utenti non valido")
		}
		photo = ""
	}

	// Inserisco la chat nella tabella `chats_table`
	result, err := db.c.Exec("INSERT INTO chats_table (name, isGroup, photo) VALUES (?, ?, ?)", name, isGroup, photo)
	if err != nil {
		return 0, fmt.Errorf("errore nell'inserimento della chat: %w", err)
	}

	// Ottengo l'ID della chat appena creata
	chatID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("errore nel recupero dell'ID della chat: %w", err)
	}

	// Inserisco gli utenti nella tabella `chat_members`
	for _, userID := range users {
		_, err := db.c.Exec("INSERT INTO chat_members (chatId, userId) VALUES (?, ?)", chatID, userID)
		if err != nil {
			return 0, fmt.Errorf("errore nell'associare utenti alla chat: %w", err)
		}
	}

	return int(chatID), nil
}
