package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// set logger
var logger = logrus.New()

func initLogger() {
	// set log format to JSON for better readability in centralized systems.
	logger.SetFormatter(&logrus.JSONFormatter{})
	// set log-output to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Failed to open log file:: %v", err)
	}
	logger.SetOutput(file)
	// set log level (e.g. Info, Warn, Error)
	logger.SetLevel(logrus.InfoLevel)
	logger.Info("Logger Initialized")
}

// Protect endpoint
func homepage(w http.ResponseWriter, r *http.Request) {
	logger.WithFields(logrus.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"ip":     r.RemoteAddr,
	}).Info("Request Received")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "this is homepage"})
}

func errorPage(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
			"ip":     r.RemoteAddr,
			"error":  err.Error(),
		}).Error("Error Occured")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

const PORT = ":8080"

func main() {
	initLogger()
	// routes
	http.HandleFunc("/", homepage)
	http.HandleFunc("/error", errorPage(fmt.Errorf("example error")))
	// starting server
	fmt.Printf("Server is starting at http::/127.0.0.1%s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Error Occur while server", err)
	}
}
