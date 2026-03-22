package main

import "net/http"

func (a *application) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("Request", "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
