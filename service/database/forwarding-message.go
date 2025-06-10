package database

import (
	"fmt"
	"time"
)

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

	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		fmt.Printf("Warning: Could not load timezone Europe/Rome: %v. Using UTC as fallback.\n", err)
		loc = time.UTC // Fallback a UTC se non riesce a caricare "Europe/Rome"
	}
	currentTime := time.Now().In(loc)
	formattedTimestamp := currentTime.Format("2006-01-02 15:04:05")

	_, err = db.c.Exec(`
		INSERT INTO messages (chatId, senderId, content, timestamp, status, isForwarded)
		VALUES (?, ?, ?, ?, 'sent', 1)
	`, destinationChatID, senderID, content, formattedTimestamp)
	if err != nil {
		return fmt.Errorf("error forwarding message: %w", err)
	}

	return nil
}
