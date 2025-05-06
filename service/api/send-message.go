package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"Wasa-Project-2024/service/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	// Recupera l'ID della chat dall'URL
	chatIDStr := ps.ByName("chatId")
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

	msg.SenderID = ctx.UserID

	// Aggiungi automaticamente il timestamp
	msg.Timestamp = time.Now().Format("2006-01-02 15:04:05")

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
	if err := json.NewEncoder(w).Encode(map[string]int{"message_id": messageID}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}

}
