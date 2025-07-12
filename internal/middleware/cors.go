package middleware

import (
	"net/http"
)

// CORS is a middleware that sets CORS headers for allowed origins.
func CORS(next http.Handler) http.Handler {
	allowedOrigins := map[string]bool{
		"http://python-api:3000": true,
		"http://frontend:3001":   true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// If preflight request (OPTIONS), respond immediately
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
