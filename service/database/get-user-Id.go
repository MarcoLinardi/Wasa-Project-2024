package database

func (db *appdbimpl) GetUserIdByName(name string) (int, error) {
	var userId int
	err := db.c.QueryRow("SELECT userId FROM users_table WHERE name = ?", name).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
