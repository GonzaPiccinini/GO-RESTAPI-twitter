package db

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckRelation checks if a user ir related to another
func CheckRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("relations")
	
	condition := bson.M{
		"user_id": relation.UserID,
		"user_relation_id": relation.UserRelationID,
	}

	var result models.Relation

	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}