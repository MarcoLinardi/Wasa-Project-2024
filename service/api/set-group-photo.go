package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	chatIDStr := ps.ByName("chatId")

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid chat ID"}`, http.StatusBadRequest)
		return
	}

	var reqBody struct {
		NewPhoto string `json:"newPhoto"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if reqBody.NewPhoto == "" {
		http.Error(w, `{"error": "New photo cannot be empty"}`, http.StatusBadRequest)
		return
	}

	err = rt.db.SetGroupPhoto(chatID, reqBody.NewPhoto)
	if err != nil {
		http.Error(w, `{"error": "Failed to update group photo"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Failed to update group photo")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Group photo updated successfully"}`))
}
