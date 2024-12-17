package handlers

import (
	"encoding/json"
	"net/http"
	"shortit/api/services"

	"github.com/gorilla/mux"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	response := services.CreateShortURLService(request["url"].(string), r.RemoteAddr, r.Header.Get("X-Forwarded-For"), r.Header.Get("user-agent"))
	if response == "" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]
	originalURL := services.RedirectService(shortURL)
	if originalURL == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(originalURL)
}
