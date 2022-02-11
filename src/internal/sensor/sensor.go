package sensor

import (
	"coffee-maker-pro/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

const EnvironmentTemp = 90
const InitialPressure = 0.5

func getNameBySensorType(sensorType SensorType) string {
	if sensorType == TEMP {
		return TEMP_NAME
	}
	return PRESSURE_NAME
}

func Init(tempSensor *Sensor, value float64) {
	tempSensor.SetSensorValue(value)
}

func Create(sensorType SensorType) Sensor {
	return Sensor{
		name:       getNameBySensorType(sensorType),
		value:      0,
		sensorType: sensorType,
	}
}

func Log(dbClient *mongo.Client, sensor *Sensor, last *Sensor) {
	filename := "/var/log/coffee-maker-pro/sensors.log"
	sensorValue := sensor.GetSensorValue()
	timestamp := time.Now().Format(time.RFC3339Nano)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	sensorDB := sensor.ToDB()
	if sensor.value != last.value {
		database.Insert[DBSensor](dbClient, database.SENSORS, sensorDB)
		last.SetSensorValue(sensor.value)
	}
	if _, err = f.WriteString(fmt.Sprintf("%s %f\n", timestamp, sensorValue)); err != nil {
		panic(err)
	}
}
