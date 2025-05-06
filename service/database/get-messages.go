package database

import "fmt"

func (db *appdbimpl) GetAllMessages(chatID int) ([]Message, error) {
	rows, err := db.c.Query(`
		SELECT messageId, senderId, content, timestamp, status, isForwarded
		FROM messages
		WHERE chatId = ?
		ORDER BY timestamp ASC
	`, chatID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving messages with status: %w", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.MessageID, &msg.SenderID, &msg.Content, &msg.Timestamp, &msg.Status, &msg.IsForwarded); err != nil {
			return nil, fmt.Errorf("error scanning message: %w", err)
		}

		// Recupera tutte le reactions per questo messaggio
		reactionRows, err := db.c.Query(
			`SELECT messageId, userId, reaction FROM message_reactions WHERE messageId = ?`,
			msg.MessageID,
		)
		if err != nil {
			return nil, fmt.Errorf("error retrieving reactions: %w", err)
		}

		var reactions []Reaction
		for reactionRows.Next() {
			var reaction Reaction
			if err := reactionRows.Scan(
				&reaction.MessageID, &reaction.UserID, &reaction.Reaction,
			); err != nil {
				reactionRows.Close()
				return nil, fmt.Errorf("error scanning reaction: %w", err)
			}
			reactions = append(reactions, reaction)
		}
		reactionRows.Close()

		// Verifica se ci sono errori durante l'iterazione
		if err := reactionRows.Err(); err != nil {
			return nil, fmt.Errorf("error iterating reactions: %w", err)
		}

		msg.Reactions = reactions
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	return messages, nil
}
