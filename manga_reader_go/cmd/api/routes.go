package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/health", a.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/user/register", a.registerUser)

	router.Handler(http.MethodGet, "/manga", a.authMiddleware(a.searchManga))
	router.Handler(http.MethodGet, "/manga/:name", a.authMiddleware(a.getManga))
	router.Handler(http.MethodGet, "/read/:mangaName/:chapterName", a.authMiddleware(a.getChapter))

	return a.logMiddleware(router)
}
