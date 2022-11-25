package db

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyUser allows modify the user profile
func ModifyUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("users")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		register["lastname"] = user.Lastname
	}
	register["birthday"] = user.Birthday
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	updatedUser := bson.M{
		"$set": register,
	}

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{ "_id": bson.M{ "$eq": objectID } }

	_, err := collection.UpdateOne(ctx, filter, updatedUser)
	if err != nil {
		return false, err
	}

	return true, nil
}