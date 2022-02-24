package api

import (
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/statemachine"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/rpc"
)

func statusWs(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbClient := database.Init()
	cursor := database.Watch(dbClient, database.STATUS)
	findOptions := options.FindOne()
	findOptions.SetSort(bson.M{"timestamp": -1})

	state, _ := json.Marshal(database.Query[statemachine.StateMachine](dbClient, database.STATUS, bson.M{}, findOptions))

	conn.WriteMessage(1, state)

	var changeStreamValue bson.M
	var template statemachine.StateMachine
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&changeStreamValue); err != nil {
			log.Fatal(err)
		}
		result := database.PrepareWatchResult(changeStreamValue["fullDocument"], template)

		conn.WriteMessage(1, result)
	}

}

func GetStatus(c *gin.Context) {
	statusWs(c.Writer, c.Request)
}

func TurnOn(c *gin.Context) {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	var reply statemachine.StateType
	err = client.Call("RPCServer.ReceiveMessage", statemachine.TurnOn, &reply)
	if err != nil {
		log.Fatal("message error", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": reply,
	})
}

func TurnOff(c *gin.Context) {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	var reply statemachine.StateType
	err = client.Call("RPCServer.ReceiveMessage", statemachine.TurnOff, &reply)
	if err != nil {
		log.Fatal("message error", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": reply,
	})
}
