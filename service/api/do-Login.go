package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// DoLogin è l'handler per l'endpoint /login
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *reqcontext.RequestContext) {
	var requestBody = struct {
		Name string `json:"name"` // Il nome dell'utente passato nella richiesta
	}{}

	var LoginResponse = struct {
		Identifier int `json:"identifier"`
	}{}

	// Decodifica il body JSON della richiesta
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Failed to decode request body")
		return
	}

	// Verifica se il nome è presente nella richiesta
	if requestBody.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		ctx.Logger.Error("Name is missing in the request")
		return
	}

	// Simulazione: crea un identificatore per l'utente
	// Qui puoi aggiungere logica per controllare se l'utente esiste già
	var err error
	LoginResponse.Identifier, err = rt.db.GetUserIdByName(requestBody.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//devo creare l'utente ed inserirlo nel database
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Error while retriving user id by the user name")
			return
		}
	}

	// Invia la risposta al client
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authentication", fmt.Sprintf("bearerAuth: %d", LoginResponse.Identifier))
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(LoginResponse); err != nil {
		ctx.Logger.Error("Failed to encode response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Logga il successo
	ctx.Logger.Infof("User login successful for name: %s", requestBody.Name)
}
