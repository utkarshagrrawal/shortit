package routers

import (
	"shortit/api/handlers"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.CreateShortURL).Methods("POST", "OPTIONS")
	router.HandleFunc("/{shortURL}", handlers.RedirectToOriginalURL).Methods("GET")

	return router
}
