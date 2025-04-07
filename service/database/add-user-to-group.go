package database

import "fmt"

func (db *appdbimpl) AddMemberToGroup(chatID, userID int) error {

	// Prima controllo che la chat sia un gruppo
	var isGroup bool
	err := db.c.QueryRow(`
		SELECT isGroup
		FROM chats_table
		WHERE chatId = ?
	`, chatID).Scan(&isGroup)
	if err != nil {
		return fmt.Errorf("error checking chat type: %w", err)
	}

	if !isGroup {
		return fmt.Errorf("chat is not a group")
	}

	// Aggiunge l'utente al gruppo
	_, err = db.c.Exec(`
		INSERT INTO chat_members (chatId, userId)
		VALUES (?, ?)
	`, chatID, userID)
	if err != nil {
		return fmt.Errorf("error adding member to group: %w", err)
	}

	return nil
}
