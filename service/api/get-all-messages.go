package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getAllMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	// Estraggo chatId dalla URL
	chatIDStr := ps.ByName("chatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	// Chiamo il database
	messages, err := rt.db.GetAllMessages(chatID)
	if err != nil {
		http.Error(w, `{"error": "Error retrieving messages"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error retrieving messages")
		return
	}

	// Risposta JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
