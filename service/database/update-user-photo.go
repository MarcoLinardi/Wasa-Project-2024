package database

func (db *appdbimpl) UpdatePhoto(photo string, userId int) error {

	_, err := db.c.Exec("UPDATE users_table SET photo = ? WHERE userId = ?", photo, userId)
	if err != nil {
		return err
	}
	return nil
}
