package models

import (
	"time"
)

// SaveTweet is the tweet structure
type SaveTweet struct {
	UserID 	string 		`bson:"user_id" json:"userId,omitempty"`
	Message string 		`bson:"message" json:"message,omitempty"`
	Date 	time.Time 	`bson:"date" json:"date,omitempty"`
}