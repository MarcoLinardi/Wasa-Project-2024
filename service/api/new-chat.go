package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CreateChatRequest struct {
	Name  string `json:"name"`
	Users []int  `json:"users"`
}

type CreateChatResponse struct {
	ChatID int `json:"chatId"`
}

func (rt *_router) createChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	var request CreateChatRequest

	// Decodifica il corpo della richiesta
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Failed to decode createChat request")
		return
	}

	// Validazione base
	if request.Name == "" || len(request.Users) == 0 {
		http.Error(w, "Missing chat name or users", http.StatusBadRequest)
		ctx.Logger.Error("Chat name or users list missing in request")
		return
	}

	// Chiamata al database per creare la chat
	chatID, err := rt.db.CreateChat(request.Name, request.Users)
	if err != nil {
		http.Error(w, "Failed to create chat", http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Database error in CreateChat")
		return
	}

	// Costruzione e invio della risposta
	response := CreateChatResponse{ChatID: chatID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		ctx.Logger.WithError(err).Error("Failed to encode CreateChat response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Log
	ctx.Logger.Infof("Chat creata con successo: %s (ID: %d)", request.Name, chatID)
}
