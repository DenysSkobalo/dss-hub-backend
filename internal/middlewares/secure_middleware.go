package middlewares

import (
	"net/http"
)

func SecurityMiddlewares(serviceName string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Strict CORS Policy
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:1313" || origin == "https://dss-hub-frontend.pages.dev" ||  origin == "https://denysskobalodev.space" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Preflight requests for browser
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Security Headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Service-ID", serviceName)

		next(w, r)
	}
}
