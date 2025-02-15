package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ak-er/golang-playground/internal/config"
	"github.com/ak-er/golang-playground/internal/http/handlers/student"
	"github.com/ak-er/golang-playground/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad() // load config
	// setup db
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Failed to connect database", err.Error())
	}
	slog.Info("database connected", slog.String("env", cfg.Env))
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", student.List(storage))
	router.HandleFunc("POST /api/students/create", student.Create(storage))
	router.HandleFunc("GET /api/students/{id}", student.Get(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.Delete(storage))
	// setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}
	slog.Info("Server started", slog.String("address", cfg.HttpServer.Addr))

	// graceful shutdown
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server", err.Error())
		}
	}()
	<-done
	slog.Info("sutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
