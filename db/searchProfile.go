package db

import (
	"context"
	"fmt"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile receives an ID, search the user with that and returns it
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("users")

	var profile models.User
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("User not found" + err.Error())
		return profile, err
	}

	return profile, nil
}