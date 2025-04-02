package database

import "fmt"

func (db *appdbimpl) DeleteChat(chatID int) error {

	result, err := db.c.Exec(`DELETE FROM chats WHERE chat_id = ?`, chatID)
	if err != nil {
		return fmt.Errorf("errore durante l'eliminazione della chat: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("errore ottenendo le righe modificate: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nessuna chat trovata con ID %d", chatID)
	}

	return nil
}
