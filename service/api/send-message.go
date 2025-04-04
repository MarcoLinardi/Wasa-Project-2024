package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"Wasa-Project-2024/service/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	// Recupera l'ID della chat dall'URL
	chatIDStr := ps.ByName("ChatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	// Decodifica il corpo della richiesta
	var msg database.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Prendi l'ID dell'utente autenticato (ipotizziamo che sia in `ctx.UserID`)
	msg.SenderID = ctx.UserID

	// Salva il messaggio nel database
	messageID, err := rt.db.SaveMessage(chatID, msg)
	if err != nil {
		http.Error(w, `{"error": "Failed to save message"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Errore durante il salvataggio del messaggio")
		return
	}

	// Risposta di successo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"message_id": messageID})
}
