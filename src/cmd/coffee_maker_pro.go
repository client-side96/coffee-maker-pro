package main

import (
	"coffee-maker-pro/internal/config"
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"
	"coffee-maker-pro/internal/server"
	"coffee-maker-pro/internal/statemachine"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func initSensors(tempSensor *sensor.Sensor, pressureSensor *sensor.Sensor, volumeSensor *sensor.Sensor, timeSensor *sensor.Sensor, grindingSensor *sensor.Sensor) {
	log.Println("Initializing sensors...")
	sensor.Init(tempSensor, sensor.EnvironmentTemp)
	sensor.Init(pressureSensor, sensor.InitialPressure)
	sensor.Init(volumeSensor, sensor.InitialVolume)
	sensor.Init(timeSensor, sensor.InitialTime)
	sensor.Init(grindingSensor, sensor.InitialGrinding)
}

func main() {
	stateId, _ := primitive.ObjectIDFromHex(statemachine.STATEID)
	dbClient := database.Init()
	tempSensor := sensor.Create(sensor.TEMP)
	pressureSensor := sensor.Create(sensor.PRESSURE)
	volumeSensor := sensor.Create(sensor.VOLUME)
	timeSensor := sensor.Create(sensor.TIME)
	grindingSensor := sensor.Create(sensor.GRINDING)
	initSensors(&tempSensor, &pressureSensor, &volumeSensor, &timeSensor, &grindingSensor)

	rpcServer := new(server.RPCServer)
	rpc.Register(rpcServer)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(l, nil)

	for true {
		database.Update(dbClient, database.STATUS, bson.M{"$set": bson.M{"value": statemachine.CoffeeMaker.State}}, bson.M{"_id": stateId})
		sensor.ReadValue(dbClient, &tempSensor)
		sensor.ReadValue(dbClient, &pressureSensor)
		sensor.ReadValue(dbClient, &volumeSensor)
		sensor.ReadValue(dbClient, &timeSensor)
		sensor.ReadValue(dbClient, &grindingSensor)

		switch statemachine.CoffeeMaker.State {
		case statemachine.Off:
			// Do nothing
		case statemachine.Idle:
			isMatchingSensorValues := config.IsAppliedConfigMatchingSensorValues(dbClient, &tempSensor, &pressureSensor, &volumeSensor, &timeSensor, &grindingSensor)
			log.Printf("Is Matching Sensor Values: %b", isMatchingSensorValues)
			if isMatchingSensorValues {
				statemachine.TransitionState(statemachine.SetReady)
			} else {
				statemachine.TransitionState(statemachine.ApplyConfig)
			}
		// Read current applied config from database
		// Read current sensor values
		// If those values are different -> transition to state applying, else -> transition to ready
		case statemachine.Applying:
			isMatchingSensorValues := config.IsAppliedConfigMatchingSensorValues(dbClient, &tempSensor, &pressureSensor, &volumeSensor, &timeSensor, &grindingSensor)
			if isMatchingSensorValues {
				statemachine.TransitionState(statemachine.SetReady)
			} else {
				config.ApplyConfig(dbClient, &tempSensor, &pressureSensor, &volumeSensor, &timeSensor, &grindingSensor)
			}

		// change values of sensors until they are the same with the applied configuration
		// transition to Ready
		case statemachine.Ready:
		// stay in this state until StartBrewing or ChangeConfig
		case statemachine.Brewing:
			// not sure maybe do nothing
		default:
			println("Do nothing")
		}
		time.Sleep(time.Second)
	}
}
