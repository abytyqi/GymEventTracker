package users

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, repo database.UserRepo) {
	// Members routes
	e.GET("/users", getUsers(repo)) // Retrieve all members
	//e.POST("/users", createUser(repo))       // Add a new member
	//e.GET("/users/:id", showUser)            // Retrieve a member by ID
	//e.PUT("/users/:id", updateUser)          // Update a member
	//e.DELETE("/users/:id", deleteUser(repo)) // Delete a member
}
