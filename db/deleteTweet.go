package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteTweet delete a specific tweet of the database
func DeleteTweet(ID string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("tweet")

	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectID,
		"user_id": userId,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}