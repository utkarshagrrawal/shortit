package middlewares

import (
	"net/http"
	"net/url"
)

func ApplyCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allowedOrigins = make(map[string]bool)
		allowedOrigins["localhost:5173"] = true
		allowedOrigins["localhost:5174"] = true
		allowedOrigins["shortitin.vercel.app"] = true

		origin := r.Header.Get("origin")

		originUrl, err := url.Parse(origin)
		if err != nil {
			http.Error(w, "Invalid origin", http.StatusBadRequest)
			return
		}

		if ok := allowedOrigins[originUrl.Host]; ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
	})
}
