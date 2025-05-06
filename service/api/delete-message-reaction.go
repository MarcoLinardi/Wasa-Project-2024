package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteReaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
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

	// Verifica che il messaggio appartenga alla chat
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

	// Elimina la reazione
	err = rt.db.DeleteReaction(messageID, ctx.UserID)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete reaction"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to delete reaction")
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "Reaction deleted successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
