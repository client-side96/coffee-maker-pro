package main

import (
	"coffee-maker-pro/internal/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	v1 := router.Group("/api")
	{
		v1.GET("/sensors", api.GetSensorData)
		v1.GET("/status", api.GetStatus)
		v1.POST("/status/power/on", api.TurnOn)
		v1.POST("/status/power/off", api.TurnOff)
		// POST /state/brew/start -> Transition StartBrewing
		// POST /state/brew/stop => Transition StopBrewing
		// GET /config -> Retrieve current configuration
		// POST /config/create -> Create new configuration
		// PUT /config/update -> Update configuration
		// DELETE /config/delete -> Delete configuration
		// POST /config/apply -> Transition ChangeConfig
	}

	router.Run()
}
