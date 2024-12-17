package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var URLCollection *mongo.Collection

func init() {
	mongodbURI := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		panic(err)
	}
	URLCollection = client.Database("shortit").Collection("urls")
}
