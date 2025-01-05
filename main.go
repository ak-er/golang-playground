package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Helper function to hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// setup database
var db *sql.DB

func init() {
	var err error
	// get db location
	db, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal("connect with db failed", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE, password TEXT, role TEXT)`)
	if err != nil {
		log.Fatal("error get while creating table", err)
	}
	hastadminpass, _ := hashPassword("adminpassword")
	hashuserpass, _ := hashPassword("userpassword")
	_, err = db.Exec(`
	   INSERT OR IGNORE INTO users (username, password, role)
	   VALUES
	   ('adminname', ?, 'admin'),
	   ('username', ?, 'user') 
	`, hastadminpass, hashuserpass)
	if err != nil {
		log.Fatal("error get inserting the user in users table", err)
	}
}

// Secret key for signin using JWT token
var jwtSecret = []byte("my-secret-key")

// user struct for authentication
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Generate JWT Token
func generateToken(username, role string) (string, string, error) {
	tokenClaim := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Minute * 5).Unix(), // expiration time 5 minutes
	}
	// access-token short lived
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// refresh-token long lived
	refreshClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // expiration time 7 days
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

// Middleware to verify JWT token
func authenticate(allowRoles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			prefix := "Bearer "
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is Required", http.StatusUnauthorized)
				return
			}
			if len(authHeader) <= len(prefix) {
				http.Error(w, "invalid authorization header: token missing", http.StatusBadRequest)
				return
			}
			tokenString := authHeader[len(prefix):]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signin method")
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid token claims", http.StatusUnauthorized)
				return
			}
			userRole := claims["role"].(string)
			for _, role := range allowRoles {
				if userRole == role {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
	}
}

// Login Handler
func login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// get data from db
	var hashPassword, role string
	err = db.QueryRow(`SELECT password, role FROM users WHERE username=?`, user.Username).Scan(&hashPassword, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := generateToken(user.Username, role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}

// implement refresh-token endpoint
func refreshTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RefreshToken string `json:"refresh-token"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.RefreshToken == "" {
		http.Error(w, "invalid refresh token", http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(body.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("token.Method", token.Method)
			return nil, fmt.Errorf("unexpected string method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "invalid or expired refresh token", http.StatusUnauthorized)
		log.Println("token refreshed failed: invalid refresh token", err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil {
		http.Error(w, "invalid token claims", http.StatusUnauthorized)
		return
	}
	username := claims["username"].(string)
	var role string
	err = db.QueryRow(`SELECT role FROM users WHERE username=?`, username).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusBadRequest)
		return
	}

	accessToken, _, err := generateToken(username, role)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		log.Println("failed to generate token--err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access-token": accessToken,
	})
}

// Protect endpoint
func userProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "welcome to user protected endpoint"})
}
func adminProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "welcome to admin protected endpoint"})
}

const PORT = ":8080"

func main() {
	// routes
	http.HandleFunc("/login", login)
	http.HandleFunc("/user-protected-endpoint", authenticate("admin", "user")(userProtectedEndpoint))
	http.HandleFunc("/admin-protected-endpoint", authenticate("admin")(adminProtectedEndpoint))
	http.HandleFunc("/refresh-token", refreshTokenEndpoint)
	// starting server
	fmt.Printf("Server is starting at http::/127.0.0.1:%s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Error Occur while server", err)
	}
}
