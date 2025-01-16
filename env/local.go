package env

import "os"

// Secret key for signin using JWT token
func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
