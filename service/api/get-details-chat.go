package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) detailsChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	chatIDStr := ps.ByName("chatId")
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		http.Error(w, "ID chat non valido", http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	chat, err := rt.db.GetChatDetails(chatID, userID)
	if err != nil {
		http.Error(w, "Errore: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chat); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
