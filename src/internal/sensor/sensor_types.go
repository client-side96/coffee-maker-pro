package sensor

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const TEMP_NAME = "temperature_sensor_0001"
const PRESSURE_NAME = "pressure_sensor_0001"
const VOLUME_NAME = "volume_sensor_0001"
const TIME_NAME = "time_sensor_0001"
const GRINDING_NAME = "grinding_sensor_0001"

const (
	TEMP     SensorType = "temperature"
	PRESSURE SensorType = "pressure"
	VOLUME   SensorType = "volume"
	TIME     SensorType = "time"
	GRINDING SensorType = "grinding"
)

type SensorType string

type Sensor struct {
	Name       string
	Value      float64
	PrevValue  float64
	SensorType SensorType
	Timestamp  string
}

func (s *Sensor) SetSensorValue(newValue float64) {
	s.Value = newValue
}

func (s *Sensor) SetSensorPrevValue(newValue float64) {
	s.PrevValue = newValue
}

func (s Sensor) GetSensorValue() float64 {
	return s.Value
}

func (s Sensor) ToDB() DBSensor {
	timestamp := time.Now()
	return DBSensor{
		ID:         primitive.ObjectID{},
		Value:      s.Value,
		SensorType: s.SensorType,
		Timestamp:  primitive.NewDateTimeFromTime(timestamp),
	}
}

type DBSensor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Value      float64            `bson:"value,omitempty" json:"value"`
	SensorType SensorType         `bson:"type,omitempty" json:"sensorType"`
	Timestamp  primitive.DateTime `bson:"timestamp,omitempty" json:"timestamp"`
}
