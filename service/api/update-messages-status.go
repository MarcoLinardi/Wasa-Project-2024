package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateMessagesStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	chatIDStr := ps.ByName("chatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	var requestBody struct {
		MessageIDs []int  `json:"messageIds"`
		NewStatus  string `json:"newStatus"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if requestBody.NewStatus != "sent" && requestBody.NewStatus != "received" && requestBody.NewStatus != "read" {
		http.Error(w, `{"error": "Invalid status value"}`, http.StatusBadRequest)
		return
	}

	err = rt.db.UpdateMessagesStatus(chatID, requestBody.MessageIDs, requestBody.NewStatus)
	if err != nil {
		http.Error(w, `{"error": "Failed to update messages"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to update messages status")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Messages status updated successfully"}`))
}
