package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) reactToMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	chatIDStr := ps.ByName("chatId")
	msgIDStr := ps.ByName("messageId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}
	messageID, err := strconv.Atoi(msgIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid message ID"}`, http.StatusBadRequest)
		return
	}
	exists, err := rt.db.CheckMessageInChat(chatID, messageID)
	if err != nil {
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error checking message in chat")
		return
	}
	if !exists {
		http.Error(w, `{"error": "Message does not belong to chat"}`, http.StatusBadRequest)
		return
	}
	var requestBody struct {
		Reaction string `json:"reaction"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	if requestBody.Reaction == "" {
		http.Error(w, `{"error": "Reaction cannot be empty"}`, http.StatusBadRequest)
		return
	}
	// Aggiunge la reazione
	err = rt.db.AddReaction(messageID, ctx.UserID, requestBody.Reaction)
	if err != nil {
		http.Error(w, `{"error": "Failed to add reaction"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to add reaction")
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(`{"message": "Reaction added successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
