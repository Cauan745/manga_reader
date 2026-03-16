package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/health", a.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/manga/search/:query", a.searchManga)
	return router
}
