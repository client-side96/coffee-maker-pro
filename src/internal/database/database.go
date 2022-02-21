package database

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const USER = "cmp-user"
const PW = "admin"
const HOST = "localhost"
const PORT = "27011"
const DB = "cmp-db"

var ctx = context.TODO()
var connectionString = "mongodb://" + USER + ":" + PW + "@" + HOST + ":" + PORT + "/" + DB + "?connect=direct"

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

func Query[T any](client *mongo.Client, collection CollectionName, queryFilter bson.M, opts *options.FindOneOptions) T {
	var result T
	coll := client.Database(DB).Collection(string(collection))
	err := coll.FindOne(ctx, queryFilter, opts).Decode(&result)
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

func Update(client *mongo.Client, collection CollectionName, payload bson.M, queryFiler bson.M) *mongo.UpdateResult {
	coll := client.Database(DB).Collection(string(collection))
	result, err := coll.UpdateOne(ctx, queryFiler, payload)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Watch(client *mongo.Client, collection CollectionName) *mongo.ChangeStream {
	coll := client.Database(DB).Collection(string(collection))
	cursor, err := coll.Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		log.Fatal(err)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return cursor

}

func PrepareWatchResult[T any](fullDocument interface{}, template T) []byte {
	bsonBytes, err := bson.Marshal(fullDocument)
	if err != nil {
		log.Fatal(err)
	}

	bson.Unmarshal(bsonBytes, &template)

	result, err := json.Marshal(template)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
