package middleware

import (
		"net/http"
		"os"
		"github.com/joho/godotenv"
)

func CorsMiddleware(w http.ResponseWriter, r *http.Request) error {
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
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return nil
		}

		return nil
}