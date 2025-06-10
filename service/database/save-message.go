package database

import (
	"fmt"
)

// SaveMessage salva un messaggio nel database e restituisce il suo ID
func (db *appdbimpl) SaveMessage(chatID int, msg Message) (int, error) {

	// Controlla se la chat esiste
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM chats_table WHERE chatId = ?)", chatID).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("errore nella verifica della chat: %w", err)
	}
	if !exists {
		return 0, fmt.Errorf("chat non trovata")
	}

	// Controlla che il mittente faccia parte della chat
	var isMember bool
	err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM chat_members WHERE chatId = ? AND userId = ?)", chatID, msg.SenderID).Scan(&isMember)
	if err != nil {
		return 0, fmt.Errorf("errore nella verifica dell'utente nella chat: %w", err)
	}
	if !isMember {
		return 0, fmt.Errorf("utente non autorizzato a inviare messaggi in questa chat")
	}

	// Inserisce il messaggio nel database
	result, err := db.c.Exec("INSERT INTO messages (chatId, senderId, content, photo, timestamp, replyToId) VALUES (?, ?, ?, ?, ?, ?)",
		chatID, msg.SenderID, msg.Content, msg.Photo, msg.Timestamp, msg.ReplyToID,
	)
	if err != nil {
		return 0, fmt.Errorf("errore nell'inserimento del messaggio: %w", err)
	}

	// Recupera l'ID del messaggio appena inserito
	messageID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("errore nel recupero dell'ID del messaggio: %w", err)
	}

	return int(messageID), nil
}
