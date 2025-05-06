package database

import (
	"fmt"
)

// GetChats recupera tutte le chat dell'utente dal database
func (db *appdbimpl) GetChats(userId int) ([]Chat, error) {

	// Esegui la query per ottenere tutte le chat a cui l'utente partecipa
	rows, err := db.c.Query(`SELECT ch.chatId, ch.name, ch.photo, ch.isGroup
							FROM chat_members cm
							JOIN chats_table ch ON cm.chatId = ch.chatId
							WHERE cm.userId = ?
							GROUP BY ch.chatId;`,
		userId)
	if err != nil {
		return nil, fmt.Errorf("error querying chats: %w", err)
	}
	defer rows.Close()

	// Array per memorizzare le chat
	var chats []Chat

	// Itera sui risultati della query
	for rows.Next() {
		var chat Chat
		if err := rows.Scan(&chat.ChatID, &chat.Name, &chat.Photo, &chat.IsGroup); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		participants, err := db.getParticipants(chat.ChatID)
		if err != nil {
			return nil, fmt.Errorf("error loading participants: %w", err)
		}
		chat.Members = participants

		chats = append(chats, chat)
	}

	// Controlla errori dopo aver iterato
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return chats, nil
}
