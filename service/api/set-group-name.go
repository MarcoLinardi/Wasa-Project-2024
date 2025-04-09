package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	chatIDStr := ps.ByName("chatId")

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	var reqBody struct {
		NewName string `json:"newName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if reqBody.NewName == "" {
		http.Error(w, `{"error": "New name cannot be empty"}`, http.StatusBadRequest)
		return
	}

	err = rt.db.SetGroupName(chatID, reqBody.NewName)
	if err != nil {
		http.Error(w, `{"error": "Failed to rename group chat"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to rename group chat")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Group chat renamed successfully"}`))
}
