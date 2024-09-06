package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, repo database.MemberRepo) {
	// Members routes
	e.GET("/members", getMembers(repo))          // Retrieve all members
	e.POST("/members", createMember(repo))       // Add a new member
	e.GET("/members/:id", showMember(repo))      // Retrieve a member by ID
	e.PUT("/members/:id", updateMember(repo))    // Update a member
	e.DELETE("/members/:id", deleteMember(repo)) // Delete a member
}
