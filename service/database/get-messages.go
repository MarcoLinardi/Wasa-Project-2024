package database

import "fmt"

func (db *appdbimpl) GetAllMessages(chatID int) ([]Message, error) {
	rows, err := db.c.Query(`
		SELECT messageId, senderId, content, timestamp, status
		FROM messages
		WHERE chatId = ?
		ORDER BY timestamp ASC
	`, chatID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving messages with status: %w", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.MessageID, &msg.SenderID, &msg.Content, &msg.Timestamp, &msg.Status); err != nil {
			return nil, fmt.Errorf("error scanning message: %w", err)
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	return messages, nil
}
