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

	// Decodifica il corpo della richiesta JSON
	var msg database.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, `{"error": "Invalid JSON body"}`, http.StatusBadRequest)
		return
	}

	// Controllo che il messaggio non sia vuoto
	if msg.Content == "" {
		http.Error(w, `{"error": "Message content cannot be empty"}`, http.StatusBadRequest)
		return
	}

	// Imposta l'ID della chat
	msg.ChatID = chatID

	// Salva il messaggio nel database
	messageID, err := rt.db.SaveMessage(msg)
	if err != nil {
		http.Error(w, `{"error": "Failed to save message"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to save message")
		return
	}

	// Risponde con 201 Created e l'ID del messaggio creato
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"message_id": messageID})
}
