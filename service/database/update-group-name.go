package database

import "fmt"

func (db *appdbimpl) SetGroupName(chatID int, newName string) error {

	// Prima controlliamo che sia effettivamente un gruppo
	var isGroup bool
	err := db.c.QueryRow(`SELECT isGroup FROM chats_table WHERE chatId = ?`, chatID).Scan(&isGroup)
	if err != nil {
		return fmt.Errorf("error checking if chat is group: %w", err)
	}
	if !isGroup {
		return fmt.Errorf("cannot rename: chat is not a group")
	}

	// Se è un gruppo, aggiorna il nome
	_, err = db.c.Exec(`UPDATE chats_table SET name = ? WHERE chatId = ?`, newName, chatID)
	if err != nil {
		return fmt.Errorf("error renaming group chat: %w", err)
	}
	return nil
}
