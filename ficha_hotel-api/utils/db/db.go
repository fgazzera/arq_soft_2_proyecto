package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database
var client *mongo.Client

func DisconnectDB() {
	client.Disconnect(context.TODO())
}

func InitDB() error {
	clientOpts := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli
	if err != nil {
		return err
	}

	MongoDb = client.Database("arq_soft_2")

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	fmt.Println("Available databases:")
	fmt.Println(dbNames)

	return nil
}
