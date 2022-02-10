package main

import (
	"bufio"
	"coffee-maker-pro/internal/sensors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const environmentTemp = 23

func initializeTemperature(tempSensor *sensors.Sensor) {
	tempSensor.SetSensorValue(environmentTemp)
}

func main() {
	tempSensor := sensors.Create(sensors.PRESSURE)
	initializeTemperature(&tempSensor)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			newValue, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Errorf("%w\n", err)
			} else {
				tempSensor.SetSensorValue(newValue)
			}
		}
	}()
	for true {
		fmt.Printf("%f\n", tempSensor.GetSensorValue())
		time.Sleep(time.Second)
	}
}
