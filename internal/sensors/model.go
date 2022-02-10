package sensors

import "time"

const TEMP_NAME = "temperature_sensor_0001"
const PRESSURE_NAME = "pressure_sensor_0001"

const (
	TEMP     SensorType = "temperature"
	PRESSURE SensorType = "pressure"
)

type SensorType string

type Sensor struct {
	id         int
	name       string
	value      float64
	sensorType SensorType
	timestamp  string
}

func (s *Sensor) SetSensorValue(newValue float64) {
	s.value = newValue
}

func (s Sensor) GetSensorValue() float64 {
	return s.value
}

func getNameBySensorType(sensorType SensorType) string {
	if sensorType == TEMP {
		return TEMP_NAME
	}
	return PRESSURE_NAME
}

func Create(sensorType SensorType) Sensor {
	return Sensor{
		id:         1,
		name:       getNameBySensorType(sensorType),
		value:      0,
		sensorType: sensorType,
		timestamp:  time.Now().Format(time.RFC3339Nano),
	}
}
