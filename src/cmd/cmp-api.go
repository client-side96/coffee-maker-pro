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
		v1.GET("/config", api.GetConfig)
		v1.GET("/config/:id", api.GetConfigById)
		v1.POST("/config", api.CreateConfig)
		v1.PUT("/config/:id", api.UpdateConfig)
		v1.DELETE("/config/:id", api.DeleteConfig)
		v1.PUT("/config/apply/:id", api.ApplyConfig)
		// POST /state/brew/start -> Transition StartBrewing
		// POST /state/brew/stop => Transition StopBrewing
		// POST /config/apply -> Transition ChangeConfig
	}

	router.Run()
}
