package main

import (
	"net/http"
	"os"
	"shortit/api/middlewares"
	"shortit/api/routers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Use(middlewares.ApplyCors)
	router.Use(middlewares.CreateLogs)

	router.PathPrefix("/api/v1").Handler(http.StripPrefix("/api/v1", routers.AppRouter()))

	port := ":" + os.Getenv("PORT")

	http.ListenAndServe(port, router)
}
