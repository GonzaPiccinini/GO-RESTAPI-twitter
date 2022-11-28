package db

import (
	"context"
	"log"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTweets return all paged tweets that match the condition
func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("tweet")

	var result []*models.GetTweets

	condition := bson.M{
		"user_id": ID,
	}

	option := options.Find()
	option.SetSort(bson.D{{ Key: "date", Value: -1 }})
	option.SetSkip((page - 1) * 20)
	option.SetLimit(20)

	index, err := collection.Find(ctx, condition, option)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	
	for index.Next(context.TODO()) {
		var register models.GetTweets
		err := index.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true 
}