package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func writeJson(data any, w http.ResponseWriter, l *slog.Logger) {
	resp, err := json.Marshal(data)
	if err != nil {
		l.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(resp)
}

func (a *application) searchManga(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	searchResult, err := a.mangaService.SearchManga(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(searchResult, w, a.logger)
}

func (a *application) getManga(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	mangaName := params.ByName("name")

	manga, err := a.mangaService.GetManga(mangaName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(manga, w, a.logger)
}

func (a *application) getChapter(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	mangaName := params.ByName("mangaName")
	chapterName := params.ByName("chapterName")

	chapter, err := a.mangaService.GetChapter(mangaName, chapterName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJson(chapter, w, a.logger)
}
