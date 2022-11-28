package db

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFollowersTweets routine gets all tweets of the followers
func GetFollowersTweets(ID string, page int) ([]models.FollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("relations")
	
	var results []models.FollowersTweets

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{ "$match": bson.M{ "user_id": ID } })
	conditions = append(conditions, bson.M{ 
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "user_relation_id",
			"foreignField": "user_id",
			"as": "tweet",
		}})
	conditions = append(conditions, bson.M{ "$unwind": "$tweet" })
	conditions = append(conditions, bson.M{ "$sort": bson.M{ "date": -1 } })
	conditions = append(conditions, bson.M{ "$skip": skip })
	conditions = append(conditions, bson.M{ "$limit": 20 })

	index, _ := collection.Aggregate(ctx, conditions)
	
	err := index.All(ctx, &results)
	if err != nil {
		return results, false
	}

	return results, true
}