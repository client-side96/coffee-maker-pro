package sensor

import (
	"log"
	"time"
)

func ChangeSensor(sensor *Sensor, target float64) {
	log.Printf("Changing to target temperature: %f", target)
	current := sensor.GetSensorValue()
	for current != target {
		current = current + 2
		if current > target {
			sensor.SetSensorValue(target)
			return
		}
		sensor.SetSensorValue(current)
		time.Sleep(time.Second)
	}
}
