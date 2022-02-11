package main

import (
	"bufio"
	"coffee-maker-pro/internal/database"
	"coffee-maker-pro/internal/sensor"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func initSensors(temp *sensor.Sensor, pressure *sensor.Sensor) {
	sensor.Init(temp)
	sensor.Init(pressure)
}

func readAndUpdateSensors(temp *sensor.Sensor, pressure *sensor.Sensor) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command := input[0]
		newValue, err := strconv.ParseFloat(input[1], 64)
		if err != nil {
			fmt.Errorf("%w\n", err)
		} else {
			if command == "temp_up" {
				temp.SetSensorValue(newValue)
				log.Println("New temperature set.")
			} else if command == "press_up" {
				pressure.SetSensorValue(newValue)
				log.Println("New pressure set.")
			}
		}
	}

}

func main() {
	log.Println("Starting coffee maker pro...")
	tempSensor := sensor.Create(sensor.TEMP)
	pressureSensor := sensor.Create(sensor.PRESSURE)

	dbClient := database.Init()
	database.Query[interface{}](dbClient, database.SENSORS)

	initSensors(&tempSensor, &pressureSensor)

	go readAndUpdateSensors(&tempSensor, &pressureSensor)

	for true {
		//sensors.Log(&tempSensor)
		//sensors.Log(&pressureSensor)
		time.Sleep(time.Second)
	}
}
