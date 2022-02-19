package main

import (
	"coffee-maker-pro/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/sensors", api.SensorEndpoint)
		v1.POST("/power/on", api.TurnOn)
		// power/off -> Transition TurnOff
		// config/apply -> Transition ChangeConfig
		// config/create -> Create new configuration
		// config/update -> Update configuration
		// config/delete -> Delete configuration
		// coffee/start -> Transition StartBrewing
		// coffee/stop => Transition StopBrewing
	}

	router.Run()
}
