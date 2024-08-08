package routes

import (
	"GymEventTracker/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home)
	e.GET("/events", handlers.ListEvents)
	e.POST("/events", handlers.CreateEvent)

	// Members routes
	e.GET("/members", handlers.GetMembers)          // Retrieve all members
	e.POST("/members", handlers.CreateMember)       // Add a new member
	e.GET("/members/:id", handlers.GetMember)       // Retrieve a member by ID
	e.PUT("/members/:id", handlers.UpdateMember)    // Update a member
	e.DELETE("/members/:id", handlers.DeleteMember) // Delete a member
}
