package sensor

import (
	"log"
	"time"
)

func ChangeTemperature(tempSensor *Sensor, target float64) {
	log.Printf("Changing to target temperature: %f", target)
	current := tempSensor.GetSensorValue()
	for current != target {
		current = current + 2
		if current > target {
			tempSensor.SetSensorValue(target)
			return
		}
		tempSensor.SetSensorValue(current)
		time.Sleep(time.Second)
	}
}
