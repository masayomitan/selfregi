package middleware

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CorsMiddleware(w http.ResponseWriter) error {
	err := godotenv.Load((".env"))
	if err != nil {
			panic ("failed reading env file from middleware")
	}
	protocol := os.Getenv("PROTOCOL")
	host := os.Getenv("FRONT_HOST")

	// こんな感じでローカルかどうか分岐
	// if tools.IsProductionEnv() {
	// 	protocol = "https://"
		// host = os.Getenv("FRONT_HOST")
	// }
	w.Header().Set("Access-Control-Allow-Origin", protocol + host)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie, application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, DELETE, OPTIONS")
	return nil
}