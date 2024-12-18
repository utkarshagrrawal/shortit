package services

import (
	"context"
	"math/rand"
	"os"
	"shortit/db"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GenerateUsernameService() string {
	generatedName := ""
	var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < 12; i++ {
		generatedName += string(characters[rand.Intn(len(characters))])
	}
	isUserPresent := db.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: generatedName}})
	if isUserPresent.Err() == mongo.ErrNoDocuments {
		_, err := db.UserCollection.InsertOne(context.TODO(), bson.D{{Key: "username", Value: generatedName}})
		if err != nil {
			return "Error creating user"
		}
		return generatedName
	}
	retryCount := 0
	for isUserPresent.Err() == nil && retryCount < 3 {
		generatedName = ""
		for i := 0; i < 12; i++ {
			generatedName += string(characters[rand.Intn(len(characters))])
		}
		isUserPresent = db.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: generatedName}})
		retryCount++
	}
	if isUserPresent.Err() == nil {
		return "Error creating user"
	}
	if isUserPresent.Err() == mongo.ErrNoDocuments {
		_, err := db.UserCollection.InsertOne(context.TODO(), bson.D{{Key: "username", Value: generatedName}})
		if err != nil {
			return "Error creating user"
		}
		return generatedName
	}
	return "Error creating user"
}

func GetUserAllURLsService(username string) []map[string]interface{} {
	cursor, err := db.URLCollection.Find(context.TODO(), bson.D{{Key: "username", Value: username}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.TODO())
	var allURLs []map[string]interface{}
	for cursor.Next(context.TODO()) {
		var url map[string]interface{}
		err := cursor.Decode(&url)
		if err != nil {
			return nil
		}
		url = map[string]interface{}{
			"originalUrl": url["original_url"],
			"shortUrl":    url["short_url"],
			"clicks":      url["clicks"],
			"createdAt":   url["created_at"],
		}
		allURLs = append(allURLs, url)
	}
	return allURLs
}

func DeleteUrlService(username, shortId string) string {
	shortUrl := os.Getenv("BASE_URL") + "/" + shortId
	_, err := db.URLCollection.DeleteOne(context.TODO(), bson.D{{Key: "username", Value: username}, {Key: "short_url", Value: shortUrl}})
	if err != nil {
		return "Error deleting URL"
	}
	return ""
}
