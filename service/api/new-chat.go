package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CreateChatRequest struct {
	Name    string `json:"name"`
	Users   []int  `json:"users"`
	IsGroup bool   `json:"isGroup"`
}

type CreateChatResponse struct {
	ChatID int `json:"chatId"`
}

func (rt *_router) createChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	var request CreateChatRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Failed to decode createChat request")
		return
	}

	if request.Name == "" || len(request.Users) == 0 {
		http.Error(w, `{"error": "Missing chat name or users"}`, http.StatusBadRequest)
		ctx.Logger.Error("Chat name or users list missing in request")
		return
	}

	if !request.IsGroup && len(request.Users) != 2 {
		http.Error(w, `{"error": "A private chat must have exactly 2 members"}`, http.StatusBadRequest)
		ctx.Logger.Error("Private chat with invalid number of members")
		return
	}

	if request.IsGroup && len(request.Users) < 2 {
		http.Error(w, `{"error": "A group must have at least 2 members"}`, http.StatusBadRequest)
		ctx.Logger.Error("Group with not enough members")
		return
	}

	// Crea la chat
	chatID, err := rt.db.CreateChat(request.Name, request.Users, request.IsGroup)
	if err != nil {
		http.Error(w, `{"error": "Failed to create chat"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Database error in CreateChat")
		return
	}

	// Risposta
	response := CreateChatResponse{ChatID: chatID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		ctx.Logger.WithError(err).Error("Failed to encode CreateChat response")
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	ctx.Logger.Infof("Chat creata con successo: %s (ID: %d)", request.Name, chatID)
}
