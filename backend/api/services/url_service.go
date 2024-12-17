package services

import (
	"context"
	"math/rand"
	"os"
	"shortit/db"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateShortURLService(url string) string {
	var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortURL string
	for i := 0; i < 12; i++ {
		shortURL += string(characters[rand.Intn(len(characters))])
	}
	shortURL = os.Getenv("BASE_URL") + "/" + shortURL
	result := db.URLCollection.FindOne(context.TODO(), bson.D{{Key: "original_url", Value: url}})
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return ""
	}
	if result.Err() == nil {
		var existingURL map[string]interface{}
		err := result.Decode(&existingURL)
		if err != nil {
			return ""
		}
		return existingURL["short_url"].(string)
	}
	retryCount := 0
	_, err := db.URLCollection.InsertOne(context.TODO(), bson.D{{Key: "original_url", Value: url}, {Key: "short_url", Value: shortURL}})
	for err != nil && retryCount < 3 {
		shortURL = ""
		for i := 0; i < 12; i++ {
			shortURL += string(characters[rand.Intn(len(characters))])
		}
		shortURL = os.Getenv("BASE_URL") + "/" + shortURL
		_, err = db.URLCollection.InsertOne(context.TODO(), bson.D{{Key: "original_url", Value: url}, {Key: "short_url", Value: shortURL}})
		retryCount++
	}
	if err != nil {
		return ""
	}
	return shortURL
}
