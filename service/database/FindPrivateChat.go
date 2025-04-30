package database

import "database/sql"

func (db *appdbimpl) FindPrivateChat(user1ID, user2ID int) (int, error) {
	query := `
				SELECT c.chatId
				FROM chats_table c
				JOIN chat_members cm ON c.chatId = cm.chatId
				WHERE c.isGroup = 0
				AND cm.userId IN (?, ?)
				GROUP BY c.chatId
				HAVING COUNT(DISTINCT cm.userId) = 2
				LIMIT 1;
				`

	row := db.c.QueryRow(query, user1ID, user2ID)

	var chatID int
	err := row.Scan(&chatID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return chatID, nil
}
