package api

import (
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	dbClient := database.Init()
	cursor := database.Watch(dbClient, database.SENSORS)
	findOptions := options.FindOne()
	findOptions.SetSort(bson.M{"timestamp": -1})

	lastTemperature := database.Query[sensor.DBSensor](dbClient, database.SENSORS, bson.M{"type": sensor.TEMP}, findOptions)
	lastPressure := database.Query[sensor.DBSensor](dbClient, database.SENSORS, bson.M{"type": sensor.PRESSURE}, findOptions)
	lastTemperatureB, _ := json.Marshal(lastTemperature)
	lastPressureB, _ := json.Marshal(lastPressure)

	conn.WriteMessage(1, lastTemperatureB)
	conn.WriteMessage(1, lastPressureB)

	var changeStreamValue bson.M
	var template sensor.DBSensor
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&changeStreamValue); err != nil {
			log.Fatal(err)
		}
		result := database.PrepareWatchResult(changeStreamValue["fullDocument"], template)

		conn.WriteMessage(1, result)
	}

}

func SensorEndpoint(c *gin.Context) {

	sensorWs(c.Writer, c.Request)
}
