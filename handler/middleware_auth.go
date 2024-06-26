package handler

import (
	"fmt"
	"net/http"

	"github.com/confusedOrca/RSS-Aggregator/internal/auth"
	"github.com/confusedOrca/RSS-Aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			respondWithError(responseWriter, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			respondWithError(responseWriter, 400, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(responseWriter, request, user)
	}
}
