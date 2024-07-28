package main

import (
	"fmt"
	"net/http"

	"github.com/harryfung22/ScrapeHive/internal/auth"
	"github.com/harryfung22/ScrapeHive/internal/databse"
)

type authHandler func(http.ResponseWriter, *http.Request, databse.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			resErr(w, 403, fmt.Sprintf("An error occured with authentication: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			resErr(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}
