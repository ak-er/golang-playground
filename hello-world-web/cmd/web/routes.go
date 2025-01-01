package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ak-er/golang-journey/pkg/config"
	"github.com/ak-er/golang-journey/pkg/handlers"
	"github.com/bmizerany/pat"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// using pat routing
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	fmt.Println(app)
	return mux
}

// using chi routing
func routesChi(app *config.AppConfig) http.Handler {
	log.Println(app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.With(WriteToConsole).Get("/about", handlers.Repo.About)
	return mux
}

// simple http
func Routes() {
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
}
