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

	authenticatedUserID := ctx.UserID
	found := false
	for _, id := range request.Users {
		if id == authenticatedUserID {
			found = true
			break
		}
	}
	if !found {
		request.Users = append(request.Users, authenticatedUserID)
	}

	if request.Name == "" || len(request.Users) == 0 {
		http.Error(w, `{"error": "Missing chat name or users"}`, http.StatusBadRequest)
		ctx.Logger.Error("Chat name or users list missing in request")
		return
	}

	// Verifica se è una chat privata già esistente
	if !request.IsGroup {
		if len(request.Users) != 2 {
			http.Error(w, `{"error": "A private chat must have exactly 2 members"}`, http.StatusBadRequest)
			ctx.Logger.Error("Private chat with invalid number of members")
			return
		}

		existingChatID, err := rt.db.FindPrivateChat(request.Users[0], request.Users[1])
		if err != nil {
			http.Error(w, `{"error": "Database error during chat lookup"}`, http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Database error in FindPrivateChat")
			return
		}
		if existingChatID != 0 {
			// Chat già esistente, restituiscila
			response := CreateChatResponse{ChatID: existingChatID}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				ctx.Logger.WithError(err).Error("Failed to encode existing chat response")
				http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
				return
			}
			ctx.Logger.Infof("Chat privata già esistente restituita (ID: %d)", existingChatID)
			return
		}
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
