package api

import (
	"Wasa-Project-2024/service/api/reqcontext"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) bearerAuth(fn httpRouterHandler) httpRouterHandler {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

		// Estraggo l'header "Authentication" dalla richiesta
		authHeader := r.Header.Get("Authorization")

		// Controllo che l'header sia presente e che sia del tipo bearer
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			ctx.Logger.WithField("authentication_header:", authHeader).Warn("Missing or invalid bearer autentication header")
			http.Error(w, "Unauthorized - missing token", http.StatusUnauthorized)
			return
		}

		// Estraggo il token dall'header (rimuovo "bearer" dall'header)
		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := strconv.Atoi(token)

		if err != nil {
			ctx.Logger.WithError(err).Warnf("Invalid token format <%s>", token)
			http.Error(w, "Unauthorized - invalid token format", http.StatusUnauthorized)
			return
		}

		// Verifico se il token corrisponde ad un userId di un utente già registrato
		exist, err := rt.db.UserIdExist(userId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.Logger.WithError(err).Warnf("Token: <%s> does not exist", token)
				http.Error(w, "Unauthorized - token not valid or deprecated", http.StatusInternalServerError)
				return
			}
			ctx.Logger.WithError(err).Warnf("Error checking token existance")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !exist {
			// Se l'userId non esiste e quindi l'utente non esiste e il token presente nell'header non è valido,
			// restituisco una richiesta Unathorized
			ctx.Logger.WithField("token", token).Warnf("User id token not present in the database")
			http.Error(w, "Unnathorized", http.StatusUnauthorized)
			return
		}

		// If the token is valid call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps, ctx)
	}
}
