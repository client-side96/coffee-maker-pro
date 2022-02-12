package main

import (
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"
	"log"
	"time"
)

func initCoffeeMaker(tempSensor *sensor.Sensor, pressureSensor *sensor.Sensor) {
	sensor.Init(tempSensor, sensor.EnvironmentTemp)
	sensor.Init(pressureSensor, sensor.InitialPressure)
	go sensor.ChangeTemperature(tempSensor, 95)
}

func main() {
	log.Println("Starting coffee maker pro...")
	tempSensor := sensor.Create(sensor.TEMP)
	pressureSensor := sensor.Create(sensor.PRESSURE)

	dbClient := database.Init()
	initCoffeeMaker(&tempSensor, &pressureSensor)

	lastTemp := sensor.Create(sensor.TEMP)
	lastPress := sensor.Create(sensor.PRESSURE)

	for true {
		sensor.Log(dbClient, &tempSensor, &lastTemp)
		sensor.Log(dbClient, &pressureSensor, &lastPress)
		time.Sleep(time.Second)
	}
}
