package config

import (
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Temp      float64            `bson:"temp" json:"temp"`
	Pressure  float64            `bson:"pressure" json:"pressure"`
	Grinding  float64            `bson:"grinding" json:"grinding"`
	Volume    float64            `bson:"volume" json:"volume"`
	Time      float64            `bson:"time" json:"time"`
	IsApplied bool               `bson:"isApplied" json:"isApplied"`
}

func IsAppliedConfigMatchingSensorValues(dbClient *mongo.Client, tempSensor *sensor.Sensor, pressureSensor *sensor.Sensor, volumeSensor *sensor.Sensor, timeSensor *sensor.Sensor, grindingSensor *sensor.Sensor) bool {
	appliedConfig := database.Query[Config](dbClient, database.CONFIG, bson.M{"isApplied": true}, options.FindOne())
	return appliedConfig.Temp == tempSensor.Value && appliedConfig.Pressure == pressureSensor.Value && appliedConfig.Volume == volumeSensor.Value && appliedConfig.Time == timeSensor.Value && appliedConfig.Grinding == grindingSensor.Value
}

func ApplyConfig(dbClient *mongo.Client, tempSensor *sensor.Sensor, pressureSensor *sensor.Sensor, volumeSensor *sensor.Sensor, timeSensor *sensor.Sensor, grindingSensor *sensor.Sensor) {
	appliedConfig := database.Query[Config](dbClient, database.CONFIG, bson.M{"isApplied": true}, options.FindOne())
	go sensor.ChangeSensor(tempSensor, appliedConfig.Temp)
	go sensor.ChangeSensor(pressureSensor, appliedConfig.Pressure)
	go sensor.ChangeSensor(volumeSensor, appliedConfig.Volume)
	go sensor.ChangeSensor(timeSensor, appliedConfig.Time)
	go sensor.ChangeSensor(grindingSensor, appliedConfig.Grinding)
}
