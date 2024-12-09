package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"Wasa-Project-2024/service/utilitytool"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Decodifica il corpo della richiesta JSON
	var reqBody struct {
		NewName string `json:"newName"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Valida il nuovo username
	if err := utilitytool.UserNameIsValid(reqBody.NewName); err != nil {
		http.Error(w, `{"error": "Invalid username format"}`, http.StatusBadRequest)
		return
	}

	// Recupera l'userId dal contesto
	userId := ctx.UserID

	// Aggiorna il database
	query := "UPDATE users_table SET name = ? WHERE id = ?"
	result, err := rt.db.Exec(query, reqBody.NewName, userId)
	if err != nil {
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	// Verifica se è stato modificato qualcosa
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, `{"error": "User not found or no changes made"}`, http.StatusNotFound)
		return
	}

	// Risposta di successo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "Username changed successfully"}`))
}
