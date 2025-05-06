package database

func (db *appdbimpl) getParticipants(chatId int) ([]User, error) {
	rows, err := db.c.Query(`
        SELECT u.userId, u.name, u.photo
        FROM chat_members cm
        JOIN users_table u ON cm.userId = u.userId
        WHERE cm.chatId = ?;`, chatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.Photo)
		if err != nil {
			return nil, err
		}
		participants = append(participants, user)
	}

	// Verifica se ci sono errori durante l'iterazione
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return participants, nil
}
