package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database
var client *mongo.Client

func DisconnectDB() {
	client.Disconnect(context.TODO())
}

func InitDB() error {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli
	if err != nil {
		return err
	}

	// Crear la base de datos "fichahotel" (colecciones se crean automáticamente)
	MongoDb = client.Database("fichahotel")

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	fmt.Println("Available databases:")
	fmt.Println(dbNames)

	return nil
}
