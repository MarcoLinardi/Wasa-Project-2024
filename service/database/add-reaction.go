package database

import "fmt"

func (db *appdbimpl) AddReaction(messageID, userID int, reaction string) error {
	_, err := db.c.Exec(`
		INSERT INTO message_reactions (messageId, userId, reaction)
		VALUES (?, ?, ?)
		ON CONFLICT(messageId, userId) DO UPDATE SET reaction=excluded.reaction
	`, messageID, userID, reaction)
	if err != nil {
		return fmt.Errorf("error adding or updating reaction: %w", err)
	}
	return nil
}
