package db

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// CreateRelation routinr save a relationship in a database
func CreateRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("relations")

	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}
	
	return true, nil
}