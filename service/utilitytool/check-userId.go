package utilitytool

import (
	"fmt"
)

func UserIdIsValid(userId int) error {

	// Controlla se l'ID è nel range valido
	if userId < 100 {
		return fmt.Errorf("userId %d è troppo corto (1-10000)", userId)
	}
	if userId > 10000 {
		return fmt.Errorf("userId %d è troppo lungo (1-10000)", userId)
	}
	// Se l'ID è valido, restituisci nil
	return nil
}
