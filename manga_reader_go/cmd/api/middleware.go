package main

import (
	"encoding/json"
	"net/http"
)

func (a *application) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("Request", "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (a *application) authMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			req := struct {
				Token string `json:"token"`
			}{}

			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Token not found", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
