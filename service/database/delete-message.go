package database

import "fmt"

func (db *appdbimpl) DeleteMessage(messageID int) error {
	_, err := db.c.Exec(`
		DELETE FROM messages
		WHERE messageId = ?
	`, messageID)
	if err != nil {
		return fmt.Errorf("error deleting message: %w", err)
	}
	return nil
}
