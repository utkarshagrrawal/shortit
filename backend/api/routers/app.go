package routers

import (
	"shortit/api/handlers"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.CreateShortURL).Methods("POST", "OPTIONS")
	router.HandleFunc("/{shortId}", handlers.RedirectToOriginalURL).Methods("GET", "OPTIONS")
	router.HandleFunc("/{shortId}", handlers.DeleteShortUrl).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user/urls", handlers.GetUserAllUrls).Methods("GET", "OPTIONS")

	return router
}
