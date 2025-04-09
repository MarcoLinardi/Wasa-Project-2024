package database

import "fmt"

func (db *appdbimpl) SetGroupPhoto(chatID int, newPhoto string) error {

	// Controlliamo che sia un gruppo
	var isGroup bool
	err := db.c.QueryRow(`SELECT isGroup FROM chats_table WHERE chatId = ?`, chatID).Scan(&isGroup)
	if err != nil {
		return fmt.Errorf("error checking if chat is group: %w", err)
	}
	if !isGroup {
		return fmt.Errorf("cannot update photo: chat is not a group")
	}

	// Se è un gruppo, aggiorna la foto
	_, err = db.c.Exec(`UPDATE chats_table SET photo = ? WHERE chatId = ?`, newPhoto, chatID)
	if err != nil {
		return fmt.Errorf("error updating group photo: %w", err)
	}
	return nil
}
