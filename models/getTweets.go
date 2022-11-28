package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTweets struct {
	ID 		primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID 	string `bson:"user_id" json:"userId,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date 	time.Time `bson:"date" json:"date,omitempty"`
}