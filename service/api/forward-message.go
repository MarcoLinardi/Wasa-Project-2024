package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	originalChatIDStr := ps.ByName("chatId")
	originalMsgIDStr := ps.ByName("messageId")

	originalChatID, err := strconv.Atoi(originalChatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	originalMsgID, err := strconv.Atoi(originalMsgIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid message ID"}`, http.StatusBadRequest)
		return
	}

	var requestBody struct {
		DestinationChatID int `json:"destinationChatId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	err = rt.db.ForwardMessage(originalChatID, originalMsgID, requestBody.DestinationChatID, ctx.UserID)
	if err != nil {
		http.Error(w, `{"error": "Failed to forward message"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to forward message")
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(`{"message": "Message forwarded successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
