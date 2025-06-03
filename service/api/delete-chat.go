package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {

	// Recupera l'id della chat dai parametri dell'URL
	chatIDStr := ps.ByName("chatId")

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	// Chiama la funzione del database per eliminare la chat
	err = rt.db.DeleteChat(chatID)
	if err != nil {
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(`{"message": "Chat deleted successfully"}`)); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
