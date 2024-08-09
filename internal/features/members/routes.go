package members

import (
	"GymEventTracker/internal/database/members"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	repo := members.NewSqlLiteMemberRepo()

	// Members routes
	e.GET("/members", getMembers(repo))          // Retrieve all members
	e.POST("/members", createMember(repo))       // Add a new member
	e.GET("/members/:id", showMember)            // Retrieve a member by ID
	e.PUT("/members/:id", updateMember)          // Update a member
	e.DELETE("/members/:id", deleteMember(repo)) // Delete a member
}
