package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/libstring"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const PORT = ":8080"

// implement rate limiting
func RateLimiterMiddleware(lmt *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := libstring.RemoteIP([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}, 0, r)
			httpError := tollbooth.LimitByKeys(lmt, []string{ip})
			if httpError != nil {
				w.WriteHeader(httpError.StatusCode)
				w.Write([]byte(httpError.Message))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// auth middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// implement router
	r := mux.NewRouter()
	// set a rate-limit of request per minutes
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetTokenBucketExpirationTTL(time.Minute)
	// Apply rate limiting middleware
	r.Use(RateLimiterMiddleware(lmt))

	// route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the api gateway"))
	})
	r.Handle("/secure", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is secure endpoint"))
	})))
	// starting server
	fmt.Printf("Server is starting at http::/127.0.0.1:%s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println("Error Occur while server", err)
	}
}

// type Message struct {
// 	Status string `json:"status"`
// 	Body   string `json:"body"`
// }

// func endpointHandler(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	writer.WriteHeader(http.StatusOK)
// 	message := Message{
// 		Status: "Successful",
// 		Body:   "Hi! You've reached the API. How may I help you?",
// 	}
// 	err := json.NewEncoder(writer).Encode(&message)
// 	if err != nil {
// 		return
// 	}
// }

// func main() {
// 	message := Message{
// 		Status: "Request Failed",
// 		Body:   "The API is at capacity, try again later.",
// 	}
// 	jsonMessage, _ := json.Marshal(message)

// 	tlbthLimiter := tollbooth.NewLimiter(3, nil)
// 	tlbthLimiter.SetMessageContentType("application/json")
// 	tlbthLimiter.SetMessage(string(jsonMessage))

// 	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Println("There was an error listening on port :8080", err)
// 	}

// }
