package middlewares

import (
	"context"
	"net/http"
	"shortit/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement rate limiter logic here
		cke, err := r.Cookie("username")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if err == http.ErrNoCookie {
			next.ServeHTTP(w, r)
			return
		}
		username := cke.Value
		if username != "" {
			result := db.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}})
			if result.Err() != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			var user map[string]interface{}
			err = result.Decode(&user)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			currentUnixTime := time.Now().Unix()
			if user["rateLimitReset"].(int64) < currentUnixTime {
				_, err = db.UserCollection.UpdateOne(context.TODO(), bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "rateLimit", Value: 10}, {Key: "rateLimitReset", Value: currentUnixTime + 300}}}})
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
			if user["rateLimit"].(int32) <= 0 {
				http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
				return
			} else {
				_, err := db.UserCollection.UpdateOne(context.TODO(), bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$inc", Value: bson.D{{Key: "rateLimit", Value: -1}}}})
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
