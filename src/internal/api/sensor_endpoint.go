package api

import (
	"coffee-maker-pro/internal/database"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sensorWs(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := database.Init()
	ctx := context.Background()
	sensorCollection := client.Database(database.DB).Collection(string(database.SENSORS))
	matchStage := bson.D{{"$match", bson.D{{"operationType", "insert"}}}}
	changeStream, _ := sensorCollection.Watch(ctx, mongo.Pipeline{matchStage})
	defer changeStream.Close(ctx)
	for changeStream.Next(ctx) {
		var result bson.M
		if err := changeStream.Decode(&result); err != nil {
			log.Fatal(err)
		}
		conn.WriteMessage(1, []byte(fmt.Sprintf("%v", result)))
	}

	if err := changeStream.Err(); err != nil {
		log.Fatal(err)
	}
}

func SensorEndpoint(c *gin.Context) {

	sensorWs(c.Writer, c.Request)
}
