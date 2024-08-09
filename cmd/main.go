package main

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/features/events"
	"GymEventTracker/internal/features/members"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the SQLite database
	database.InitDB("gym_event_tracker.db")
	defer database.CloseDB()

	// Create a new Echo instance
	e := echo.New()

	// Setup routes
	events.SetupRoutes(e)
	members.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
