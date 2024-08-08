package main

import (
	"GymEventTracker/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":3000"))
}
