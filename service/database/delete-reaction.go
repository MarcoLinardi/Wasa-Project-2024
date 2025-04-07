package database

import "fmt"

func (db *appdbimpl) DeleteReaction(messageID, userID int) error {
	_, err := db.c.Exec(`
		DELETE FROM message_reactions
		WHERE messageId = ? AND userId = ?
	`, messageID, userID)
	if err != nil {
		return fmt.Errorf("error deleting reaction: %w", err)
	}
	return nil
}
