package database

import "fmt"

func (db *appdbimpl) GetReactions(messageID int) ([]Reaction, error) {
	rows, err := db.c.Query(`
		SELECT userId, reaction
		FROM message_reactions
		WHERE messageId = ?
	`, messageID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving reactions: %w", err)
	}
	defer rows.Close()

	var reactions []Reaction
	for rows.Next() {
		var r Reaction
		if err := rows.Scan(&r.UserID, &r.Reaction); err != nil {
			return nil, fmt.Errorf("error scanning reaction: %w", err)
		}
		reactions = append(reactions, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating reactions: %w", err)
	}

	return reactions, nil
}
