package database

func (db *appdbimpl) UpdateUsername(newName string, userId int) error {

	_, err := db.c.Exec("UPDATE users_table SET name = ? WHERE userId = ?", newName, userId)
	if err != nil {
		return err
	}
	return nil
}
