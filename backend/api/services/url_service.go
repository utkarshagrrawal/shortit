package services

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"os"
	"shortit/db"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateShortURLService(url, proxyIP, agent, ip, username string) (string, string) {
	var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortURL string
	for i := 0; i < 12; i++ {
		shortURL += string(characters[rand.Intn(len(characters))])
	}
	shortURL = os.Getenv("BASE_URL") + "/" + shortURL
	result := db.URLCollection.FindOne(context.TODO(), bson.D{{Key: "original_url", Value: url}})
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return "Error checking if short URL already exists", ""
	}
	if result.Err() == nil {
		var existingURL map[string]interface{}
		err := result.Decode(&existingURL)
		if err != nil {
			return "Error decoding existing URL", ""
		}
		return "", existingURL["short_url"].(string)
	}
	reqPayload := `{"client": {"clientId": "shortit", "clientVersion": "1.0.0"}, "threatInfo": {"threatTypes": ["MALWARE", "SOCIAL_ENGINEERING", "UNWANTED_SOFTWARE", "POTENTIALLY_HARMFUL_APPLICATION"], "platformTypes": ["ANY_PLATFORM"], "threatEntryTypes": ["URL"], "threatEntries": [{"url": "` + url + `"}]}}`
	request, err := http.NewRequest("POST", "https://safebrowsing.googleapis.com/v4/threatMatches:find?key="+os.Getenv("GOOGLE_SAFE_BROWSING_API_KEY"), strings.NewReader(reqPayload))
	if err != nil {
		return "Error creating threat check request", ""
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Referer", os.Getenv("GOOGLE_SAFE_BROWSING_API_REFERER"))
	threatMatchResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		return "Error checking for threats in URL", ""
	}
	if threatMatchResponse.StatusCode != 200 {
		return "Error checking for threats in URL", ""
	}
	defer threatMatchResponse.Body.Close()
	threatMatchResponseData, err := io.ReadAll(threatMatchResponse.Body)
	if err != nil {
		return "Error reading response from threat check", ""
	}
	if string(threatMatchResponseData) != "{}\n" {
		return "URL is not safe", ""
	}
	retryCount := 0
	_, err = db.URLCollection.InsertOne(context.TODO(), bson.D{{Key: "original_url", Value: url}, {Key: "short_url", Value: shortURL}, {Key: "clicks", Value: 0}, {Key: "created_at", Value: time.Now()}, {Key: "IP", Value: ip}, {Key: "proxy_IP", Value: proxyIP}, {Key: "user_agent", Value: agent}, {Key: "username", Value: username}})
	for err != nil && retryCount < 3 {
		shortURL = ""
		for i := 0; i < 12; i++ {
			shortURL += string(characters[rand.Intn(len(characters))])
		}
		shortURL = os.Getenv("BASE_URL") + "/" + shortURL
		_, err = db.URLCollection.InsertOne(context.TODO(), bson.D{{Key: "original_url", Value: url}, {Key: "short_url", Value: shortURL}, {Key: "clicks", Value: 0}, {Key: "created_at", Value: time.Now()}, {Key: "IP", Value: ip}, {Key: "proxy_IP", Value: proxyIP}, {Key: "user_agent", Value: agent}, {Key: "username", Value: username}})
		retryCount++
	}
	if err != nil {
		return "Error creating short URL", ""
	}
	return "", shortURL
}

func RedirectService(shortURL string) (string, string) {
	result := db.URLCollection.FindOne(context.TODO(), bson.D{{Key: "short_url", Value: os.Getenv("BASE_URL") + "/" + shortURL}})
	if result.Err() != nil {
		return "Error finding short URL", ""
	}
	var url map[string]interface{}
	err := result.Decode(&url)
	if err != nil {
		return "Error decoding short URL", ""
	}
	_, err = db.URLCollection.UpdateOne(context.TODO(), bson.D{{Key: "short_url", Value: os.Getenv("BASE_URL") + "/" + shortURL}}, bson.D{{Key: "$inc", Value: bson.D{{Key: "clicks", Value: 1}}}})
	if err != nil {
		return "Error updating short URL analytics", ""
	}
	return "", url["original_url"].(string)
}
