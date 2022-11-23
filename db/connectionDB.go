package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCNN is the connection object to the database
var MongoCNN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user_node_cafe:gonzapi1@miclustercafe.d0z9e.mongodb.net/GolangDB")

// ConnectDB is the function that allows connecting the database
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("BASE DE DATOS CONECTADA EXISTOSAMENTE")
	return client
}

// CheckConnection checks the ping of the database
func CheckConnection() int {
	err := MongoCNN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}