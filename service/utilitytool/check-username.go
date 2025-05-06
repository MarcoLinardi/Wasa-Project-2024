package utilitytool

import (
	"fmt"
)

func UserNameIsValid(username string) error {

	// Controlla se l'ID è nel range valido
	if len(username) < 3 {
		return fmt.Errorf("l'username %s è troppo corto (3-20)", username)
	}
	if len(username) > 20 {
		return fmt.Errorf("l'username %s è troppo lungo (3-20)", username)
	}
	// Se l'ID è valido, restituisci nil
	return nil
}
