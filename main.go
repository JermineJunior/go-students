package main

import (
	"log/slog"
	"net/http"

	"github.com/JermaineJunior/goStudents/internal/db"
	"github.com/JermaineJunior/goStudents/internal/handlers"
	"github.com/JermaineJunior/goStudents/internal/records"
	"github.com/JermaineJunior/goStudents/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		slog.Error(err.Error())
	}
	defer db.Close()
	// init services,records {Bootstrap}
	studentRecord := records.NewStudentRecord(db)
	studentService := services.NewStudentService(studentRecord)
	studentHanler := handlers.NewStudentHandler(studentService)

	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// handle static assets
	router.Handle("/public/*", public())
	//web routes
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.NotFound(handlers.Make(handlers.HandleNotFound))
	// student Routes
	router.Get("/students", handlers.Make(studentHanler.Index))
	srv := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	slog.Info("Server running  on port 8000")
	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
