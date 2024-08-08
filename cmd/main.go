package main

import (
	"GymEventTracker/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	//comment 1
	r.Run(":3000")
}
