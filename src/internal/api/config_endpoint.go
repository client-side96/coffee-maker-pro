package api

import (
	"coffee-maker-pro/internal/config"
	"coffee-maker-pro/internal/database"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
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
	result := database.Insert[config.Config](client, database.CONFIG, payload)
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
		"temp":      payload.Temp,
		"pressure":  payload.Pressure,
		"grinding":  payload.Grinding,
		"isApplied": payload.IsApplied,
	},
	}, bson.M{"_id": id})
	c.JSON(http.StatusOK, response)
}
func DeleteConfig(c *gin.Context) {
	client := database.Init()
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	response := database.Delete(client, database.CONFIG, bson.M{"_id": id})
	c.JSON(http.StatusOK, response)
}
