package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", a.healthcheckHandler)

	return router
}
