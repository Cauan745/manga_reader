package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"manga-reader/internal/services"
)

type config struct {
	port int
	env  string
}

type application struct {
	logger       *slog.Logger
	config       config
	mangaService services.MangaService
}

func main() {
	mangaService := services.NewMangaService()

	cfg := config{
		port: 8000,
		env:  "development",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{
		config:       cfg,
		logger:       logger,
		mangaService: mangaService,
	}

	// Use the httprouter instance returned by app.routes() as the server handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
