package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetChatDetails(chatID int, userID int) (*Chat, error) {

	// Verifica che l'utente sia membro della chat
	var tmp int
	check := `SELECT 1 FROM chat_members WHERE chatId = ? AND userId = ? LIMIT 1`
	err := db.c.QueryRow(check, chatID, userID).Scan(&tmp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("accesso non autorizzato alla chat")
		}
		return nil, err
	}

	// Recupera i dati della chat
	chat := &Chat{}
	query := `SELECT chatId, name, photo, isGroup FROM chats_table WHERE chatId = ?`
	err = db.c.QueryRow(query, chatID).Scan(&chat.ChatID, &chat.Name, &chat.Photo, &chat.IsGroup)
	if err != nil {
		return nil, err
	}

	// Recupera i membri della chat
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.photo 
							FROM chat_members cm JOIN users_table u ON cm.userId = u.userId 
							WHERE cm.chatId = ?`, chatID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Photo); err != nil {
			return nil, err
		}
		chat.Members = append(chat.Members, user)
	}

	return chat, nil

}
