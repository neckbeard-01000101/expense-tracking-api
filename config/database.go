package config

import (
	"context"
	"fmt"
	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func InitDB() {
	uri, err := env.MustGet("MONGODB_URI")
	if err != nil {
		log.Fatal("Error reading enviroment variable <MONGODB_URI>")
	}
	_ = uri
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// send a ping to confirm the db connection
	var result bson.M
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
