package main

import (
	"fmt"
	"net/http"

	"github.com/sachintha-lk/go-rss-agg/internal/auth"
	"github.com/sachintha-lk/go-rss-agg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %s", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %s", err))
			return
		}

		handler(w, r, user)
	}
}
