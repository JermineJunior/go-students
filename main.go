package main

import (
	"log/slog"
	"net/http"

	"github.com/JermaineJunior/goStudents/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	// handle static assets
	router.Handle("/public/*",public())
	//web routes
	router.Get("/", handlers.Make(handlers.HandleHome))

	srv := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	slog.Info("Server running  on port 8000")
	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
