package database

func (db *appdbimpl) GetUserById(userID int) (User, error) {
	var user User
	err := db.c.QueryRow("SELECT userId, name, photo FROM users_table WHERE userId = ?", userID).
		Scan(&user.UserID, &user.Name, &user.Photo)
	return user, err
}
