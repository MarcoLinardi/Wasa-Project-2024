package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeMemberToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	chatIDStr := ps.ByName("chatId")
	userIDStr := ps.ByName("userId")

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	err = rt.db.RemoveMemberFromGroup(chatID, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to remove member from group"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to remove member from group")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Member removed successfully"}`))
}
