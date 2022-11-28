package models

// Tweet decode the body request
type Tweet struct {
	Message string `bson:"message" json:"message"`
}