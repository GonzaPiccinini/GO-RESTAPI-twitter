package db

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveTweet save a new tweet in the database
func SaveTweet(tweet models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("tweet")

	register := bson.M{
		"user_id": tweet.UserID,
		"message": tweet.Message,
		"date": tweet.Date,
	}

	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	objectID, _ := result.InsertedID.(primitive.ObjectID)
	
	return objectID.String(), true, nil
}