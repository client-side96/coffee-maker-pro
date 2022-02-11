package sensor

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const TEMP_NAME = "temperature_sensor_0001"
const PRESSURE_NAME = "pressure_sensor_0001"

const (
	TEMP     SensorType = "temperature"
	PRESSURE SensorType = "pressure"
)

type SensorType string

type Sensor struct {
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

func (s Sensor) ToDB() DBSensor {
	timestamp := time.Now()
	return DBSensor{
		ID:         primitive.ObjectID{},
		Value:      s.value,
		SensorType: s.sensorType,
		Timestamp:  primitive.NewDateTimeFromTime(timestamp),
	}
}

type DBSensor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Value      float64            `bson:"value,omitempty"`
	SensorType SensorType         `bson:"type,omitempty"`
	Timestamp  primitive.DateTime `bson:"timestamp,omitempty"`
}
