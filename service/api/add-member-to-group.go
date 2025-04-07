package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addMemberToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	chatIDStr := ps.ByName("chatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	var requestBody struct {
		UserID int `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if requestBody.UserID == 0 {
		http.Error(w, `{"error": "User ID is required"}`, http.StatusBadRequest)
		return
	}

	// Aggiunge l'utente al gruppo
	err = rt.db.AddMemberToGroup(chatID, requestBody.UserID)
	if err != nil {
		if err.Error() == "chat is not a group" {
			http.Error(w, `{"error": "Cannot add members to a non-group chat"}`, http.StatusBadRequest)
			return
		}
		http.Error(w, `{"error": "Failed to add member"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to add member to group")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User added to group successfully"}`))
}
