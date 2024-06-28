package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/klaital/todolist/cmd/server/api"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	server := api.NewServer()

	r := chi.NewMux()

	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
