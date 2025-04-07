package database

import "fmt"

func (db *appdbimpl) UpdateMessagesStatus(chatID int, messageIDs []int, newStatus string) error {
	tx, err := db.c.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	stmt, err := tx.Prepare("UPDATE messages SET status = ? WHERE chatId = ? AND messageId = ?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	for _, messageID := range messageIDs {
		_, err := stmt.Exec(newStatus, chatID, messageID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error updating message status: %w", err)
		}
	}

	return tx.Commit()
}
