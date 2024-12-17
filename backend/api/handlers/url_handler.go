package handlers

import (
	"encoding/json"
	"net/http"
	"shortit/api/services"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	response := services.CreateShortURLService(request["url"].(string))
	if response == "" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
