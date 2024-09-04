package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri, err := env.MustGet("MONGODB_URI")
	if err != nil {
		log.Fatal("Error reading enviroment variable <MONGODB_URI>")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	// send a ping to confirm the db connection
	var result bson.M
	if err := Client.Database("admin").RunCommand(
		context.TODO(),
		bson.D{{"ping", 1}},
	).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("connected to database successfully")
}
