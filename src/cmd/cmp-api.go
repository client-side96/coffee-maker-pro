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
	}

	router.Run()
}
