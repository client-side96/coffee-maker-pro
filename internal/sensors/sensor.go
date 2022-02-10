package sensors

import (
	"fmt"
	"os"
	"time"
)

const environmentTemp = 23

func Init(tempSensor *Sensor) {
	tempSensor.SetSensorValue(environmentTemp)
}

func Log(sensor *Sensor) {
	filename := "/var/log/coffee-maker-pro/" + string(sensor.sensorType) + ".log"
	sensorValue := sensor.GetSensorValue()
	timestamp := time.Now().Format(time.RFC3339Nano)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s %f\n", timestamp, sensorValue)); err != nil {
		panic(err)
	}
}
