package sensor

import (
	"coffee-maker-pro/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
)

const EnvironmentTemp = 23
const InitialPressure = 0.5
const InitialVolume = 18
const InitialTime = 20
const InitialGrinding = 25

func getNameBySensorType(sensorType SensorType) string {
	if sensorType == TEMP {
		return TEMP_NAME
	} else if sensorType == VOLUME {
		return VOLUME_NAME
	} else if sensorType == TIME {
		return TIME_NAME
	} else if sensorType == GRINDING {
		return GRINDING_NAME
	}
	return PRESSURE_NAME
}

func logValue(sensorValue float64, timestamp string) {
	filename := "/var/log/coffee-maker-pro/sensors.log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprintf("%s %f\n", timestamp, sensorValue)); err != nil {
		panic(err)
	}
}

func writeToDB(dbClient *mongo.Client, sensor DBSensor) {
	database.Insert[DBSensor](dbClient, database.SENSORS, sensor)
}

func Create(sensorType SensorType) Sensor {
	log.Printf("Starting sensor: %s", sensorType)
	return Sensor{
		Name:       getNameBySensorType(sensorType),
		Value:      0,
		SensorType: sensorType,
	}
}

func Init(tempSensor *Sensor, value float64) {
	log.Printf("Initializing sensor %s with %f", tempSensor.Name, value)
	tempSensor.SetSensorValue(value)
}

func ReadValue(dbClient *mongo.Client, sensor *Sensor) {
	sensorValue := sensor.GetSensorValue()
	timestamp := time.Now().Format(time.RFC3339Nano)
	sensorDB := sensor.ToDB()
	if sensor.Value != sensor.PrevValue {
		log.Printf("%s: Sensor Value changed to %f", sensor.Name, sensor.Value)
		writeToDB(dbClient, sensorDB)
		sensor.SetSensorPrevValue(sensor.Value)
	}
	logValue(sensorValue, timestamp)
}
