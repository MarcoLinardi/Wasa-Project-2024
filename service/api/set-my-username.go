package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"Wasa-Project-2024/service/utilitytool"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	var reqBody struct {
		NewName string `json:"newName"`
	}

	// Decodifica il corpo della richiesta JSON
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Valida il nuovo username
	if err := utilitytool.UserNameIsValid(reqBody.NewName); err != nil {
		http.Error(w, `{"error": "Invalid username format"}`, http.StatusBadRequest)
		return
	}

	// Aggiorna il database
	err := rt.db.UpdateUsername(reqBody.NewName, ctx.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			ctx.Logger.WithError(err).Error("User not found")
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Internal server error")
		return
	}

	// Risposta di successo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "Username changed successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
