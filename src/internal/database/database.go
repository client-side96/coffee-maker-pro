package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const USER = "cmp-user"
const PW = "admin"
const HOST = "localhost"
const PORT = "27017"
const DB = "cmp-db"

var ctx = context.TODO()
var connectionString = "mongodb://" + USER + ":" + PW + "@" + HOST + ":" + PORT + "/" + DB

func Init() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//defer client.Disconnect(ctx)

	return client
}

func Query[T any](client *mongo.Client, collection CollectionName) T {
	var result T
	coll := client.Database(DB).Collection(string(collection))
	err := coll.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Insert[T any](client *mongo.Client, collection CollectionName, payload T) *mongo.InsertOneResult {
	coll := client.Database(DB).Collection(string(collection))
	result, err := coll.InsertOne(ctx, payload)
	if err != nil {
		log.Fatal(err)
	}
	return result
}