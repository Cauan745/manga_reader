package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/health", a.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/manga", a.searchManga)
	router.HandlerFunc(http.MethodGet, "/manga/:name", a.getManga)
	router.HandlerFunc(http.MethodGet, "/read/:mangaName/:chapterName", a.getChapter)

	router.HandlerFunc(http.MethodPost, "/user/register", a.registerUser)

	return a.logMiddleware(router)
}
