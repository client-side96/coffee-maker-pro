package api

import (
	"coffee-maker-pro/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	result := database.Watch(client, database.SENSORS)

	conn.WriteMessage(1, []byte(fmt.Sprintf("%v", result)))

}

func SensorEndpoint(c *gin.Context) {

	sensorWs(c.Writer, c.Request)
}
