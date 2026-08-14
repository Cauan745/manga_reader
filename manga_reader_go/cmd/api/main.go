package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	repo "manga-reader/internal/adapters/postgresql/sqlc"
	"manga-reader/internal/services"

	"github.com/jackc/pgx/v5"
)

type dbConfig struct {
	port int
	dsn  string
}

type config struct {
	port int
	env  string
	db   dbConfig
}

type application struct {
	logger       *slog.Logger
	config       config
	mangaService services.MangaService
	// db           *pgx.Conn

	db *repo.Queries
}

func main() {
	var cfg config
	var dbCfg dbConfig

	flag.IntVar(&cfg.port, "addr", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.IntVar(&dbCfg.port, "dbport", 5432, "Database port")
	flag.StringVar(
		&dbCfg.dsn,
		"dbdsn",
		"host=localhost user=postgres password=postgres dbname=manga_reader_db sslmode=disable",
		"Database dsn string")
	flag.Parse()

	cfg.db = dbCfg

	ctx := context.Background()

	mangaService := services.NewMangaService()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	//

	app := application{
		config: cfg,
		logger: logger,
		// db:           conn,
		db:           repo.New(conn),
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

	err = srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
