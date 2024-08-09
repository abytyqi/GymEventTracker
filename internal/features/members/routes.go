package members

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Members routes
	e.GET("/members", getMembers)          // Retrieve all members
	e.POST("/members", createMember)       // Add a new member
	e.GET("/members/:id", getMember)       // Retrieve a member by ID
	e.PUT("/members/:id", updateMember)    // Update a member
	e.DELETE("/members/:id", deleteMember) // Delete a member
}
