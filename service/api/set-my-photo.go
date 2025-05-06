package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// isBase64Encoded verifica se una stringa è codificata in Base64
func isBase64Encoded(s string) bool {

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx *reqcontext.RequestContext) {

	// Decodifica il corpo della richiesta JSON
	var reqBody struct {
		Photo string `json:"photo"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Invalid request body")
		return
	}

	// Verifica che il contenuto della foto non sia vuoto e non superi 1 MB
	if len(reqBody.Photo) < 4 || len(reqBody.Photo) > 1374389 {
		http.Error(w, `{"error": "Invalid photo content"}`, http.StatusBadRequest)
		return
	}
	// Verifica se il contenuto è in formato Base64
	if !isBase64Encoded(reqBody.Photo) {
		http.Error(w, `{"error": "Photo is not in valid Base64 format"}`, http.StatusBadRequest)
		return
	}

	// Aggiorna il campo "photo" nella tabella "users_table"
	err = rt.db.UpdateUserPhoto(reqBody.Photo, ctx.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			ctx.Logger.Error("User not found")
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Internal server error")
		return
	}

	// Risposta di successo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "Photo updated successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
