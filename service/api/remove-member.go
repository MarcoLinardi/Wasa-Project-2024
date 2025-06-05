package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeMemberToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	chatIDStr := ps.ByName("chatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	var requestBody struct {
		UserIDs []int `json:"userIds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if len(requestBody.UserIDs) == 0 {
		http.Error(w, `{"error": "userIds array is required"}`, http.StatusBadRequest)
		return
	}

	for _, userID := range requestBody.UserIDs {
		err = rt.db.RemoveMemberFromGroup(chatID, userID)
		if err != nil {
			http.Error(w, `{"error": "Failed to remove members from group"}`, http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Failed to remove members from group")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "Members removed successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
