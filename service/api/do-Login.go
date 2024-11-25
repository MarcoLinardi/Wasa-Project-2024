package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// User è la struttura che rappresenta l'input richiesto
type User struct {
	Name string `json:"name"` // Il nome dell'utente passato nella richiesta
}

// LoginResponse è la struttura per la risposta
type LoginResponse struct {
	Identifier string `json:"identifier"` // Identificatore restituito al client
}

// DoLogin è l'handler per l'endpoint /login
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Decodifica il body JSON della richiesta
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		ctx.Logger.Error("Failed to decode request body:", err)
		return
	}

	// Verifica se il nome è presente
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		ctx.Logger.Error("Name is missing in the request")
		return
	}

	// Simulazione: crea un identificatore per l'utente
	// Qui puoi aggiungere logica per controllare se l'utente esiste già
	identifier := "abcdef012345" // Puoi generare un identificatore univoco

	// Prepara la risposta
	response := LoginResponse{
		Identifier: identifier,
	}

	// Invia la risposta al client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		ctx.Logger.Error("Failed to encode response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Logga il successo
	ctx.Logger.Info("User login successful for name:", user.Name)
}
