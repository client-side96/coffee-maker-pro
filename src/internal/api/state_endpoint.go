package api

import (
	"coffee-maker-pro/internal/statemachine"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/rpc"
)

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
