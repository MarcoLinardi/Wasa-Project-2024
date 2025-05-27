package database

import (
	"database/sql"
	"fmt"
)

func (db *appdbimpl) getLastMessage(chatId int) (*LastMessage, error) {

	row := db.c.QueryRow(`SELECT content, timestamp, senderId 
							FROM messages
							WHERE chatId = ?
							ORDER BY timestamp DESC
							LIMIT 1;`,
		chatId)

	var lm LastMessage
	var senderID sql.NullInt64
	err := row.Scan(&lm.Content, &lm.Timestamp, &senderID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Nessun messaggio trovato per questa chat, è normale.
			return nil, nil
		}
		// Altro errore durante la query o lo scan
		return nil, fmt.Errorf("error scanning last message for chatID %d: %w", chatId, err)
	}
	if senderID.Valid {
		lm.SenderID = int(senderID.Int64)
	}
	return &lm, nil
}
