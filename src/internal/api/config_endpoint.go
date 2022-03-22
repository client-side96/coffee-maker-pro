package api

import (
	"coffee-maker-pro/internal/config"
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/statemachine"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConfig(c *gin.Context) {
	var configs []*config.Config
	client := database.Init()
	findOptions := options.Find()
	results := database.QueryAll(client, database.CONFIG, bson.M{}, findOptions)
	defer results.Close(context.Background())
	for results.Next(context.Background()) {
		var c config.Config
		err := results.Decode(&c)
		if err != nil {
			log.Fatalln(err)
		}
		configs = append(configs, &c)
	}
	c.JSON(http.StatusOK, configs)
}
func GetConfigById(c *gin.Context) {
	client := database.Init()
	findOptions := options.FindOne()
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	oneConfig := database.Query[config.Config](client, database.CONFIG, bson.M{"_id": id}, findOptions)
	c.JSON(http.StatusOK, oneConfig)
}
func CreateConfig(c *gin.Context) {
	var payload config.Config
	client := database.Init()
	byteResult, err := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(byteResult, &payload)
	if err != nil {
		log.Fatalln(err)
	}
	result := database.Insert(client, database.CONFIG, payload)
	c.JSON(http.StatusCreated, result)
}
func UpdateConfig(c *gin.Context) {
	var payload config.Config
	client := database.Init()
	byteResult, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(byteResult, &payload)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	response := database.Update(client, database.CONFIG, bson.M{"$set": bson.M{
		"name":      payload.Name,
		"temp":      payload.Temp,
		"pressure":  payload.Pressure,
		"grinding":  payload.Grinding,
		"volume":    payload.Volume,
		"time":      payload.Time,
		"isApplied": payload.IsApplied,
	},
	}, bson.M{"_id": id})
	c.JSON(http.StatusOK, response)
}
func ApplyConfig(c *gin.Context) {
	client := database.Init()
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	appliedConfig := database.Query[config.Config](client, database.CONFIG, bson.M{"isApplied": true}, options.FindOne())
	database.Update(client, database.CONFIG, bson.M{"$set": bson.M{
		"name":      appliedConfig.Name,
		"temp":      appliedConfig.Temp,
		"pressure":  appliedConfig.Pressure,
		"grinding":  appliedConfig.Grinding,
		"volume":    appliedConfig.Volume,
		"time":      appliedConfig.Time,
		"isApplied": false,
	}}, bson.M{"_id": appliedConfig.ID})

	configToApply := database.Query[config.Config](client, database.CONFIG, bson.M{"_id": id}, options.FindOne())
	response := database.Update(client, database.CONFIG, bson.M{"$set": bson.M{
		"name":      configToApply.Name,
		"temp":      configToApply.Temp,
		"pressure":  configToApply.Pressure,
		"grinding":  configToApply.Grinding,
		"volume":    configToApply.Volume,
		"time":      configToApply.Time,
		"isApplied": true,
	}}, bson.M{"_id": id})
	rpcClient, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	var reply statemachine.StateType
	err = rpcClient.Call("RPCServer.ReceiveMessage", statemachine.TurnOn, &reply)
	if err != nil {
		log.Fatal("message error", err)
	}
	c.JSON(http.StatusOK, response)
}
func DeleteConfig(c *gin.Context) {
	client := database.Init()
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	response := database.Delete(client, database.CONFIG, bson.M{"_id": id})
	c.JSON(http.StatusOK, response)
}
