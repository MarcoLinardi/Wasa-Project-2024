package database

import "fmt"

func (db *appdbimpl) RemoveMemberFromGroup(chatID int, userID int) error {
	_, err := db.c.Exec(
		`DELETE FROM chat_members WHERE chatId = ? AND userId = ?`,
		chatID, userID,
	)
	if err != nil {
		return fmt.Errorf("error removing member from group: %w", err)
	}
	return nil
}
