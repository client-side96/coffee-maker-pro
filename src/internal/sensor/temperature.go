package sensor

import "time"

func ChangeTemperature(tempSensor *Sensor, target float64) {
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
