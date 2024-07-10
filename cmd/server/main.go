package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/klaital/todolist/cmd/server/api"
	"github.com/klaital/todolist/internal/config"
	"github.com/klaital/todolist/internal/storage/queries"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DbDsn(""))
	if err != nil {
		slog.Error("Failed to connect to db", "error", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	spa := spaHandler{
		staticPath: "web/todo/dist",
		indexPath:  "index.html",
	}
	router.PathPrefix("/web").Handler(spa)

	srv := api.NewServer(queries.New(db))
	h := api.HandlerFromMux(&srv, router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080", "http://bedroom-tv"},
		AllowedMethods: []string{"GET", "PUT"},
	})
	s := &http.Server{
		Handler: c.Handler(h),
		Addr:    "0.0.0.0:8080",
	}

	s.ListenAndServe()
}
