package db

import (
	"context"
	"fmt"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUsers returns all users
func GetAllUsers(ID string, page int64, search string, typeUser string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	database := MongoCNN.Database("GolangDB")
	collection := database.Collection("users")

	var results []*models.User
	
	findOptions := options.Find()
	findOptions.SetSkip((page -1) * 20)
	findOptions.SetLimit(20)
	
	condition := bson.M{
		"name": bson.M{ "$regex": `(?i)` + search },
	}

	index, err := collection.Find(ctx, condition, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var status, include bool

	for index.Next(ctx) {
		var user models.User
		err := index.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation
		relation.UserID = ID
		relation.UserRelationID = user.ID.Hex()

		include = false
		
		status, _ = CheckRelation(relation)
		if typeUser == "unfollow" && !status {
			include = true
		}
		if typeUser == "follow" && status {
			include = true
		}
		if relation.UserRelationID  == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Biography = ""
			user.Banner = ""
			user.Email = ""
			user.Location = ""
			user.Website = ""

			results = append(results, &user)
		}

	}

	err = index.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	index.Close(ctx)

	return results, true
}