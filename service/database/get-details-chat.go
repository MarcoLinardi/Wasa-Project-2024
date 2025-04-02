package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetChatDetails(chatID int, userID int) (*Chat, error) {

	// Verifica che l'utente sia membro della chat
	var tmp int
	check := `SELECT 1 FROM chat_members WHERE chat_id = ? AND user_id = ? LIMIT 1`
	err := db.c.QueryRow(check, chatID, userID).Scan(&tmp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("accesso non autorizzato alla chat")
		}
		return nil, err
	}

	// Recupera i dati della chat
	chat := &Chat{}
	query := `SELECT id, name, is_group FROM chats_table WHERE id = ?`
	err = db.c.QueryRow(query, chatID).Scan(&chat.ChatID, &chat.Name, &chat.IsGroup)
	if err != nil {
		return nil, err
	}

	// Recupera i messaggi della chat
	rows, err := db.c.Query(`
		SELECT id, sender_id, content, timestamp
		FROM messages
		WHERE chat_id = ?
		ORDER BY timestamp ASC
	`, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.MessageID, &msg.SenderID, &msg.Content, &msg.Timestamp); err != nil {
			return nil, err
		}
		chat.Messages = append(chat.Messages, msg)
	}

	return chat, nil
}
