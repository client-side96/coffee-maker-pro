package main

import (
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"
	"coffee-maker-pro/internal/server"
	"coffee-maker-pro/internal/statemachine"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func initSensors(tempSensor *sensor.Sensor, pressureSensor *sensor.Sensor) {
	log.Println("Initializing sensors...")
	sensor.Init(tempSensor, sensor.EnvironmentTemp)
	sensor.Init(pressureSensor, sensor.InitialPressure)
	//go sensor.ChangeTemperature(tempSensor, 95)
}

func main() {
	stateId, _ := primitive.ObjectIDFromHex(statemachine.STATEID)
	dbClient := database.Init()
	tempSensor := sensor.Create(sensor.TEMP)
	pressureSensor := sensor.Create(sensor.PRESSURE)
	initSensors(&tempSensor, &pressureSensor)

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

		switch statemachine.CoffeeMaker.State {
		case statemachine.Off:
			// Do nothing
		case statemachine.Idle:
		// Read current applied config from database
		// Read current sensor values
		// If those values are different -> transition to state applying, else -> transition to ready
		case statemachine.Applying:
		// change values of sensors until they are the same with the applied configuration
		// transition to Ready
		case statemachine.Ready:
		// stay in this state until StartBrewing or ChangeConfig
		case statemachine.Brewing:
			// not sure maybe do nothing
		default:
			println("Do nothing")
		}
	}
}
