package handlers

import (
	"encoding/json"
	"net/http"
	"os"
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
	ip, ok := request["ip"].(string)
	if !ok {
		ip = r.RemoteAddr
	}
	usernameCookie, err := r.Cookie("username")
	if err != nil && err != http.ErrNoCookie {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	var username string
	if err == http.ErrNoCookie {
		username = services.GenerateUsernameService()
		if username == "Error creating user" {
			http.Error(w, username, http.StatusInternalServerError)
			return
		}
	} else {
		username = usernameCookie.Value
	}
	response, url := services.CreateShortURLService(request["url"].(string), r.Header.Get("X-Forwarded-For"), r.Header.Get("user-agent"), ip, username)
	if response != "" {
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "username",
		Value:    username,
		HttpOnly: true,
		MaxAge:   86400 * 400,
		Path:     "/api/v1",
	}
	if os.Getenv("ENV") == "production" {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(url)
}

func RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortId := params["shortId"]
	err, originalURL := services.RedirectService(shortId)
	if err != "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(originalURL)
}

func GetUserAllUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	username := usernameCookie.Value
	allURLs := services.GetUserAllURLsService(username)
	if allURLs == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allURLs)
}

func DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	username := cookie.Value
	params := mux.Vars(r)
	shortId := params["shortId"]
	response := services.DeleteUrlService(username, shortId)
	if response != "" {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
