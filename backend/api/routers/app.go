package routers

import (
	"net/http"
	"shortit/api/handlers"
	"shortit/api/middlewares"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/shorten", middlewares.RateLimit(http.HandlerFunc(handlers.CreateShortURL))).Methods("POST", "OPTIONS")
	router.HandleFunc("/{shortId}", handlers.RedirectToOriginalURL).Methods("GET", "OPTIONS")
	router.HandleFunc("/{shortId}", handlers.DeleteShortUrl).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user/urls", handlers.GetUserAllUrls).Methods("GET", "OPTIONS")

	return router
}
