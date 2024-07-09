package main

import (
	"github.com/gorilla/mux"
	"github.com/klaital/todolist/cmd/server/api"
	"net/http"
)

func main() {
	server := api.NewServer()

	router := mux.NewRouter()
	spa := spaHandler{
		staticPath: "web/todo/dist",
		indexPath:  "index.html",
	}
	router.PathPrefix("/web").Handler(spa)

	srv := api.NewServer()
	h := api.HandlerFromMux(&srv, router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080", "http://bedroom-tv"},
		AllowedMethods: []string{"GET", "PUT"},
	})
	s := &http.Server{
		Handler: c.Handler(h),
		Addr:    "0.0.0.0:8080",
	}
}
