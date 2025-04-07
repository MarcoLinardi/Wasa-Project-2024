package database

import "fmt"

func (db *appdbimpl) ForwardMessage(originalChatID, originalMsgID, destinationChatID, senderID int) error {

	var content string
	err := db.c.QueryRow(`
		SELECT content
		FROM messages
		WHERE chatId = ? AND messageId = ?
	`, originalChatID, originalMsgID).Scan(&content)
	if err != nil {
		return fmt.Errorf("error retrieving original message: %w", err)
	}

	_, err = db.c.Exec(`
		INSERT INTO messages (chatId, senderId, content, timestamp, status, isForwarded)
		VALUES (?, ?, ?, datetime('now'), 'sent', 1)
	`, destinationChatID, senderID, content)
	if err != nil {
		return fmt.Errorf("error forwarding message: %w", err)
	}

	return nil
}
