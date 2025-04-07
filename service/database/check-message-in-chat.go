package database

import "fmt"

func (db *appdbimpl) CheckMessageInChat(chatID, messageID int) (bool, error) {
	var exists bool
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM messages
			WHERE chatId = ? AND messageId = ?
		)
	`, chatID, messageID).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking if message belongs to chat: %w", err)
	}

	return exists, nil
}
